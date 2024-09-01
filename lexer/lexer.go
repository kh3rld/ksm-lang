package lexer

import "unicode"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
		// Add more cases for other tokens like minus, asterisk, etc.
		case ';':  // Handle semicolon
        tok = newToken(SEMICOLON, l.ch)
    case ',':
        tok = newToken(COMMA, l.ch)  // Handle comma
    case '[':
        tok = newToken(LBRACKET, l.ch)  // Handle opening bracket
    case ']':
        tok = newToken(RBRACKET, l.ch) 

	case 0:
		tok.Literal = ""
		tok.Type = EOF
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	// Handling for true, false, and null literals
	case 't':
		if l.peekString("true") {
			tok = Token{Type: TRUE, Literal: "true"}
			l.readCharN(4)
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case 'f':
		if l.peekString("false") {
			tok = Token{Type: FALSE, Literal: "false"}
			l.readCharN(5)
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case 'n':
		if l.peekString("null") {
			tok = Token{Type: NULL, Literal: "null"}
			l.readCharN(4)
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekString(expected string) bool {
	for i := 0; i < len(expected); i++ {
		if l.input[l.position+i] != expected[i] {
			return false
		}
	}
	return true
}

func (l *Lexer) readCharN(n int) {
	for i := 0; i < n; i++ {
		l.readChar()
	}
}
