package obfuscator

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
)

//  ArrayFetchObfuscator Conversion to array and fetch it
var ArrayFetchObfuscator = nodeProcess.NewReturnProcess("arrayFetch", func(n *node.Node) (app []node.Node, cur node.Node) {
	nn := util.GetArrayFetch(
		util.GetArray(&expr.ArrayItem{Val: *n}),
		&scalar.String{Value: "0"},
	)
	return []node.Node{}, nn
})
