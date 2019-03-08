package obfuscator

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

// TwoNotObfuscator add !! at a bool value
var TwoNotObfuscator = nodeProcess.NewBoolProcess("two not ", func(n *node.Node) (app []node.Node, cur node.Node) {
	var nn node.Node
	nn = &expr.BooleanNot{
		Expr: &expr.BooleanNot{
			Expr: *n,
		},
	}
	return []node.Node{}, nn
})
