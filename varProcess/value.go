package varProcess

import (
	"errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
)

func GetStingTypeValue(n node.Node) (str string, err error) {
	switch n.(type) {
	case *scalar.String:
		return n.(*scalar.String).Value, nil
	case *scalar.Heredoc:
		return ProcessVar(n.(*scalar.Heredoc).Parts), nil
	case *scalar.Encapsed:
		return ProcessVar(n.(*scalar.Encapsed).Parts), nil
	default:
		err = errors.New("not a string type")
		return
	}
}

func ProcessVar(n []node.Node) (s string) {
	for _, value := range n {
		switch value.(type) {
		case *scalar.EncapsedStringPart:
			s = s + value.(*scalar.EncapsedStringPart).Value
		case *expr.Variable:
			s = s + NewVarStr(value.(*expr.Variable)).String()
		}
	}
	return
}
