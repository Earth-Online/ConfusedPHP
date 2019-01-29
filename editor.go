package confusedPHP

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"math/rand"
)

type Editor struct {
	Root *node.Node
}

func NewEditor(root *node.Node) *Editor {
	return &Editor{Root: root}
}

func (e *Editor) Edit() error {
	return e.EditNode(e.Root)
}

func (e *Editor) EditNode(n *node.Node) (err error) {
	value := *n
	switch value.(type) {
	case *node.Root:
		return e.EditRoot(&value)
	case *node.Identifier:
		return nil
	//	return e.EditIdent(n)
	case *node.Argument:
		return e.EditNode(&(value.(*node.Argument).Expr))
	case *node.ArgumentList:
		return e.EditNodes(&(value.(*node.ArgumentList).Arguments))
	case *node.Nullable:
		return nil
	case *node.Parameter:
		return nil

	case *name.Name:
		return nil
	case *name.NamePart:
		return nil
	case *name.Relative:
		return nil
	case *name.FullyQualified:
		return nil

	case *expr.Eval:
		//err =  e.EditEval(n)
		err = e.EditNode(&(value.(*expr.Eval).Expr))
	case *expr.ShellExec:
		err = e.EditNodes(&(value.(*expr.ShellExec).Parts))
	case *expr.Print:
		err = e.EditNode(&(value.(*expr.Print).Expr))
	case *expr.Exit:
		err = e.EditNode(&value.(*expr.Exit).Expr)
	case *expr.Die:
		err = e.EditNode(&value.(*expr.Die).Expr)
	case *expr.UnaryMinus:
		err = e.EditNode(&value.(*expr.UnaryMinus).Expr)
	case *expr.UnaryPlus:
		err = e.EditNode(&value.(*expr.UnaryPlus).Expr)
	case *expr.Clone:
		err = e.EditNode(&value.(*expr.Clone).Expr)
	case *expr.ErrorSuppress:
		return e.EditNode(&value.(*expr.ErrorSuppress).Expr)
	case *expr.Isset:
		err = e.EditNodes(&value.(*expr.Isset).Variables)
	//case *expr.List:
	//	err =   e.EditNodes(n.(*expr.List).Items)

	//	case *expr.Variable:
	//	return e.EditVar(&n)
	case *expr.New:
		return nil
	case *expr.StaticCall:
		return nil

	case *expr.InstanceOf:
		err = e.EditNode(&value.(*expr.InstanceOf).Expr)

	case *expr.Ternary:
		return nil
	case *expr.StaticPropertyFetch:
		return nil
		// php array and list
	case *expr.Array:
		err = e.EditNodes(&value.(*expr.Array).Items)
	case *expr.ArrayItem:
		nn := value.(*expr.ArrayItem)
		err = e.EditNode(&nn.Val)
		if err != nil {
			return err
		}
		err = e.EditNode(&nn.Key)
		return
	case *expr.List:
		err = e.EditNodes(&value.(*expr.List).Items)
	case *expr.ShortList:
		err = e.EditNodes(&value.(*expr.ShortList).Items)
	case *expr.ShortArray:
		err = e.EditNodes(&value.(*expr.ShortArray).Items)

		// &1
	case *expr.BitwiseNot:
		err = e.EditNode(&value.(*expr.BitwiseNot).Expr)
		// !1
	case *expr.BooleanNot:
		err = e.EditNode(&value.(*expr.BooleanNot).Expr)

		// php include file
	case *expr.RequireOnce:
		return nil
	case *expr.Require:
		return nil
	case *expr.Include:
		return
	case *expr.IncludeOnce:
		return

	case *expr.Reference:
		return e.EditNode(&value.(*expr.Reference).Variable)
	case *expr.PropertyFetch:
		return nil

	case *expr.Closure:
		return nil
	case *expr.ClosureUse:
		return nil

	case *expr.Yield:
		return nil
	case *expr.YieldFrom:
		return nil

		// ++1
	case *expr.PreInc:
		return nil
		// --1
	case *expr.PreDec:
		return nil
		// 1++
	case *expr.PostInc:
		return nil
		// 1--
	case *expr.PostDec:
		return nil

	case *expr.Empty:
		return e.EditNode(&value.(*expr.Empty).Expr)
	case *expr.FunctionCall:
		nn := value.(*expr.FunctionCall)
		err = e.EditNode(&nn.Function)
		if err != nil {
			return err
		}
		var nnn node.Node = nn.ArgumentList
		err = e.EditNode(&nnn)
	case *expr.ArrayDimFetch:
		nn := value.(*expr.ArrayDimFetch)
		err = e.EditNode(&nn.Variable)
		if err != nil {
			return
		}
		return e.EditNode(&nn.Dim)

		// php __xxx__
	case *scalar.MagicConstant:
		return nil

	// php if
	case *stmt.If:
		err = e.EditNode(&value.(*stmt.If).Cond)
		if err != nil {
			return err
		}
		err = e.EditNode(&value.(*stmt.If).Stmt)
		if err != nil {
			return err
		}
		err = e.EditNodes(&value.(*stmt.If).ElseIf)
		if err != nil {
			return err
		}
		return e.EditNode(&value.(*stmt.If).Else)
	case *stmt.ElseIf:
		err = e.EditNode(&value.(*stmt.ElseIf).Cond)
		if err != nil {
			return err
		}
		return e.EditNode(&value.(*stmt.ElseIf).Stmt)
	case *stmt.Else:
		return e.EditNode(&value.(*stmt.Else).Stmt)

	// php switch
	case *stmt.Switch:
		err = e.EditNode(&value.(*stmt.Switch).Cond)
		if err != nil {
			return err
		}
		var nn node.Node = value.(*stmt.Switch).CaseList
		err = e.EditNode(&nn)
	case *stmt.CaseList:
		err = e.EditNodes(&value.(*stmt.CaseList).Cases)
		return
	case *stmt.Case:
		err = e.EditNode(&value.(*stmt.Case).Cond)
		if err != nil {
			return
		}
		err = e.EditNodes(&value.(*stmt.Case).Stmts)

	case *stmt.For:
		return nil
	case *stmt.Foreach:
		return nil

	case *stmt.Global:
		return nil
	case *stmt.Use:
		return nil
	case *stmt.UseList:
		return nil
	case *stmt.Finally:
		return nil

	case *stmt.While:
		err = e.EditNode(&value.(*stmt.While).Cond)
		if err != nil {
			return
		}
		return e.EditNode(&value.(*stmt.While).Stmt)
	case *stmt.Do:
		err = e.EditNode(&value.(*stmt.Do).Cond)
		if err != nil {
			return
		}
		return e.EditNode(&value.(*stmt.Do).Stmt)
	case *stmt.Expression:
		return e.EditNode(&value.(*stmt.Expression).Expr)
	case *stmt.Echo:
		err = e.EditNodes(&value.(*stmt.Echo).Exprs)
	case *stmt.Nop:
		return nil

	case *stmt.Try:
		err = e.EditNodes(&value.(*stmt.Try).Stmts)
		if err != nil {
			return
		}
		err = e.EditNodes(&value.(*stmt.Try).Catches)
		if err != nil {
			return
		}
		err = e.EditNode(&value.(*stmt.Try).Finally)
		return err
	case *stmt.Catch:
		err = e.EditNodes(&value.(*stmt.Catch).Stmts)
		return nil
	case *stmt.Throw:
		err = e.EditNode(&value.(*stmt.Throw).Expr)
		return err
	// php 7 not support break and continue  non-constant operand
	case *stmt.Break:
		return nil
	case *stmt.Continue:
		return nil

	case *stmt.AltIf:
		return nil
	case *stmt.AltElseIf:
		return nil
	case *stmt.AltElse:
		return nil
	case *stmt.AltWhile:
		return nil
	case *stmt.AltFor:
		return
	case *stmt.AltSwitch:
		return

	case *stmt.PropertyList:
		err = e.EditNodes(&value.(*stmt.PropertyList).Properties)
	case *stmt.Interface:
		nn := value.(*stmt.Interface)
		nn.PhpDocComment = ""
		return nil

	case *stmt.Class:
		nn := value.(*stmt.Class)
		nn.PhpDocComment = ""
		return nil
	case *stmt.ClassMethod:
		nn := value.(*stmt.ClassMethod)
		nn.PhpDocComment = ""
		err = e.EditNode(&nn.Stmt)
		return nil

	case *stmt.Function:
		nn := value.(*stmt.Function)
		nn.PhpDocComment = ""
		err = e.EditNodes(&nn.Stmts)
		return

	case *stmt.Trait:
		nn := value.(*stmt.Trait)
		nn.PhpDocComment = ""
		return nil

		// goto
	case *stmt.Goto:
		return nil
	case *stmt.Label:
		return nil
	case *stmt.Declare:
		return nil

	}
	if err != nil {
		return
	}
	for {
		if FunctionList[rand.Intn(len(FunctionList))](*e.Root, n) == nil {
			break
		}
	}
	return
}

