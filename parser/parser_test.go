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
			want: nil,
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

func TestParser_ParseExpression(t *testing.T) {
	type fields struct {
		l         *lexer.Lexer
		curToken  token.Token
		peekToken token.Token
		errors    []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *BinaryExpr
	}{
		{
			name: "Simple addition",
			fields: fields{
				l:         lexer.New("2 + 3"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "2"},
				peekToken: token.Token{Type: token.PLUS, Literal: "+"},
			},
			want: &BinaryExpr{
				Left:     &NumberExpr{Value: 2},
				Operator: "+",
				Right:    &NumberExpr{Value: 3},
			},
		},
		{
			name: "Simple subtraction",
			fields: fields{
				l:         lexer.New("5 - 1"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "5"},
				peekToken: token.Token{Type: token.MINUS, Literal: "-"},
			},
			want: &BinaryExpr{
				Left:     &NumberExpr{Value: 5},
				Operator: "-",
				Right:    &NumberExpr{Value: 1},
			},
		},
		{
			name: "Addition and subtraction",
			fields: fields{
				l:         lexer.New("1 + 2 - 3"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "1"},
				peekToken: token.Token{Type: token.PLUS, Literal: "+"},
			},
			want: &BinaryExpr{
				Left: &BinaryExpr{
					Left:     &NumberExpr{Value: 1},
					Operator: "+",
					Right:    &NumberExpr{Value: 2},
				},
				Operator: "-",
				Right:    &NumberExpr{Value: 3},
			},
		},
		{
			name: "Complex expression",
			fields: fields{
				l:         lexer.New("3 + 5 - 2"),
				curToken:  token.Token{Type: token.NUMBER, Literal: "3"},
				peekToken: token.Token{Type: token.PLUS, Literal: "+"},
			},
			want: &BinaryExpr{
				Left: &BinaryExpr{
					Left:     &NumberExpr{Value: 3},
					Operator: "+",
					Right:    &NumberExpr{Value: 5},
				},
				Operator: "-",
				Right:    &NumberExpr{Value: 2},
			},
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
			if got := p.ParseExpression(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
