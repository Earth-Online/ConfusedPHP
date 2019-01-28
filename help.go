package confusedPHP

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"math/rand"
	"reflect"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// get random string
func RandStringBytes(n uint) string {
	if n < 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandConfused() bool {
	rand.Seed(int64(time.Now().UnixNano()))
	if rand.Int()%2 == 1 {
		return true
	}
	return false
}

// php definition E.g function
var DefinitionType = []reflect.Type{
	reflect.TypeOf(&stmt.Class{}),
	reflect.TypeOf(&stmt.Function{}),
}

var ConstantType = []reflect.Type{
	reflect.TypeOf(&scalar.String{}),
	reflect.TypeOf(&scalar.Lnumber{}),
	reflect.TypeOf(&scalar.Dnumber{}),
	reflect.TypeOf(&scalar.Encapsed{}),
	reflect.TypeOf(&scalar.Heredoc{}),
}

// Determine if node  is  sDefinitionType
func IsDefinitionType(n *node.Node) bool {
	for _, val := range DefinitionType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}

func IsConstantType(n *node.Node) bool {
	for _, val := range ConstantType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}
