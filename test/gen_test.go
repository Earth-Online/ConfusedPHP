package test

import (
	"github.com/blue-bird1/ConfusedPHP"
	"github.com/z7zmey/php-parser/node/scalar"
	"testing"
)

// var p = printer.NewPrinter(os.Stdout, "    ")

func TestFunctionRet(t *testing.T) {
	n := confusedPHP.GetFunctionRet("test", scalar.NewString("'test'"))
	p.Print(n)
}
