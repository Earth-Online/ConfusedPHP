package phpread

import "testing"

func TestNewPhpFile(t *testing.T) {
	shell, err := NewPhpFile("./test.php")
	err = shell.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	err = shell.Parser()
	if err != nil {
		t.Error(err)
	}
	shell, err = NewPhpFile("./fake.php")
	if err == nil {
		t.Error("no open file error")
		return
	}
}

func TestNewPhpString(t *testing.T) {
	src := `<?php $a;`
	shell, err := NewPhpString(src)
	err = shell.Parser()
	if err != nil {
		t.Error(err)
	}
	err = shell.Parser()
	if err != nil {
		t.Error(err)
	}
}
