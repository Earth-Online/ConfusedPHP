package confusedPHP

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
)

func GetAssign(name string, value node.Node) node.Node {
	var n node.Node
	n = &stmt.Expression{
		Expr: &assign.Assign{
			Variable: &expr.Variable{
				VarName: &node.Identifier{
					Value: name,
				},
			},
			Expression: value,
		},
	}
	return n
}

func GetFunctionCall(name node.Node, args *node.ArgumentList) (n node.Node) {
	n = &expr.FunctionCall{
		Function:     name,
		ArgumentList: args,
	}
	return
}

func GetFunctionRet(funcName string, ret node.Node) (n node.Node) {
	n = &stmt.Function{
		FunctionName: &name.Name{
			Parts: []node.Node{&name.NamePart{funcName}},
		},
		Stmts: []node.Node{
			&stmt.Return{
				Expr: ret,
			},
		},
	}
	return
}

func GetFunctionArg(nodes ...node.Node) (n node.Node) {
	var args []node.Node
	for _, value := range nodes {
		args = append(args, node.NewArgument(
			value,
			false,
			false))
	}
	n = node.NewArgumentList(args)
	return
}

func GetClass(name string, stmts []node.Node) (n node.Node) {
	n = &stmt.Class{
		PhpDocComment: "",
		ClassName:     &node.Identifier{Value: name},
		Stmts:         stmts,
	}
	return
}

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
