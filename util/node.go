package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
)

//GetFunctionArg get a arg list
func GetFunctionArg(nodes ...node.Node) (n node.Node) {
	var args []node.Node
	for _, value := range nodes {
		args = append(args, node.NewArgument(
			value,
			false,
			false,
		))
	}
	n = node.NewArgumentList(args)
	return
}

// GetFunctionCall get function call. E.g. base64("123")
func GetFunctionCall(name node.Node, args *node.ArgumentList) (n node.Node) {
	n = &expr.FunctionCall{
		Function:     name,
		ArgumentList: args,
	}
	return
}

// GetFunctionRet get a only ret value function.E.g function a(){ return 1;}
func GetFunctionRet(funcName string, ret node.Node) (n node.Node) {
	n = &stmt.Function{
		FunctionName: &name.Name{
			Parts: []node.Node{&name.NamePart{Value: funcName}},
		},
		Stmts: []node.Node{
			&stmt.Return{
				Expr: ret,
			},
		},
	}
	return
}

// GetClass get a php class. E.g class a {}
func GetClass(name string, stmts []node.Node) (n node.Node) {
	n = &stmt.Class{
		PhpDocComment: "",
		ClassName:     &node.Identifier{Value: name},
		Stmts:         stmts,
	}
	return
}

// GetStaticPropertyFetch get a php class static fetch e.g foo::abc
func GetStaticPropertyFetch(className string, varName string) (n node.Node) {
	n = &expr.StaticPropertyFetch{
		Class: &name.Name{
			Parts: []node.Node{
				&name.NamePart{Value: className},
			},
		},
		Property: &expr.Variable{
			VarName: &node.Identifier{
				Value: varName,
			},
		},
	}
	return
}

// GetIfTrueExec get true if E.g if(1){ echo 1}
func GetIfTrueExec(stmts node.Node, elseif []node.Node, Else node.Node) (n node.Node) {
	n = &stmt.If{
		Cond: &scalar.Lnumber{
			Value: "1",
		},
		Stmt:   stmts,
		ElseIf: elseif,
		Else:   Else,
	}
	return
}

// GetArrayFetch get a array and fetch it.
func GetArrayFetch(array node.Node, fetch node.Node) (n node.Node) {
	n = &expr.ArrayDimFetch{
		Variable: array,
		Dim:      fetch,
	}
	return
}

// GetArray get a Array
func GetArray(value ...node.Node) (n node.Node) {
	n = &expr.Array{
		Items: value,
	}
	return
}

// GetStaticCall get a class static call
func GetStaticCall(className string, funcName string, args node.ArgumentList) (n node.Node) {
	n = &expr.StaticCall{
		Class: &name.Name{
			Parts: []node.Node{
				&name.NamePart{Value: className},
			},
		},
		Call:         &node.Identifier{Value: funcName},
		ArgumentList: &args,
	}
	return
}
