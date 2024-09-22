package eval

import "github.com/kh3rld/ksm-lang/ast"

type Evaluator struct{}

func (e *Evaluator) Evaluate(node ast.Node) interface{} {
	switch node := node.(type) {
	case *ast.Num:
		return node.Value
	case *ast.InfixExpression:
		left := e.Evaluate(node.Left)
		right := e.Evaluate(node.Right)

		switch node.Operator {
		case "+":
			return left.(int64) + right.(int64)
		}
	}
	return nil
}
