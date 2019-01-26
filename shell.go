package confusedPHP

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"os"
)

type Shell struct {
	Filename string
	Code     *php7.Parser
}

func NewShell(filename string) *Shell {
	return &Shell{Filename: filename}
}

func (s *Shell) Parser() (err error) {
	file, err := os.Open(s.Filename)
	if err != nil {
		return
	}
	parser := php7.NewParser(file, "example.php")
	parser.Parse()
	s.Code = parser
	return
}

func (s *Shell) GetRoot() node.Node {
	return s.Code.GetRootNode()
}
