package editor

import (
	"fmt"
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/obfuscator"
	"github.com/blue-bird1/ConfusedPHP/phpread"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/printer"
	"os"
	"testing"
)

func TestEditWalker(t *testing.T) {
	editor := NewEditWalker([]nodeProcess.NodePrecess{obfuscator.FunctionRetObfuscator})
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
	fmt.Printf("%v+", editor.addNode)
	p2 := printer.NewPrinter(os.Stdout)
	p2.Print(node.NewRoot(editor.addNode))
	fmt.Print("?>")
	p := NewPrinter(os.Stdout, editor.modifyNode)
	p.Print(root)
}
