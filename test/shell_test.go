package test

import (
	"github.com/blue-bird1/ConfusedPHP"
	"github.com/z7zmey/php-parser/node"
	"testing"
)

func TestNewShell(t *testing.T) {
	shell := confusedPHP.NewShell("./test.php")
	err := shell.Parser()
	if err != nil {
		t.Error(err)
	}
	shell = confusedPHP.NewShell("./fake.php")
	err = shell.Parser()
	if err == nil {
		t.Error("error parser")
	}
}

func TestGetRootNode(t *testing.T) {
	shell := confusedPHP.NewShell("./test.php")
	err := shell.Parser()
	if err != nil {
		t.Error(err)
	}
	n := shell.GetRoot()
	n, ok := n.(*node.Root)
	if !ok {
		t.Error("error get root node")
	}
}
