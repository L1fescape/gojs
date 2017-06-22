%{
package parser

import "fmt"
%}

%union {
  num int
}

%token NUMBER
%token ADD
%token SUB

%left ADD
%left SUB

%%
program: expr
{
  setParseResult(yylex, $1)
}

expr: NUMBER
{
  $$ = $1
}
| expr ADD expr
{
  fmt.Println($1.num, "+", $3.num)
  $$.num = $1.num + $3.num
}
| expr SUB expr
{
  fmt.Println($1.num, "-", $3.num)
  $$.num = $1.num - $3.num
}
