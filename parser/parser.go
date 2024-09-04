package parser

import (
	"strconv"

	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	Value     int
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.curToken
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case token.NUMBER:
		return p.parseNumber()
	default:
		return nil
	}
}

func (p *Parser) parseNumber() *Parser {
	value, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		return nil
	}
	return &Parser{Value: value}
}

