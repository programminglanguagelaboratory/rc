package parser

import (
	"errors"
	"fmt"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

type Parser struct {
	lexer *lexer.Lexer
	tok   token.Token
}

func (p *Parser) next() {
	p.tok, _ = p.lexer.Lex()
}

func NewParser(l *lexer.Lexer) *Parser {
	p := Parser{}
	p.lexer = l
	p.tok, _ = p.lexer.Lex()
	return &p
}

func (p *Parser) Parse() (ast.Expr, error) {
	return p.parseExpr()
}

func (p *Parser) parseExpr() (ast.Expr, error) {
	return p.parseBinaryExpr(0)
}

func (p *Parser) parseBinaryExpr(prevPrec int) (ast.Expr, error) {
	left, err := p.parseUnaryExpr()
	if err != nil {
		return nil, err
	}

	for {
		prec := p.tok.Kind.GetPrec()
		if prec < prevPrec {
			return left, nil
		}

		tok := p.tok
		p.next()

		right, err := p.parseBinaryExpr(prec + 1)
		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpr{Left: left, Right: right, Token: tok}
	}
}

func (p *Parser) parseUnaryExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.MINUS, token.EXCLAMATION:
		t := ast.UnaryExpr{Token: p.tok}
		p.next()

		left, err := p.parseUnaryExpr()
		if err != nil {
			return nil, err
		}

		t.Left = left
		return t, nil
	}

	left, err := p.parseLitOrParenExpr()
	if err != nil {
		return nil, err
	}

	return left, nil
}

func (p *Parser) parseLitOrParenExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.ID, token.STRING, token.NUMBER, token.BOOL:
		t := ast.LitExpr{Token: p.tok}
		p.next()
		return t, nil
	case token.LPAREN:
		return p.parseParenExpr()
	}

	return nil, errors.New("unexpected token")
}

func (p *Parser) parseParenExpr() (ast.Expr, error) {
	if p.tok.Kind != token.LPAREN {
		return nil, fmt.Errorf("expected lparen, but got: %v", p.tok.Kind)
	}
	p.next()

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if p.tok.Kind != token.RPAREN {
		return nil, fmt.Errorf("expected rparen, but got: %v", p.tok.Kind)
	}
	p.next()

	return expr, nil
}
