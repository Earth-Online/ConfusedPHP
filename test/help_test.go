package test

import (
	"github.com/blue-bird1/ConfusedPHP"
	"testing"
)

func TestRandString(t *testing.T) {
	ret := confusedPHP.RandStringBytes(5)
	if len(ret) != 5 {
		t.Error("Error generating  string")
	}
}
