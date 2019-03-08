package nodetype

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"reflect"
)

var definitionType = []reflect.Type{
	reflect.TypeOf(&stmt.Class{}),
	reflect.TypeOf(&stmt.Function{}),
	reflect.TypeOf(&stmt.ClassMethod{}),
	reflect.TypeOf(&stmt.Interface{}),
	reflect.TypeOf(&stmt.Trait{}),
}

var processControlType = append([]reflect.Type{
	reflect.TypeOf(&stmt.If{}),
	reflect.TypeOf(&stmt.Else{}),
	reflect.TypeOf(&stmt.ElseIf{}),
	reflect.TypeOf(&stmt.Switch{}),
	reflect.TypeOf(&stmt.Case{}),
	reflect.TypeOf(&stmt.Do{}),
	reflect.TypeOf(&stmt.While{}),
	reflect.TypeOf(&stmt.CaseList{}),
	reflect.TypeOf(&stmt.For{}),
	reflect.TypeOf(&stmt.Foreach{}),
}, altProcessControlType...)

var altProcessControlType = []reflect.Type{
	reflect.TypeOf(&stmt.AltIf{}),
	reflect.TypeOf(&stmt.AltElseIf{}),
	reflect.TypeOf(&stmt.AltElse{}),
	reflect.TypeOf(&stmt.AltSwitch{}),
	reflect.TypeOf(&stmt.AltFor{}),
	reflect.TypeOf(&stmt.AltForeach{}),
	reflect.TypeOf(&stmt.AltWhile{}),
}

var haveReturnType = append([]reflect.Type{
	reflect.TypeOf(&expr.FunctionCall{}),
	reflect.TypeOf(&expr.ShortArray{}),
	reflect.TypeOf(&expr.Array{}),
	//	reflect.TypeOf(&expr.Variable{}),
	reflect.TypeOf(&expr.InstanceOf{}),
	reflect.TypeOf(&expr.Eval{}),
	reflect.TypeOf(&expr.List{}),
	reflect.TypeOf(&expr.Empty{}),
	reflect.TypeOf(&expr.Isset{}),
	reflect.TypeOf(&expr.ShellExec{}),
	reflect.TypeOf(&expr.Print{}),
	reflect.TypeOf(&expr.StaticPropertyFetch{}),
	reflect.TypeOf(&expr.StaticCall{}),
	reflect.TypeOf(&stmt.Expression{}),
}, constantType...)

var numType = []reflect.Type{
	reflect.TypeOf(&scalar.Lnumber{}),
	reflect.TypeOf(&scalar.Dnumber{}),
}

var stringType = []reflect.Type{
	reflect.TypeOf(&scalar.String{}),
	reflect.TypeOf(&scalar.Encapsed{}),
	reflect.TypeOf(&scalar.Heredoc{}),
}

var boolType = []reflect.Type{
	reflect.TypeOf(&expr.BooleanNot{}),
	reflect.TypeOf(&expr.BitwiseNot{}),
}
var retBoolType = append([]reflect.Type{
	reflect.TypeOf(&expr.Isset{}),
	reflect.TypeOf(&expr.Empty{}),
}, boolType...)

var constantType = append(stringType, numType...)

var listType = []reflect.Type{
	reflect.TypeOf(&expr.List{}),
	reflect.TypeOf(&expr.ShortList{}),
	reflect.TypeOf(&expr.Array{}),
	reflect.TypeOf(&expr.ShortArray{}),
}

var valueType = append(constantType, listType...)

// NodeIsType determine if the node is in the list
func NodeIsType(nodeType []reflect.Type) func(n node.Node) bool {
	return func(n node.Node) bool {
		for _, val := range nodeType {
			if val == reflect.TypeOf(n) {
				return true
			}
		}
		return false
	}
}

var IsDefinitionType = NodeIsType(definitionType)

var IsHaveReturnType = NodeIsType(haveReturnType)

var IsConstantType = NodeIsType(constantType)

var IsProcessControlType = NodeIsType(processControlType)

var IsStringType = NodeIsType(stringType)

var IsValueType = NodeIsType(valueType)

var IsBoolType = NodeIsType(boolType)

var IsRetBoolType = NodeIsType(retBoolType)
