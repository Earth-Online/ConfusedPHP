package nodeProcess

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"testing"
)

const TESTStr = "test"

var str node.Node = &scalar.String{Value: TESTStr}

func TestStringPrecess(t *testing.T) {
	testProcess := NewStringPrecess(TESTStr, func(n string) (nodes []node.Node, node node.Node) {
		return
	})
	if testProcess.Name() != TESTStr {
		t.Error("error name")
	}
	app, cur := testProcess.Precess(&str)
	if len(app) != 0 || cur != nil {
		t.Error("error process")
	}
	if !testProcess.Check(&str, &str) {
		t.Error("error check")
	}
}
