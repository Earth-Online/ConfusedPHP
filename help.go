package confusedPHP

import (
	"bytes"
	"compress/zlib"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"math/rand"
	"reflect"
	"time"
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

func RandConfused() bool {
	rand.Seed(int64(time.Now().UnixNano()))
	if rand.Int()%2 == 1 {
		return true
	}
	return false
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

var ConstantType = []reflect.Type{
	reflect.TypeOf(&scalar.String{}),
	reflect.TypeOf(&scalar.Lnumber{}),
	reflect.TypeOf(&scalar.Dnumber{}),
	reflect.TypeOf(&scalar.Encapsed{}),
	reflect.TypeOf(&scalar.Heredoc{}),
}

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
