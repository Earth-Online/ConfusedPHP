package confusedPHP

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"strconv"
)

func FunctionRet(root node.Node, n *node.Node) (err error) {
	if IsDefinitionType(*n) {
		return errors.New("not support Definition Type ")
	}
	rootNode := root.(*node.Root)
	name := RandStringBytes(5)
	nn := GetFunctionRet(name, *n)
	rootNode.Stmts = append([]node.Node{nn}, rootNode.Stmts...)
	var nameNode node.Node
	nameNode = node.NewIdentifier(name)
	ret := GetFunctionCall(nameNode, node.NewArgumentList([]node.Node{}))
	*n = ret
	return
}

func StringSplit(root node.Node, n *node.Node) (err error) {
	nn, ok := (*n).(*scalar.String)
	if !ok {
		return errors.New("")
	}
	value := nn.Value
	split := len(value) / 2
	string1 := scalar.NewString(fmt.Sprintf("%s\"", value[:split]))
	string2 := scalar.NewString(fmt.Sprintf("\"%s", value[split:]))
	t := binary.NewPlus(string1, string2)
	*n = t
	return
}

func StringBase64(root node.Node, n *node.Node) (err error) {
	nn, ok := (*n).(*scalar.String)
	if !ok {
		return errors.New("")
	}
	value := fmt.Sprintf("\"%s\"", base64.StdEncoding.EncodeToString([]byte(nn.Value)))
	var nameNode node.Node
	nameNode = node.NewIdentifier("base64")
	args := GetFunctionArg(scalar.NewString(value))
	f := GetFunctionCall(nameNode, args.(*node.ArgumentList))
	*n = f
	return
}

func IntSplit(root node.Node, n *node.Node) (err error) {
	switch (*n).(type) {
	case *scalar.Dnumber:
		i, err := strconv.Atoi((*n).(*scalar.Dnumber).Value)
		if err != nil {
			return err
		}
		value := i / 2
		value2 := i - value
		var valueNode node.Node = scalar.NewDnumber(strconv.Itoa(value))
		var value2Node node.Node = scalar.NewDnumber(strconv.Itoa(value2))
		var nn node.Node = binary.NewPlus(valueNode, value2Node)
		n = &nn
	case *scalar.Lnumber:
		v := (*n).(*scalar.Lnumber).Value
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		value := f - 1
		value2 := f - value
		var valueNode node.Node = scalar.NewLnumber(strconv.FormatFloat(value, 'f', -1, 64))
		var value2Node node.Node = scalar.NewLnumber(strconv.FormatFloat(value2, 'f', -1, 64))
		var nn node.Node = binary.NewPlus(valueNode, value2Node)
		n = &nn
	default:
		return errors.New("")
	}
	return
}

func ValueEqual(root node.Node, n *node.Node) (err error) {
	assign := GetAssign(RandStringBytes(5), *n)
	*n = assign
	return
}

func FunctionUserCall(root node.Node, n *node.Node) (err error) {
	nn, ok := (*n).(*expr.FunctionCall)
	if !ok {
		return errors.New("")
	}
	var nameNode node.Node
	nameNode = node.NewIdentifier("call_user_func")
	call := GetFunctionCall(nameNode, GetFunctionArg(append([]node.Node{nn.Function}, nn.ArgumentList.Arguments...)...).(*node.ArgumentList))
	*n = call
	return
}

func BoolTwoNot(root node.Node, n *node.Node) (err error) {
	var nn node.Node
	switch (*n).(type) {
	case *expr.BitwiseNot:
		nn = &expr.BitwiseNot{
			Expr: &expr.BitwiseNot{
				Expr: *n,
			},
		}
	case *expr.BooleanNot:
		nn = &expr.BooleanNot{
			Expr: &expr.BooleanNot{
				Expr: *n,
			},
		}
	default:
		return errors.New("")
	}
	*n = nn
	return
}

func ClassStaticAttr(root node.Node, n *node.Node) (err error) {
	if IsDefinitionType(*n) {
		return errors.New("only supoort var")
	}
	attrName := RandStringBytes(20)
	var stmts node.Node = &stmt.PropertyList{
		Modifiers: []node.Node{
			&node.Identifier{
				Value: "public",
			},
			&node.Identifier{
				Value: "static",
			},
		},
		Properties: []node.Node{&stmt.Property{
			PhpDocComment: "",
			Variable: &expr.Variable{
				VarName: &node.Identifier{
					attrName,
				},
			},
			Expr: *n,
		}},
	}
	className := RandStringBytes(20)
	nn := GetClass(className, []node.Node{stmts})
	rootNode := root.(*node.Root)
	rootNode.Stmts = append([]node.Node{nn}, rootNode.Stmts...)
	*n = GetStaticPropertyFetch(className, attrName)
	return
}

func ArrayFetch(root node.Node, n *node.Node) (err error) {
	if !IsHaveReturnType(*n) {
		return errors.New("only support have return")
	}
	nn := GetArrayFetch(
		GetArray(&expr.ArrayItem{Val: *n}),
		&scalar.String{Value: "1"},
	)
	*n = nn
	return
}

var FunctionList = []func(root node.Node, n *node.Node) (err error){
	FunctionRet,
	StringSplit,
	StringBase64,
	IntSplit,
	//	ValueEqual,
	FunctionUserCall,
	BoolTwoNot,
	ClassStaticAttr,
	ArrayFetch,
}
