package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

// get a arg list
func GetFunctionArg(nodes ...node.Node) (n node.Node) {
	var args []node.Node
	for _, value := range nodes {
		args = append(args, node.NewArgument(
			value,
			false,
			false))
	}
	n = node.NewArgumentList(args)
	return
}

// get function call. E.g. base64("123")
func GetFunctionCall(name node.Node, args *node.ArgumentList) (n node.Node) {
	n = &expr.FunctionCall{
		Function:     name,
		ArgumentList: args,
	}
	return
}
