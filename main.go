package main

import (
	"github.com/l1fescape/gojs/parser"
	"bytes"
	"code.justin.tv/web/audrey/_vendor/github.com/davecgh/go-spew/spew"
)

func main() {
	node := parser.Parse(bytes.NewReader([]byte("3 - 1 - 1")))
	spew.Dump(node)
}
