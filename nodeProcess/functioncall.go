package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

type FunctionCallPrecess struct {
	BasePrecess
}

func (f FunctionCallPrecess) Check(n node.Node, preNode util.EnterNode) bool {
	_, ok := (n).(*expr.FunctionCall)
	return ok
}

func NewFunctionCallPrecess(name string, f func(n node.Node) ([]node.Node, node.Node)) *FunctionCallPrecess {
	precess := &FunctionCallPrecess{
		BasePrecess: BasePrecess{
			name: name,
			precess: func(n node.Node) (append []node.Node, replace node.Node) {
				return f(n)
			},
		},
	}
	return precess
}
