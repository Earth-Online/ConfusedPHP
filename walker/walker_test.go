package walker

import (
	"github.com/blue-bird1/ConfusedPHP/phpread"
	"github.com/z7zmey/php-parser/visitor"
	"os"
	"testing"
)

func Test(t *testing.T) {
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
	nsResolver := visitor.NewNamespaceResolver()
	root.Walk(nsResolver)
	dumper := &visitor.Dumper{
		Writer:     os.Stdout,
		Indent:     "| ",
		NsResolver: nsResolver,
	}
	root.Walk(dumper)
	t.Error()
}
