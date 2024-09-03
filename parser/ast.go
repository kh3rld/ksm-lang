package parser

type Visitor interface {
	VisitNumberExpr(expr *NumberExpr) interface{}
}

type Expr interface {
	Accept(visitor Visitor) interface{}
}

type Program struct {
	Statements []Statement
}

type Statement interface{}

type Expression interface {
	Accept(visitor Visitor) interface{}
}

type NumberExpr struct {
	Value float64
}

func (n *NumberExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitNumberExpr(n)
}
