// package parser

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/kh3rld/ksm-lang/lexer"
// 	"github.com/kh3rld/ksm-lang/token"
// )

// func TestParseNumber(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		input          string
// 		expectedValue  float64
// 		expectedErrors []string
// 	}{
// 		{
// 			name:          "Valid number",
// 			input:         "123",
// 			expectedValue: 123,
// 		},
// 		{
// 			name:          "Valid floating-point number",
// 			input:         "123.45",
// 			expectedValue: 123.45,
// 		},
// 		{
// 			name:           "Invalid number format",
// 			input:          "abc",
// 			expectedErrors: []string{"Error parsing number: invalid syntax"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := lexer.New(tt.input)
// 			p := New(l)

// 			// Parse the input
// 			p.nextToken() // move to the first token
// 			if p.curToken.Type != token.NUMBER {
// 				fmt.Printf("Error parsing number: %s", p.curToken.Type)
// 			}

// 			// Call parseNumber method (indirectly through public methods)
// 			result := p.ParseNumber()

// 			if len(p.Errors()) > 0 {
// 				for i, err := range p.Errors() {
// 					if i >= len(tt.expectedErrors) || err != tt.expectedErrors[i] {
// 						t.Errorf("unexpected error: got %q, want %q", err, tt.expectedErrors[i])
// 					}
// 				}
// 				return
// 			}

// 			if result != nil {
// 				if result.Value != tt.expectedValue {
// 					t.Errorf("expected %f, got %f", tt.expectedValue, result.Value)
// 				}
// 			} else if len(tt.expectedErrors) == 0 {
// 				t.Error("expected a result but got nil")
// 			}
// 		})
// 	}
// }

package parser

import (
	"testing"

	"github.com/kh3rld/ksm-lang/lexer"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "1 + 2  + 4 + 5 - 3",
			expected: 5,
		},
		{
			input:    "3 - 4 + 5",
			expected: 3,
		},
		{
			input:    "5",
			expected: 1,
		},
		{
			input:    " ",
			expected: 0,
		},
		{
			input:    "2 + 3 - 4 + 5",
			expected: 4,
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		if len(program.Statements) != tt.expected {
			t.Errorf("Expected %d statements, got %d", tt.expected, len(program.Statements))
		}
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue float64
		expectError   bool
	}{
		{
			input:         "-42",
			expectedValue: float64(-42),
			expectError:   false,
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		p.nextToken()

		if tt.expectError {
			p.nextToken()
			numberExpr := p.ParseNumber()
			if numberExpr != nil {
				t.Errorf("Expected error for input %q, but got valid number: %v", tt.input, numberExpr)
			}
			continue
		}

		numberExpr := p.ParseNumber()
		if numberExpr.Value != tt.expectedValue {
			t.Errorf("Expected %v for this input %q,  got %v", tt.expectedValue, tt.input, numberExpr.Value)
		}
	}
}
