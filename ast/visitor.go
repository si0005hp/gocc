package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// GoccParser tokens.
const (
	GoccParserEOF     = antlr.TokenEOF
	GoccParserINT     = 1
	GoccParserADD     = 2
	GoccParserSUB     = 3
	GoccParserMUL     = 4
	GoccParserDIV     = 5
	GoccParserNEWLINE = 6
	GoccParserWS      = 7
)

type Visitor interface {
	VisitInt(n IntNode)
	VisitBinOp(n BinOpNode)
}

type CodeGenerator struct {
}

func (g *CodeGenerator) VisitInt(n IntNode) {
	fmt.Printf("mov $%d, %%eax\n", n.Val)
}

func (g *CodeGenerator) VisitBinOp(n BinOpNode) {
	var op string
	switch n.Op {
	case GoccParserADD:
		op = "add"
	case GoccParserSUB:
		op = "sub"
	case GoccParserMUL:
		op = "mul"
	case GoccParserDIV:
	}
	n.Right.Accept(g)
	fmt.Printf("push %%rax\n")
	n.Left.Accept(g)

	if n.Op == GoccParserDIV {
		fmt.Println("pop %%rcx")
		fmt.Println("mov $0, %%edx")
		fmt.Println("idiv %%rcx")
	} else {
		fmt.Println("pop %%rcx")
		fmt.Printf("%s %%rcx, %%rax\n", op)
	}
}
