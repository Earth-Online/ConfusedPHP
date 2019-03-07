package tool

import "testing"

func TestDeleteBlankLine(t *testing.T) {
	testCode := `
	<?php
	 		
echo "hello world";
$a=1; 
	`
	code, err := DeleteBlankLine(testCode)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(code)
}
