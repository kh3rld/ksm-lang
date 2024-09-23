package parser

import (
	"reflect"
	"testing"

	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/token"
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

func TestParser_ParseNumber(t *testing.T) {
	type fields struct {
		l         *lexer.Lexer
		curToken  token.Token
		peekToken token.Token
		errors    []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *NumberExpr
	}{
		{
			name: "Parse positive number",
			fields: fields{
				l:         lexer.New("45"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "45"},
				peekToken: token.Token{Type: token.EOF, Literal: ""},
				errors:    []string{},
			},
			want: &NumberExpr{Value: 45},
		},
		{
			name: "Parse negative number",
			fields: fields{
				l:         lexer.New("-45"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "-45"},
				peekToken: token.Token{Type: token.EOF, Literal: ""},
				errors:    []string{},
			},
			want: &NumberExpr{Value: -45},
		},
		{
			name: "Parse zero",
			fields: fields{
				l:         lexer.New("0"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "0"},
				peekToken: token.Token{Type: token.EOF, Literal: ""},
				errors:    []string{},
			},
			want: &NumberExpr{Value: 0},
		},
		{
			name: "Invalid number input",
			fields: fields{
				l:         lexer.New("hello"),
				curToken:  token.Token{Type: token.IDENT, Literal: "hello"},
				peekToken: token.Token{Type: token.EOF, Literal: ""},
				errors:    []string{"expected number, got identifier"},
			},
			want: nil,
		},
		{
			name: "Invalid mixed input '5d2'",
			fields: fields{
				l:         lexer.New("5d2"),
				curToken:  token.Token{Type: token.IDENT, Literal: "5d2"},
				peekToken: token.Token{Type: token.EOF, Literal: ""},
				errors:    []string{"invalid number format: '5d2'"},
			},
			want: nil, // Adjust this to match how your parser handles errors or invalid input
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				l:         tt.fields.l,
				curToken:  tt.fields.curToken,
				peekToken: tt.fields.peekToken,
				errors:    tt.fields.errors,
			}
			if got := p.ParseNumber(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
