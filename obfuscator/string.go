package obfuscator

import (
	"encoding/base64"
	"fmt"
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/scalar"
)

var Base64Obfuscator = nodeProcess.NewStringPrecess("base64", func(n string) (append []node.Node, cur node.Node) {
	value := fmt.Sprintf("\"%s\"", base64.StdEncoding.EncodeToString([]byte(n)))
	var nameNode node.Node
	nameNode = node.NewIdentifier("base64")
	args := util.GetFunctionArg(scalar.NewString(value))
	f := util.GetFunctionCall(nameNode, args.(*node.ArgumentList))
	return []node.Node{}, f
})

var StringSplitObfuscator = nodeProcess.NewStringPrecess("string split", func(str string) (append []node.Node, cur node.Node) {

	split := len(str) / 2
	string1 := scalar.NewString(fmt.Sprintf("\"%s\"", str[:split]))
	string2 := scalar.NewString(fmt.Sprintf("\"%s\"", str[split:]))
	t := binary.NewPlus(string1, string2)
	return []node.Node{}, t
})

var GzCompressObfuscator = nodeProcess.NewStringPrecess("gzcompress", func(str string) (app []node.Node, cur node.Node) {
	compress, err := util.ZlibCompress([]byte(str))
	if err != nil {
		return
	}

	nn := &scalar.String{
		Value: fmt.Sprintf("\"%s\"", compress),
	}
	var nameNode node.Node = node.NewIdentifier("gzcompress")
	var args = util.GetFunctionArg(nn)
	nnn := util.GetFunctionCall(nameNode, args.(*node.ArgumentList))
	return nil, nnn
})
