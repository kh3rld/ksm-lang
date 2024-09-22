package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    TokenType = "illegal"
	EOF        TokenType = "eof"
	IDENT      TokenType = "ident"
	FLOAT      TokenType = "float"
	ASSIGN     TokenType = "="
	PLUS       TokenType = "+"
	MINUS      TokenType = "-"
	TIMES      TokenType = "*"
	DIVIDE     TokenType = "/"
	MODULUS    TokenType = "%"
	EQUAL      TokenType = "=="
	NOTEQUAL   TokenType = "!="
	NOT        TokenType = "!"
	COMMA      TokenType = ","
	SEMICOLON  TokenType = ";"
	COLON      TokenType = ":"
	SPACE      TokenType = " "
	CURLYOPEN  TokenType = "{"
	CURLYCLOSE TokenType = "}"
	BRACKOPEN  TokenType = "("
	BRACKCLOSE TokenType = ")"
	NUM        TokenType = "num"
	STRING     TokenType = "str"
	BOOLEAN    TokenType = "bool"
	NULL       TokenType = "nil"
	ARRAY      TokenType = "arr"
	OBJECT     TokenType = "object"
	IDENTIFIER TokenType = "identifier"
	GREATER    TokenType = ">"
	LESS       TokenType = "<"
	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "rtn"
)

var keywords = map[string]TokenType{
	"fn":    FUNCTION,
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
	"rtn":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
