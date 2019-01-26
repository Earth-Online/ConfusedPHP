package confusedPHP

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"os"
)

// php file
type Shell struct {
	// php file name
	Filename string
	// parser
	Code *php7.Parser
}

// create new shell
func NewShell(filename string) *Shell {
	return &Shell{Filename: filename}
}

// parser code
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

// get code root node
func (s *Shell) GetRoot() node.Node {
	return s.Code.GetRootNode()
}
