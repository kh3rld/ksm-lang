package eval

import (
	"github.com/kh3rld/ksm-lang/parser"
)

type Evaluator struct{}

func (e *Evaluator) Eval(node parser.Node) *Number {
	switch n := node.(type) {
	case *parser.NumberExpr:
		return &Number{Value: n.Value}
	case *parser.BinaryExpr:
		left := e.Eval(n.Left)
		right := e.Eval(n.Right)
		if left == nil || right == nil {
			return nil
		}
		return EvaluateArithmetic(left, right, n.Operator)

func (e *Evaluator) VisitNumberExpr(expr *parser.NumberExpr) interface{} {
	return &Number{Value: expr.Value}
}

func EvaluateArithmetic(left, right *Number, op string) *Number {
	switch op {
	case "+":
		return left.Add(right)
	case "-":
		return left.Subtract(right)
	default:
		return nil
	}
}
