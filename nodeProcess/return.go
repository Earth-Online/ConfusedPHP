package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/z7zmey/php-parser/node"
)

type ReturnProcess struct {
	BasePrecess
}

func NewReturnProcess(name string, f func(n node.Node) ([]node.Node, node.Node)) *ReturnProcess {
	precess := &ReturnProcess{
		BasePrecess: BasePrecess{
			name: name,
			precess: func(n node.Node) (append []node.Node, replace node.Node) {
				return f(n)
			},
		},
	}
	return precess
}

func (b ReturnProcess) Check(n node.Node, preNode node.Node) bool {
	return nodetype.IsHaveReturnType(n)
}
