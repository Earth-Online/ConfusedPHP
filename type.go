package confusedPHP

import (
	"fmt"
	"github.com/z7zmey/php-parser/node/expr"
)

type VarStr struct {
	expr.Variable
}

func (v VarStr) String() string {
	return fmt.Sprintf("$%s", v.VarName)
}
