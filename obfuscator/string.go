package obfuscator

import (
	"encoding/base64"
	"fmt"
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
)

var Base64Obfuscator = nodeProcess.NewStringPrecess("base54", func(n string) (append []node.Node, cur node.Node) {
	value := fmt.Sprintf("\"%s\"", base64.StdEncoding.EncodeToString([]byte(n)))
	var nameNode node.Node
	nameNode = node.NewIdentifier("base64")
	args := util.GetFunctionArg(scalar.NewString(value))
	f := util.GetFunctionCall(nameNode, args.(*node.ArgumentList))
	return []node.Node{}, f
})
