package eval

import (
	"fmt"
	"testing"

	"github.com/kh3rld/ksm-lang/ast"
	"github.com/kh3rld/ksm-lang/token"
)

func TestEvaluateAddition(t *testing.T) {
	tests := []struct {
		expression ast.Expression
		expected   float64
	}{
		{
			expression: &ast.InfixExpression{
				Token: token.Token{Type: token.PLUS, Literal: "+"},
				Left:  &ast.Num{Value: 3},
				Right: &ast.Num{Value: 5},
			},
			expected: 8,
		},
		{
			expression: &ast.InfixExpression{
				Token: token.Token{Type: token.PLUS, Literal: "+"},
				Left:  &ast.Num{Value: 3},
				Right: &ast.Num{Value: 2},
			},
			expected: 5.85,
		},
	}

	interp := &Evaluator{}

	for _, tt := range tests {
		result := interp.Evaluate(tt.expression)
		fmt.Printf("Result: %v\n", result) // Print the result for debugging
		if result != tt.expected {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