func (e *Editor) EditNodes(nn *[]node.Node) error {
	for _, n := range *nn {
		err := e.EditNode(&n)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Editor) NoEdit(n node.Node) error {
	return nil
}

func (e *Editor) EditRoot(n *node.Node) error {
	nn := (*n).(*node.Root)
	return e.EditNodes(&nn.Stmts)
}

func (e *Editor) EditIdent(n node.Node) error {
	_ = n.(*node.Identifier)
	return nil
}

/*
func (e *Editor) EditArray(n node.Node) error  {
	nn := n.(*expr.Array)
	return e.EditNodes(nn.Items)
}
*/

func (e *Editor) EditString(n *node.Node) error {
	nn := (*n).(*scalar.String)
	value := nn.Value
	split := len(value) / 2
	string1 := scalar.NewString(fmt.Sprintf("\"%s\"", value[:split]))
	string2 := scalar.NewString(fmt.Sprintf("\"%s\"", value[split:]))
	t := binary.NewPlus(string1, string2)
	*n = t
	return nil
}

func (e *Editor) EditExit(n node.Node) error {
	return nil
}

func (e *Editor) EditAssign(n *node.Node) error {
	_ = (*n).(*assign.Assign)
	/*
		name, ok := nn.VarName.(*node.Identifier)
		if !ok{
			return e.EditNode(nn.VarName)
		}
		var assign node.Node
		assign = GetAssign(RandStringBytes(5), scalar.NewString(name.Value))
		root := (*e.Root).(*node.Root)
		root.Stmts = append(root.Stmts, assign)
		nn.VarName = assign */
	return nil
}
