package editor

import "github.com/z7zmey/php-parser/node"

type BaseEditor struct {
	rootNode node.Node
}

func (b *BaseEditor) RootNode() node.Node {
	return b.rootNode
}

func (b *BaseEditor) SetRootNode(rootNode node.Node) {
	b.rootNode = rootNode
}

func (b *BaseEditor) Edit() error {
	panic("implement me")
}
