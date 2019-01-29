package test

import (
	"github.com/blue-bird1/ConfusedPHP"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"reflect"
	"testing"
)

func TestRandString(t *testing.T) {
	ret := confusedPHP.RandStringBytes(5)
	if len(ret) != 5 {
		t.Error("Error generating  string")
	}
}

func TestNodeIsType(t *testing.T) {
	testType := []reflect.Type{
		reflect.TypeOf(&scalar.String{}),
	}
	testString := scalar.NewString("test")
	if !confusedPHP.NodeIsType(testString, testType) {
		t.Error("isType check error")
	}
	testInt := scalar.NewDnumber("1")
	if confusedPHP.NodeIsType(testInt, testType) {
		t.Error("isType check error")
	}

}

func TestIsDefinitionType(t *testing.T) {
	var function node.Node = &stmt.Function{}
	if !confusedPHP.IsDefinitionType(function) {
		t.Error("function check error")
	}
	var class node.Node = &stmt.Class{}
	if !confusedPHP.IsDefinitionType(class) {
		t.Error("class check error")
	}
	var classMethod node.Node = &stmt.ClassMethod{}
	if !confusedPHP.IsDefinitionType(classMethod) {
		t.Error("classMethod check error")
	}
	var Interface node.Node = &stmt.Interface{}
	if !confusedPHP.IsDefinitionType(Interface) {
		t.Error("interface check error")
	}

	var trait node.Node = &stmt.Trait{}
	if !confusedPHP.IsDefinitionType(trait) {
		t.Error("trait check error")
	}
}

func TestIsConstantType(t *testing.T) {
	var str node.Node = &scalar.String{}
	if !confusedPHP.IsConstantType(str) {
		t.Error("sting constant check error")
	}
	var Int node.Node = &scalar.Dnumber{}
	if !confusedPHP.IsConstantType(Int) {
		t.Error("Int constant check error")
	}

	var Float node.Node = &scalar.Lnumber{}
	if !confusedPHP.IsConstantType(Float) {
		t.Error("float constant check error")
	}
}

func TestZlibCompress(t *testing.T) {
	test := []byte("123456")
	data, err := confusedPHP.ZlibCompress(test)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestIsFullyStringType(t *testing.T) {
	testString := scalar.NewString("test")
	if !confusedPHP.IsFullyStringType(testString) {
		t.Error()
	}
	testHereDoc := scalar.NewHeredoc("test", []node.Node{scalar.NewEncapsedStringPart("test")})
	if !confusedPHP.IsFullyStringType(testHereDoc) {
		t.Error()
	}
	testHereDoc = scalar.NewHeredoc("test", []node.Node{expr.NewVariable(node.NewIdentifier("test"))})
	if confusedPHP.IsFullyStringType(testHereDoc) {
		t.Error()
	}

}
