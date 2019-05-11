package ast

type ExprNode interface {
	Accept(v Visitor)
}

type (
	IntNode struct {
		Val int
	}

	BinOpNode struct {
		Op    int
		Left  ExprNode
		Right ExprNode
	}
)

func (n *IntNode) Accept(v Visitor)   { v.VisitInt(*n) }
func (n *BinOpNode) Accept(v Visitor) { v.VisitBinOp(*n) }
