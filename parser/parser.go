package parser

import "io"

func Parse(reader io.Reader) interface{} {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}
