package nodeProcess

import (
	"github.com/z7zmey/php-parser/node"
)

type NodePrecess interface {
	Precess(n *node.Node) ([]node.Node, node.Node)
	Check(n *node.Node, preNode *node.Node) bool
	Name() string
}

type BasePrecess struct {
	precess func(n *node.Node) (append []node.Node, replace node.Node)
	name    string
}

func (i BasePrecess) SetName(name string) {
	i.name = name
}

func (i BasePrecess) SetPrecess(precess func(n *node.Node) ([]node.Node, node.Node)) {
	i.precess = precess
}

func (i BasePrecess) Check(n *node.Node, preNode *node.Node) bool {
	panic("implement me")
}

func (i BasePrecess) Precess(n *node.Node) ([]node.Node, node.Node) {
	return i.precess(n)
}

func (i BasePrecess) Name() string {
	return i.name
}
