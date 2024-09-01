package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let str = "Hello, World!";
              let isTrue = true;
              let isFalse = false;
              let nothing = null;`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LET, "let"},
		{IDENT, "str"},
		{ASSIGN, "="},
		{STRING, "Hello, World!"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "isTrue"},
		{ASSIGN, "="},
		{TRUE, "true"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "isFalse"},
		{ASSIGN, "="},
		{FALSE, "false"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "nothing"},
		{ASSIGN, "="},
		{NULL, "null"},
		{SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
