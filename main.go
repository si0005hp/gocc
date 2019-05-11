package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/si0005hp/gocc/parser"
)

func main() {
	i, _ := antlr.NewFileStream(os.Args[1])
	p := getParser(i)

	fmt.Println(p)
}

func getParser(input *antlr.FileStream) *parser.GoccParser {
	lexer := parser.NewGoccLexer(input)
	parser := parser.NewGoccParser(antlr.NewCommonTokenStream(lexer, 0))
	return parser
}

// import (
// 	"fmt"

// 	"github.com/si0005hp/gocc/ast"
// )

// func main() {
// 	i := &ast.IntNode{}
// 	fmt.Println(i)
// }
