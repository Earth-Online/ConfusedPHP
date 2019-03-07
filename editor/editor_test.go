package editor

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/obfuscator"
	"github.com/blue-bird1/ConfusedPHP/phpread"
	"os"
	"testing"
)

func TestEditWalker(t *testing.T) {
	editor := NewEditWalker([]nodeProcess.NodePrecess{obfuscator.Base64Obfuscator})
	testCode := `
	<?php
	 		
echo "hello world";
$a=1; 
	`
	parser, err := phpread.NewPhpString(testCode)
	if err != nil {
		t.Error(err)
		return
	}
	err = parser.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	root := parser.GetRootNode()
	root.Walk(editor)
	p := NewPrinter(os.Stdout, editor.modifyNode)
	p.Print(root)
}
