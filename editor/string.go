package editor

import (
	"github.com/blue-bird1/ConfusedPHP/nodetype"
	"github.com/blue-bird1/ConfusedPHP/obfuscator"
	"github.com/z7zmey/php-parser/node"
)

// only a hello world example
type Base64Editor struct {
	BaseEditor
}

func (b *Base64Editor) Edit() error {
	return nil
}

func (b *Base64Editor) EditNode(n node.Node) error {
	if nodetype.IsStringType(n) {
		return b.confuse(n, obfuscator.Base64Obfuscator)
	}
	return nil
}
