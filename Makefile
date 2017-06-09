.PHONY: all precommit setup test build run fix gen

setup:
	go get -u github.com/blynn/nex
	go get -u golang.org/x/tools/cmd/goyacc

clean:
	-rm y.output

gen: clean
	goyacc -o=parser/parser.gen.go parser/parser.y
	nex -o parser/lexer.gen.go parser/lexer.nex

fix:
	gofmt -w -l .
	goimports -w -l .

test:
	go test ./...

precommit: fix test

build: clean gen test
	go build

all: clean setup build
