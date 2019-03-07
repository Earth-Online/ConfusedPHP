package tool

import (
	"bytes"
	"github.com/blue-bird1/ConfusedPHP/phpread"
	"github.com/z7zmey/php-parser/printer"
)

func DeleteBlankLine(code string) (newCode string, err error) {
	parser, err := phpread.NewPhpString(code)
	if err != nil {
		return
	}
	err = parser.Parser()
	if err != nil {
		return
	}
	root := parser.GetRootNode()

	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)

	p.Print(root)

	return o.String(), nil
}
