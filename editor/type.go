package editor

import "github.com/z7zmey/php-parser/node"

type Editor interface {
	RootNode() node.Node
	Edit() error
}
