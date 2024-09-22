package ast

import (
	"testing"

	"github.com/kh3rld/ksm-lang/token"
)

func TestNum(t *testing.T) {
	num := &Num{
		Token: token.Token{Type: token.NUM, Literal: "5"},
		Value: 5,
	}
	if num.TokenLiteral() != "5" {
		t.Fatalf("TokenLiteral() wrong. got=%s", num.TokenLiteral())
	}
	if num.String() != "5" {
		t.Fatalf("String() wrong. got=%s", num.String())
	}
}
