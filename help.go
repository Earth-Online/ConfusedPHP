package confusedPHP

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"math/rand"
	"reflect"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// get random string
func RandStringBytes(n uint) string {
	if n < 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// php definition E.g function
var DefinitionType = []reflect.Type{
	reflect.TypeOf(&stmt.Class{}),
	reflect.TypeOf(&stmt.Function{}),
	reflect.TypeOf(&stmt.ClassMethod{}),
	reflect.TypeOf(&stmt.Interface{}),
	reflect.TypeOf(&stmt.Trait{}),
}

var ProcessControlType = []reflect.Type{
	reflect.TypeOf(&stmt.If{}),
	reflect.TypeOf(&stmt.Else{}),
	reflect.TypeOf(&stmt.ElseIf{}),
	reflect.TypeOf(&stmt.Switch{}),
	reflect.TypeOf(&stmt.Case{}),
	reflect.TypeOf(&stmt.Do{}),
	reflect.TypeOf(&stmt.While{}),
	reflect.TypeOf(&stmt.CaseList{}),
}

var HaveReturnType = append([]reflect.Type{
	reflect.TypeOf(&expr.FunctionCall{}),
	reflect.TypeOf(&expr.ShortArray{}),
	reflect.TypeOf(&expr.Array{}),
	reflect.TypeOf(&expr.Variable{}),
	reflect.TypeOf(&expr.InstanceOf{}),
	reflect.TypeOf(&expr.FunctionCall{}),
	reflect.TypeOf(&expr.Eval{}),
	reflect.TypeOf(&expr.List{}),
	reflect.TypeOf(&expr.Empty{}),
	reflect.TypeOf(&expr.Isset{}),
	reflect.TypeOf(&expr.ShellExec{}),
	reflect.TypeOf(&expr.Print{}),
	reflect.TypeOf(&expr.StaticPropertyFetch{}),
	reflect.TypeOf(&expr.StaticCall{}),
}, ConstantType...)

var NumType = []reflect.Type{
	reflect.TypeOf(&scalar.Lnumber{}),
	reflect.TypeOf(&scalar.Dnumber{}),
}

var StringType = []reflect.Type{
	reflect.TypeOf(&scalar.String{}),
	reflect.TypeOf(&scalar.Encapsed{}),
	reflect.TypeOf(&scalar.Heredoc{}),
}

var ConstantType = append(StringType, NumType...)

// Determine if node  is  sDefinitionType
func IsDefinitionType(n node.Node) bool {
	for _, val := range DefinitionType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

func IsHaveReturnType(n node.Node) bool {
	for _, val := range HaveReturnType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

func IsConstantType(n node.Node) bool {
	for _, val := range ConstantType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

func IsStringType(n node.Node) bool {
	for _, val := range StringType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

func IsFullyStringType(n node.Node) bool {
	if !IsStringType(n) {
		return false
	}
	switch n.(type) {
	case *scalar.Heredoc:
		for _, value := range n.(*scalar.Heredoc).Parts {
			switch value.(type) {
			case *expr.Variable:
				return false
			}
		}
	case *scalar.Encapsed:
		for _, value := range n.(*scalar.Heredoc).Parts {
			switch value.(type) {
			case *expr.Variable:
				return false
			}
		}
	}
	return true
}

func GetStingTypeValue(n node.Node) (str string, err error) {
	if !IsStringType(n) {
		return "", errors.New("not string type")
	}
	switch n.(type) {
	case *scalar.String:
		return n.(*scalar.String).Value, nil
	case *scalar.Heredoc:
		tmp := ""
		for _, value := range n.(*scalar.Heredoc).Parts {
			switch value.(type) {
			case *VarStr:
				tmp = tmp + value.(*VarStr).String()
			case *scalar.EncapsedStringPart:
				tmp = fmt.Sprintf("%s%s", tmp, value.(*scalar.EncapsedStringPart).Value)
			}
		}
		return tmp, nil
	case *scalar.Encapsed:
		tmp := ""
		for _, value := range n.(*scalar.Encapsed).Parts {
			switch value.(type) {
			case *VarStr:
				tmp = tmp + value.(*VarStr).String()
			case *scalar.EncapsedStringPart:
				tmp = fmt.Sprintf("%s%s", tmp, value.(*scalar.EncapsedStringPart).Value)
			}
		}
		return tmp, nil
	}
	return
}

func ZlibCompress(src []byte) (data string, err error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err = w.Write(src)
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	return in.String(), nil
}
