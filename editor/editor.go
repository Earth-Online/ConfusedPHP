package editor

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
)

type BaseEditor struct {
	rootNode    node.Node
	currentNode node.Node
}

func (b *BaseEditor) CurrentNode() node.Node {
	return b.currentNode
}

func (b *BaseEditor) SetCurrentNode(currentNode node.Node) {
	b.currentNode = currentNode
}

func (b *BaseEditor) RootNode() node.Node {
	return b.rootNode
}

func (b *BaseEditor) SetRootNode(rootNode node.Node) {
	b.rootNode = rootNode
}

func (b *BaseEditor) Edit() error {
	b.getDeepNode(b.rootNode)
	return nil
}

func (b *BaseEditor) EditNode(n node.Node) error {
	panic("implement me")
}

func (b *BaseEditor) confuse(node node.Node, precess nodeProcess.NodePrecess) (err error) {
	appendNode, cur := precess.Precess(&node)
	if len(appendNode) != 0 {
		err = b.appendNode(appendNode, b.currentNode.GetPosition())
		if err != nil {
			return
		}
	}
	if cur != nil {
		cur.SetPosition(b.currentNode.GetPosition())
		b.SetCurrentNode(cur)
	}
	return
}

func (b *BaseEditor) appendNode(n []node.Node, position *position.Position) (err error) {
	root := b.rootNode.(*node.Root)
	root.Stmts = append(n, root.Stmts...)
	return
}

func (b *BaseEditor) getDeepNode(n node.Node) {
	tmp := b.currentNode
	b.currentNode = n
	ok, value := nodetype.IsHaveValueType(n)
	if ok {
		b.getDeepNode(value.(node.Node))
	}
	ok, value = nodetype.IsHaveExprType(n)
	if ok {
		b.getDeepNode(value.(node.Node))
	}
	ok, value = nodetype.IsHaveStmtsType(n)
	if ok {
		for _, v := range value.([]node.Node) {
			b.getDeepNode(v)
		}
	}
	ok, value = nodetype.IsHavePartsType(n)
	if ok {
		for _, v := range value.([]node.Node) {
			b.getDeepNode(v)
		}
	}

	_ = b.EditNode(n)
	b.currentNode = tmp
}
