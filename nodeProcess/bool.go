package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/z7zmey/php-parser/node"
)

type BoolProcess struct {
	BasePrecess
}

func NewBoolProcess(name string, f func(n *node.Node) ([]node.Node, node.Node)) *BoolProcess {
	precess := &BoolProcess{
		BasePrecess: BasePrecess{name: name},
	}
	precess.SetPrecess(f)
	return precess
}

func (b BoolProcess) Check(n *node.Node, preNode *node.Node) bool {
	return nodetype.IsRetBoolType(*n)
}
