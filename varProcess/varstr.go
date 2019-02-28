package varProcess

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
)

type VarStr struct {
	*expr.Variable
}

func NewVarStr(variable *expr.Variable) *VarStr {
	return &VarStr{Variable: variable}
}

func (v VarStr) Name() string {
	switch v.Variable.VarName.(type) {
	case *node.Identifier:
		return v.Variable.VarName.(*node.Identifier).Value
	case *scalar.String:
		return v.Variable.VarName.(*scalar.String).Value
	default:
		return ""
	}
}

func (v VarStr) String() string {
	return fmt.Sprintf("$%s", v.Name())
}

func (v VarStr) Len() int {
	return len(v.Name())
}
