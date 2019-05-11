package ast

type ExprNode interface {
	Accept()
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

func (n *IntNode) Accept()   {}
func (n *BinOpNode) Accept() {}
