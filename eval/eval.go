package eval

import (
	"github.com/kh3rld/ksm-lang/parser"
)

type Evaluator struct{}

func (e *Evaluator) VisitNumberExpr(expr *parser.NumberExpr) interface{} {
	return &Number{Value: expr.Value}
}

func Eval(node parser.Expression) *Number {
	switch n := node.(type) {
	case *parser.NumberExpr:
		return &Number{Value: n.Value}
	}
	return nil
}

func EvaluateArithmetic(left, right *Number, op string) *Number {
	switch op {
	case "+":
		return left.Add(right)
	case "-":
		return left.Subtract(right)
	}
	return nil
}
