package confusedPHP

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"math/rand"
	"reflect"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
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

var DefinitionType = []reflect.Type{
	reflect.TypeOf(&stmt.Class{}),
	reflect.TypeOf(&stmt.Function{}),
}

func IsDefinitionType(n *node.Node) bool {
	for _, val := range DefinitionType {
		if val == reflect.TypeOf(n) {
			return true
		}
	}
	return false
}
