package test

import (
	"bytes"
	"fmt"
	"github.com/blue-bird1/ConfusedPHP"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/visitor"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

var shell = confusedPHP.NewShell("./test.php")
var _ = shell.Parser()

func TestNewEditor(t *testing.T) {
	n := shell.GetRoot()
	confusedPHP.NewEditor(&n)
}

func TestEdit(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	n := shell.GetRoot()
	var editor = confusedPHP.NewEditor(&n)
	err := editor.Edit()
	err = editor.Edit()
	err = editor.Edit()
	if err != nil {
		t.Error(err)
	}
	p.Print(*editor.Root)
}

var visit = visitor.Dumper{
	Writer: os.Stdout,
	Indent: "",
	//Comments:  parser.GetComments(),
	//Positions: parser.GetPositions(),
}

var p = printer.NewPrinter(os.Stdout, "    ")

func TestPrint(t *testing.T) {
	//file := os.Stdout
	//p := printer.NewPrinter(file, "    ")
	// var n node.Node
	//	n = node.NewNullable(scalar.NewString("123"))
	//	 p.Print(n)
	src := `<?php echo array("1", "2")[1];`
	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	root := php7parser.GetRootNode()
	root.Walk(visit)
	p.Print(root)
}

func TestReflect(t *testing.T) {
	fmt.Println(reflect.TypeOf(&stmt.Class{}))
	fmt.Println(reflect.TypeOf(confusedPHP.GetClass("test", []node.Node{})))
}
