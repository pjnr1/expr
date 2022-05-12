# expr
Golang rule engine, support logical expression

Support type: int string float bool

Support operation: > < >= <= && || + - * /

Support two inner funnction: in_array(1, []int{1,2,3,4}), ver_compare(x, ">", "10.1.1") with no nested

I believe it satisfy most filter scenarios using rule engine for logical expression. 

Finally, Welcome **contribution or requirement**.

# How to use it

```go
package main

import (
	expr "github.com/pjnr1/expr"
	log "github.com/s00500/env_logger"
)

func main() {
	expr := "1.0 * 2.0 > 1.5 + _test_"
	control := map[string]interface{}{"_test_": 1.5}
	engine, err := expr.NewEngine(expr)
	log.Must(err)
	
	result, err := engine.RunRule(control)
	log.Must(err)
	
	log.Infoln(result)
}
```

the ast parser using Golang go/ast and go/parser, rule can be update using engine.UpdateAst to avoid parse the rule each time.
more demo please refer to the expr_test.go

# Contributing
Pull requests, bug fixes and issue reports are welcome and appreciated.

# Licence
MIT

