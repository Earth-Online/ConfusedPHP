package nodetype

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"reflect"
)

// php definition E.g function
var DefinitionType = []reflect.Type{
	reflect.TypeOf(&stmt.Class{}),
	reflect.TypeOf(&stmt.Function{}),
	reflect.TypeOf(&stmt.ClassMethod{}),
	reflect.TypeOf(&stmt.Interface{}),
	reflect.TypeOf(&stmt.Trait{}),
}

var ProcessControlType = append([]reflect.Type{
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
}, AltProcessControlType...)

var AltProcessControlType = []reflect.Type{
	reflect.TypeOf(&stmt.AltIf{}),
	reflect.TypeOf(&stmt.AltElseIf{}),
	reflect.TypeOf(&stmt.AltElse{}),
	reflect.TypeOf(&stmt.AltSwitch{}),
	reflect.TypeOf(&stmt.AltFor{}),
	reflect.TypeOf(&stmt.AltForeach{}),
	reflect.TypeOf(&stmt.AltWhile{}),
}

var HaveReturnType = append([]reflect.Type{
	reflect.TypeOf(&expr.FunctionCall{}),
	reflect.TypeOf(&expr.ShortArray{}),
	reflect.TypeOf(&expr.Array{}),
	reflect.TypeOf(&expr.Variable{}),
	reflect.TypeOf(&expr.InstanceOf{}),
	reflect.TypeOf(&expr.Eval{}),
	reflect.TypeOf(&expr.List{}),
	reflect.TypeOf(&expr.Empty{}),
	reflect.TypeOf(&expr.Isset{}),
	reflect.TypeOf(&expr.ShellExec{}),
	reflect.TypeOf(&expr.Print{}),
	reflect.TypeOf(&expr.StaticPropertyFetch{}),
	reflect.TypeOf(&expr.StaticCall{}),
}, ConstantType...)

var NumType = []reflect.Type{
	reflect.TypeOf(&scalar.Lnumber{}),
	reflect.TypeOf(&scalar.Dnumber{}),
}

var StringType = []reflect.Type{
	reflect.TypeOf(&scalar.String{}),
	reflect.TypeOf(&scalar.Encapsed{}),
	reflect.TypeOf(&scalar.Heredoc{}),
}

var BoolType = []reflect.Type{
	reflect.TypeOf(&expr.BooleanNot{}),
	reflect.TypeOf(&expr.BitwiseNot{}),
}
var RetBoolType = append([]reflect.Type{
	reflect.TypeOf(&expr.Isset{}),
	reflect.TypeOf(&expr.Empty{}),
}, BoolType...)

var ConstantType = append(StringType, NumType...)

var ListType = []reflect.Type{
	reflect.TypeOf(&expr.List{}),
	reflect.TypeOf(&expr.ShortList{}),
	reflect.TypeOf(&expr.Array{}),
	reflect.TypeOf(&expr.ShortArray{}),
}

var ValueType = append(ConstantType, ListType...)

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

var IsDefinitionType = NodeIsType(DefinitionType)

var IsHaveReturnType = NodeIsType(HaveReturnType)

var IsConstantType = NodeIsType(ConstantType)

var IsProcessControlType = NodeIsType(ProcessControlType)

var IsStringType = NodeIsType(StringType)

var IsValueType = NodeIsType(ValueType)

var IsBoolType = NodeIsType(BoolType)

var IsRetBoolType = NodeIsType(RetBoolType)
