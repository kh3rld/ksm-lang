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

func newToken(tokenType token.TokenType, ch string) token.Token {
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
		l.ch = "0"
	} else {
		l.ch = string(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case "=":
		t = newToken(token.ASSIGN, l.ch)
	case "+":
		t = newToken(token.PLUS, l.ch)
	case "-":
		t = newToken(token.MINUS, l.ch)
	case "*":
		t = newToken(token.TIMES, l.ch)
	case "/":
		t = newToken(token.DIVIDE, l.ch)
	case "%":
		t = newToken(token.MODULUS, l.ch)
	case ",":
		t = newToken(token.COMMA, l.ch)
	case ";":
		t = newToken(token.SEMICOLON, l.ch)
	case ":":
		t = newToken(token.COLON, l.ch)
	case "==":
		t = newToken(token.EQUAL, l.ch)
	case " ":
		t = newToken(token.SPACE, l.ch)
	case "{":
		t = newToken(token.CURLYOPEN, l.ch)
	case "}":
		t = newToken(token.CURLYCLOSE, l.ch)
	case "(":
		t = newToken(token.BRACKOPEN, l.ch)
	case ")":
		t = newToken(token.BRACKCLOSE, l.ch)
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
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}
