package eval

import (
	"github.com/kh3rld/ksm-lang/parser"
)

type Evaluator struct{}

func (e *Evaluator) VisitNumberExpr(expr *parser.NumberExpr) interface{} {
	return &Number{Value: expr.Value}
}
