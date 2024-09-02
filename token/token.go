package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    = "illegal"
	EOF        = "eof"
	IDENT      = "ident"
	FLOAT      = "float"
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	TIMES      = "*"
	DIVIDE     = "/"
	MODULUS    = "%"
	EQUAL      = "=="
	COMMA      = ","
	SEMICOLON  = ";"
	COLON      = ":"
	SPACE      = " "
	CURLYOPEN  = "{"
	CURLYCLOSE = "}"
	BRACKOPEN  = "("
	BRACKCLOSE = ")"
	NUMBER     = "number"
	STRING     = "string"
	BOOLEAN    = "boolean"
	NULL       = "null"
	ARRAY      = "array"
	OBJECT     = "object"
	IDENTIFIER = "identifier"
)
