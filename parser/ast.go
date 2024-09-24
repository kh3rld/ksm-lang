package parser

type Visitor interface {
	VisitNumberExpr(expr *NumberExpr) interface{}
	VisitBinaryExpr(expr *BinaryExpr) interface{}
}

type Node interface {
	Accept(visitor Visitor) interface{}
}

type Program struct {
	Statements []Node
}

type NumberExpr struct {
	Value float64
}

func (n *NumberExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitNumberExpr(n)
}

type BinaryExpr struct {
	Left     Node
	Operator string
	Right    Node
}

func (b *BinaryExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryExpr(b)
}
