package parser

import (
	"go/token"

	"github.com/kh3rld/ksm-lang/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	// p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() Program {
	var program Program
	return program
}

func (p *Parser) parseNumber() {
	// literal := p.curToken.Literal
	// value, err := strconv.ParseFloat(literal, 64)
	// if err != nil {
	// 	// Handle error appropriately
	// 	return nil
	// }
	// return &NumberExpr{Value: value}
}
