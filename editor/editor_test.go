package editor

import (
	"github.com/blue-bird1/ConfusedPHP/phpread"
	"github.com/z7zmey/php-parser/printer"
	"os"
	"testing"
)

const TestStrCode = "<?php $a = \"test\";"

func TestBase64Editor_Edit(t *testing.T) {

	f, err := phpread.NewPhpString(TestStrCode)
	if err != nil {
		t.Error(err)
		return
	}
	err = f.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	editor := NewBase64Editor(BaseEditor{rootNode: f.GetRootNode()})
	err = editor.Edit()
	if err != nil {
		t.Error(err)
		return
	}
	file := os.Stdout
	p := printer.NewPrinter(file)
	p.Print(editor.rootNode)
}
