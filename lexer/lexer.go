package lexer

import (
	"github.com/kh3rld/ksm-lang/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func isLetter(r byte) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func NewToken(tokenType token.TokenType, r byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(r)}
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
	return l.input[position:l.position]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	// Skip whitespace
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

	switch l.ch {
	case '+':
		t = NewToken(token.PLUS, l.ch)
	case '-':
		t = NewToken(token.MINUS, l.ch)
	case '*':
		t = NewToken(token.TIMES, l.ch)
	case '/':
		t = NewToken(token.DIVIDE, l.ch)
	case '%':
		t = NewToken(token.MODULUS, l.ch)
	case ',':
		t = NewToken(token.COMMA, l.ch)
	case ';':
		t = NewToken(token.SEMICOLON, l.ch)
	case ':':
		t = NewToken(token.COLON, l.ch)
	case '{':
		t = NewToken(token.CURLYOPEN, l.ch)
	case '}':
		t = NewToken(token.CURLYCLOSE, l.ch)
	case '(':
		t = NewToken(token.BRACKOPEN, l.ch)
	case ')':
		t = NewToken(token.BRACKCLOSE, l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			t = NewToken(token.ASSIGN, l.ch)
		}
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) || l.ch == '.' {
			t.Type = token.NUM
			t.Literal = l.readNumber()
			return t
		} else {
			t = NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}
