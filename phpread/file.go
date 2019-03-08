package phpread

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php7"
	"io"
)

// PhpCode is a php file
type PhpCode struct {
	// php file name
	Content io.Reader
	// parser
	Code *php7.Parser
}

// NewPhpCode Construct PhpCode
func NewPhpCode(content io.Reader) *PhpCode {
	return &PhpCode{Content: content}
}

// Parser parser code
func (s *PhpCode) Parser() (err error) {
	parser := php7.NewParser(s.Content, "example.php")
	parser.Parse()
	s.Code = parser
	return
}

// GetRootNode get code root node
func (s *PhpCode) GetRootNode() node.Node {
	return s.Code.GetRootNode()
}
