package confusedPHP

import (
	"os"
	"strings"
)

type PhpRead func(path string) (code *PhpCode, err error)

func NewPhpFile(filepath string) (code *PhpCode, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return NewPhpCode(file), nil
}

func NewPhpString(str string) (code *PhpCode, err error) {
	return NewPhpCode(strings.NewReader(str)), nil
}
