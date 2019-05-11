package main

import (
	"flag"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/si0005hp/gocc/ast"
	"github.com/si0005hp/gocc/parser"
)

func main() {
	flag.Parse()
	i, _ := antlr.NewFileStream(flag.Arg(0))

	if flag.Arg(1) == "v" {
		treeView(i)
	} else {
		run(i)
	}
}

func run(input *antlr.FileStream) {
	p := getParser(input)
	g := &ast.CodeGenerator{}
	p.Expr().GetN().Accept(g)
}

func treeView(input *antlr.FileStream) {
	p := getParser(input)
	s := antlr.TreesStringTree(p.Expr(), p.GetRuleNames(), nil)
	fmt.Println(s)
}

func getParser(input *antlr.FileStream) *parser.GoccParser {
	lexer := parser.NewGoccLexer(input)
	parser := parser.NewGoccParser(antlr.NewCommonTokenStream(lexer, 0))
	return parser
}
