package ast

type Expr interface {
	Accept()
}

type (
	IntNode struct {
		Val int
	}

	ExprNode struct {
		op    string
		left  *ExprNode
		right *ExprNode
	}
)

func (n *IntNode) Accept()  {}
func (n *ExprNode) Accept() {}
