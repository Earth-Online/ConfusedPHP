package obfuscator

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

var UserCallObfuscator = nodeProcess.NewFunctionCallPrecess("call_user_func", func(n *node.Node) (app []node.Node, cur node.Node) {
	nn := (*n).(*expr.FunctionCall)
	var nameNode node.Node
	nameNode = node.NewIdentifier("call_user_func")
	call := util.GetFunctionCall(nameNode, util.GetFunctionArg(append([]node.Node{nn.Function}, nn.ArgumentList.Arguments...)...).(*node.ArgumentList))
	return []node.Node{}, call
})
