package parser

import (
	"testing"

	"github.com/kh3rld/ksm-lang/ast"
	"github.com/kh3rld/ksm-lang/eval"
	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/token"
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

// Test function for parseNum.
func TestParseNum(t *testing.T) {
	tests := []struct {
		name          string
		inputToken    token.Token
		expectedValue int64
		expectError   bool
		expectedError string
	}{
		{
			name:          "Valid number",
			inputToken:    token.Token{Type: token.NUM, Literal: "123"},
			expectedValue: 123,
			expectError:   false,
		},
		{
			name:          "Valid negative number",
			inputToken:    token.Token{Type: token.NUM, Literal: "-456"},
			expectedValue: -456,
			expectError:   false,
		},
		{
			name:          "Invalid number format",
			inputToken:    token.Token{Type: token.NUM, Literal: "abc"},
			expectError:   true,
			expectedError: "Error parsing number: strconv.ParseInt: parsing \"abc\": invalid syntax",
		},
		{
			name:          "Out of range number",
			inputToken:    token.Token{Type: token.NUM, Literal: "9999999999999999999999999999"},
			expectError:   true,
			expectedError: "Error parsing number: strconv.ParseInt: parsing \"9999999999999999999999999999\": value out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{curToken: tt.inputToken, errors: []string{}}
			result := parser.parseNum()

			if tt.expectError {
				if result != nil {
					t.Errorf("Expected nil, got %v", result)
				}
				if len(parser.errors) == 0 || parser.errors[0] != tt.expectedError {
					t.Errorf("Expected error: %v, got: %v", tt.expectedError, parser.errors)
				}
			} else {
				if result == nil {
					t.Errorf("Expected a valid ast.Num, got nil")
					return
				}

				num, ok := result.(*ast.Num)
				if !ok {
					t.Errorf("Expected result type *ast.Num, got %T", result)
					return
				}

				if num.Value != tt.expectedValue {
					t.Errorf("Expected value %d, got %d", tt.expectedValue, num.Value)
				}
			}
		})
	}
}
