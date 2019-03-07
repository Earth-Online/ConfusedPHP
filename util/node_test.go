package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"testing"
)

const TESTStr = "test"

var str = &scalar.String{Value: "test"}
var item = &expr.ArrayItem{Val: str}

func TestGetArray(t *testing.T) {
	array := GetArray(item)
	v, ok := array.(*expr.Array)
	if !ok {
		t.Error("getArray type error")
		return
	}
	if len(v.Items) != 1 {
		t.Error("getArray len error")
	}
	if v.Items[0] != item {
		t.Error("getArray item error")
	}

}

func TestGetClass(t *testing.T) {

	class := GetClass(TESTStr, []node.Node{})
	v, ok := class.(*stmt.Class)
	if !ok {
		t.Error("getClass type error")
		return
	}
	if v.PhpDocComment != "" {
		t.Error()
	}
	if len(v.Stmts) != 0 {
		t.Error()
	}
}

func TestGetFunctionArg(t *testing.T) {
	call := GetFunctionArg(str)
	v, ok := call.(*node.ArgumentList)
	if !ok {
		t.Error()
		return
	}
	if len(v.Arguments) != 1 {
		t.Error()
		return
	}
}

func TestGetFunctionCall(t *testing.T) {
	call := GetFunctionCall(str, &node.ArgumentList{})
	v, ok := call.(*expr.FunctionCall)
	if !ok {
		t.Error()
		return
	}
	if len(v.ArgumentList.Arguments) != 0 {
		t.Error()
	}

}
