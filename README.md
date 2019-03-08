## ConfusedPHP 
一个开发中的php混淆器. 基于AST等价修改. 目前版本0.05

### 下载
`go get  github.com/blue-bird1/ConfusedPHP`

### 使用
#### hello world
```go
package main

import (
	`os`
	`fmt`
  `github.com/blue-bird1/ConfusedPHP/editor`
  `github.com/blue-bird1/ConfusedPHP/nodeProcess`
  `github.com/blue-bird1/ConfusedPHP/obfuscator`
   `github.com/blue-bird1/ConfusedPHP/phpread`
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/node"
)

func main() {
	edit := editor.NewEditWalker([]nodeProcess.NodePrecess{obfuscator.Base64Obfuscator})
    testCode := `
    	<?php
    		eval("ls");
    	`
    parser, err := phpread.NewPhpString(testCode)
    if err != nil {
    		panic(err)
    		return
    }
    err = parser.Parser()
    if err != nil {
    	panic(err)
    	return
    }
    root := parser.GetRootNode()
    root.Walk(edit)
    if len(edit.AddNode()) != 0 {
    		p2 := printer.NewPrinter(os.Stdout)
    		p2.Print(node.NewRoot(edit.AddNode()))
    		fmt.Print("?>")
    }
    	p := editor.NewPrinter(os.Stdout, edit.ModifyNode())
    	p.Print(root)
  }



```

