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

// func TestParseNumber(t *testing.T) {
// 	tests := []struct {
// 		input         string
// 		expectedValue float64
// 		expectError   bool
// 	}{
// 		{
// 			input:         "-42",
// 			expectedValue: float64(-42),
// 			expectError:   false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		l := lexer.New(tt.input)
// 		p := New(l)
// 		p.nextToken()

// 		if tt.expectError {
// 			p.nextToken()
// 			numberExpr := p.ParseNumber()
// 			if numberExpr != nil {
// 				t.Errorf("Expected error for input %q, but got valid number: %v", tt.input, numberExpr)
// 			}
// 			continue
// 		}

// 		numberExpr := p.ParseNumber()
// 		if numberExpr.Value != tt.expectedValue {
// 			t.Errorf("Expected %v for this input %q,  got %v", tt.expectedValue, tt.input, numberExpr.Value)
// 		}
// 	}
// }

// TestParseNumber function
func TestParseNumber(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue float64
		expectError   bool
	}{
		{"-42", -42, false},
		{"42", 42, false},
		{"-abc", 0, true}, // Invalid number
		{"abc", 0, true},  // Not a number
	}

	for _, tt := range tests {
		lexer := lexer.New(tt.input)
		parser := New(lexer)
		parser.nextToken() // Setup initial token

		numberExpr := parser.ParseNumber()

		if tt.expectError {
			if numberExpr != nil {
				t.Errorf("Expected error for input %q, but got valid number: %v", tt.input, numberExpr)
			}
			if len(parser.errors) == 0 {
				t.Errorf("Expected parsing errors for input %q, but none were found", tt.input)
			}
			continue
		}

		if numberExpr == nil {
			t.Errorf("Expected valid number for input %q, but got nil", tt.input)
			continue
		}

		if numberExpr.Value != tt.expectedValue {
			t.Errorf("Expected %v for input %q, got %v", tt.expectedValue, tt.input, numberExpr.Value)
		}
	}
}
