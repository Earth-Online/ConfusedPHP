package varProcess

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"testing"
)

var testVar = expr.NewVariable(&node.Identifier{Value: "test"})

func TestNewVarStr(t *testing.T) {
	NewVarStr(testVar)
}

func TestVarStr_Name(t *testing.T) {
	v := NewVarStr(testVar)
	if v.Name() != "test" {
		t.Error()
	}
}

func TestVarStr_Len(t *testing.T) {
	v := NewVarStr(testVar)
	if v.Len() != 4 {
		t.Error()
	}
}

func TestVarStr_String(t *testing.T) {
	v := NewVarStr(testVar)
	if v.String() != "$test" {
		t.Error()
	}
}
