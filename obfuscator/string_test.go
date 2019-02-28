package obfuscator

import (
	"encoding/base64"
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"testing"
)

var str node.Node = &scalar.String{Value: "test"}

func TestBase64Obfuscator(t *testing.T) {
	app, cur := Base64Obfuscator.Precess(&str)
	if len(app) != 0 {
		t.Error()
	}
	v, ok := cur.(*expr.FunctionCall)
	if !ok {
		t.Error()
		return
	}
	name, ok := v.Function.(*node.Identifier)
	if !ok {
		t.Error()
		return
	}
	if name.Value != "base64" {
		t.Error("error function name")
	}
	if len(v.ArgumentList.Arguments) != 1 {
		t.Error("error arg len")
		return

	}
	base64str := v.ArgumentList.Arguments[0]
	value, ok := base64str.(*node.Argument)
	if !ok {
		t.Error("arg error")
		return
	}
	strValue, ok := value.Expr.(*scalar.String)
	if !ok {
		t.Error("arg error")
		return
	}
	if strValue.Value != fmt.Sprintf("\"%s\"", base64.StdEncoding.EncodeToString([]byte("test"))) {
		t.Error("error base64")
	}
}
