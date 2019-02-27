package nodetype

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"reflect"
	"testing"
)

func TestNodeIsType(t *testing.T) {
	testType := []reflect.Type{
		reflect.TypeOf(&scalar.String{}),
	}
	testString := scalar.NewString("test")
	if !NodeIsType(testType)(testString) {
		t.Error("isType check error")
	}
	testInt := scalar.NewDnumber("1")
	if NodeIsType(testType)(testInt) {
		t.Error("isType check error")
	}

}

func TestIsDefinitionType(t *testing.T) {
	var function node.Node = &stmt.Function{}
	if !IsDefinitionType(function) {
		t.Error("function check error")
	}
	var class node.Node = &stmt.Class{}
	if !IsDefinitionType(class) {
		t.Error("class check error")
	}
	var classMethod node.Node = &stmt.ClassMethod{}
	if !IsDefinitionType(classMethod) {
		t.Error("classMethod check error")
	}
	var Interface node.Node = &stmt.Interface{}
	if !IsDefinitionType(Interface) {
		t.Error("interface check error")
	}

	var trait node.Node = &stmt.Trait{}
	if !IsDefinitionType(trait) {
		t.Error("trait check error")
	}
}

func TestIsConstantType(t *testing.T) {
	var str node.Node = &scalar.String{}
	if !IsConstantType(str) {
		t.Error("sting constant check error")
	}
	var Int node.Node = &scalar.Dnumber{}
	if !IsConstantType(Int) {
		t.Error("Int constant check error")
	}

	var Float node.Node = &scalar.Lnumber{}
	if !IsConstantType(Float) {
		t.Error("float constant check error")
	}
}

func TestIsBoolType(t *testing.T) {
	var not = &expr.BooleanNot{}
	var bitBool = &expr.BitwiseNot{}
	if !IsBoolType(not) || !IsBoolType(bitBool) {
		t.Error("error bool type check")
	}
}
