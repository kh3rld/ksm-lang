package parser

import (
	"fmt"
	"strconv"

	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}

	for p.curToken.Type != token.EOF {
		if p.curToken.Type == token.SPACE {
			p.nextToken()
			continue
		}
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() Node {
	switch p.curToken.Type {
	case token.NUMBER:
		return p.ParseExpression()
	case token.PLUS, token.MINUS:
		return p.ParseExpression()
	default:
		return nil
	}
}

func (p *Parser) ParseNumber() *NumberExpr {
	var value float64
	var err error

	if p.curToken.Type == token.MINUS {
		p.nextToken()
		if p.curToken.Type != token.NUMBER {
			p.errors = append(p.errors, "Expected a number after minus")
			return nil
		}
		value, err = strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Sprintf("Error parsing number: %s", err))
			return nil
		}
		value = -value
	} else if p.curToken.Type == token.NUMBER {
		value, err = strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Sprintf("Error parsing number: %s", err))
			return nil
		}
	} else {
		p.errors = append(p.errors, "Expected a number or a minus sign")
		return nil
	}

	return &NumberExpr{Value: value}
}

func (p *Parser) ParseExpression() *BinaryExpr {
	left := p.ParseNumber()
	if left == nil {
		return nil
	}

	operator := p.curToken
	p.nextToken()
	if p.curToken.Type != token.NUMBER {
		p.errors = append(p.errors, "Expected a number after operator")
		return nil
	}
	right := p.ParseNumber()
	if right == nil {
		return nil
	}

	return &BinaryExpr{
		Left:     left,
		Operator: operator.Literal,
		Right:    right,
	}
}
