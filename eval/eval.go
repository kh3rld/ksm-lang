package eval

import (
	"log"

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
		return EvaluateArithmetic(left, n.Operator, right)
	}
	return nil
}

func EvaluateArithmetic(left *Number, op string, right *Number) *Number {
	switch op {
	case "+":
		return left.Add(right)
	case "-":
		return left.Subtract(right)
	default:
		log.Printf("Unexpected node type: %T\n", op)
		return nil
	}
}
