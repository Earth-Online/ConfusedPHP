package util

import "testing"

func TestRandStringBytes(t *testing.T) {
	if len(RandStringBytes(5)) != 5 {
		t.Error("error random string")
	}
}
