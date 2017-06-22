package main

import (
	"bytes"

	"code.justin.tv/web/audrey/_vendor/github.com/davecgh/go-spew/spew"
	"github.com/l1fescape/gojs/parser"
	"github.com/l1fescape/gojs/syntax"
)

func main() {
	node := parser.Parse(bytes.NewReader([]byte("1")))
	spew.Dump(node)

	n := syntax.BinaryOpNode{
		Operator: "-",
		Left: syntax.BinaryOpNode{
			Operator: "-",
			Left: syntax.NumberNode{
				Number: 6,
			},
			Right: syntax.NumberNode{
				Number: 2,
			},
		},
		Right: syntax.NumberNode{
			Number: 1,
		},
	}
	spew.Dump(n.Eval())
}
