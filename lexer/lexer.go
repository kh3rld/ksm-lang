package lexer

import (
	"github.com/kh3rld/ksm-lang/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           string
}

func isLetter(ch string) bool {
	return ch >= "a" && ch <= "z" || ch >= "A" && ch <= "Z"
}

func isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func NewToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	// Check if there is a decimal point
	if l.ch == "." {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = "\x00" // EOF
	} else {
		l.ch = string(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	// Skip whitespace
	for l.ch == " " || l.ch == "\t" || l.ch == "\n" {
		l.readChar()
	}

	switch l.ch {
	case "+":
		t = NewToken(token.PLUS, l.ch)
	case "-":
		t = NewToken(token.MINUS, l.ch)
	case "*":
		t = NewToken(token.TIMES, l.ch)
	case "/":
		t = NewToken(token.DIVIDE, l.ch)
	case "%":
		t = NewToken(token.MODULUS, l.ch)
	case ",":
		t = NewToken(token.COMMA, l.ch)
	case ";":
		t = NewToken(token.SEMICOLON, l.ch)
	case ":":
		t = NewToken(token.COLON, l.ch)
	case "==":
		t = NewToken(token.EQUAL, l.ch)
	case " ":
		t = NewToken(token.SPACE, l.ch)
	case "{":
		t = NewToken(token.CURLYOPEN, l.ch)
	case "}":
		t = NewToken(token.CURLYCLOSE, l.ch)
	case "(":
		t = NewToken(token.BRACKOPEN, l.ch)
	case ")":
		t = NewToken(token.BRACKCLOSE, l.ch)
	case "\x00":
		t.Type = token.EOF
		t.Literal = ""
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.IDENT
			return t
		} else if isDigit(l.ch) || l.ch == "." {
			t.Literal = l.readNumber()
			t.Type = token.NUMBER
			return t
		} else {
			t = NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}
