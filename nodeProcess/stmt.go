package nodeProcess

import (
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
)

type StmtPrecess struct {
	BasePrecess
}

func (f StmtPrecess) Check(n node.Node, preNode util.EnterNode) bool {
	if preNode.Key != "Stmts" {
		return false
	}
	return true
}

func NewStmtPrecess(name string, f func(n node.Node) ([]node.Node, node.Node)) *StmtPrecess {
	precess := &StmtPrecess{
		BasePrecess: BasePrecess{
			name: name,
			precess: func(n node.Node) (append []node.Node, replace node.Node) {
				return f(n)
			},
		},
	}
	return precess
}
