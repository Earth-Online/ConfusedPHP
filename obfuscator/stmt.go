package obfuscator

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
)

var IfTrueExecObfuscator = nodeProcess.NewStmtPrecess("iftrue", func(n node.Node) (nodes []node.Node, cur node.Node) {
	if !nodetype.IsHaveReturnType(n) {
		return
	}
	nn := util.GetIfTrueExec(&stmt.StmtList{Stmts: []node.Node{n}}, nil, nil)
	return []node.Node{}, nn
})
