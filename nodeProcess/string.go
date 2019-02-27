package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/blue-bird1/ConfusedPHP/varProcess"
	"github.com/z7zmey/php-parser/node"
)

type StringPrecess struct {
	BasePrecess
}

func NewStringPrecess(name string, f func(n string) ([]node.Node, node.Node)) *StringPrecess {
	precess := &StringPrecess{
		BasePrecess: BasePrecess{
			name: name,
		},
	}
	precess.SetPrecess(func(n *node.Node) ([]node.Node, node.Node) {
		s := precess.GetString(n)
		return f(s)
	})
	return precess
}

func (s StringPrecess) Check(n *node.Node, preNode *node.Node) bool {
	return nodetype.IsStringType(*n)
}

func (s StringPrecess) GetString(n *node.Node) string {
	str, _ := varProcess.GetStingTypeValue(*n)
	return str
}
