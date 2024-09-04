package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL            TokenType = "illegal"
	EOF                TokenType = "eof"
	IDENT              TokenType = "ident"
	FLOAT              TokenType = "float"
	ASSIGN             TokenType = "="
	PLUS               TokenType = "+"
	MINUS              TokenType = "-"
	TIMES              TokenType = "*"
	DIVIDE             TokenType = "/"
	MODULUS            TokenType = "%"
	EQUAL              TokenType = "=="
	COMMA              TokenType = ","
	SEMICOLONTokenType TokenType = ";"
	COLON              TokenType = ":"
	SPACE              TokenType = " "
	CURLYOPENTokenType TokenType = "{"
	CURLYCLOSE         TokenType = "}"
	BRACKOPENTokenType TokenType = "("
	BRACKCLOSE         TokenType = ")"
	NUMBER             TokenType = "number"
	STRING             TokenType = "string"
	BOOLEAN            TokenType = "boolean"
	NULL               TokenType = "null"
	ARRAY              TokenType = "array"
	OBJECT             TokenType = "object"
	IDENTIFIER         TokenType = "identifier"
)
