package parser

import (
	"io"

	"github.com/golang/go/src/pkg/strconv"
	"github.com/l1fescape/gojs/syntax"
)

func Parse(reader io.Reader) syntax.Node {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult.(yySymType).node
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}

func createNumberNode(o yySymType) yySymType {
	number, err := strconv.Atoi(o.text)
	if err != nil {
		panic(err)
	}
	return yySymType{
		node: syntax.NumberNode{Number: number},
	}
}
