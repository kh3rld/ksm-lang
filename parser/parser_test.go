package parser

import (
	"testing"

	"github.com/kh3rld/ksm-lang/ast"
	"github.com/kh3rld/ksm-lang/eval"
	"github.com/kh3rld/ksm-lang/lexer"
)

func TestAddition(t *testing.T) {
	input := `5 + 10;`
	rst := int64(15)
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		t.Fatalf("parser has %d errors", len(p.Errors()))
	}

	interp := &eval.Evaluator{}
	result := interp.Evaluate(program.Statements[0].(*ast.ExpressionStatement).Expression)

	if result != rst {
		t.Fatalf("Expected %v got %v", rst, result)
	}
}
