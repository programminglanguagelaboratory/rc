package parser

import (
	"errors"
	"fmt"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

type parser struct {
	lexer *lexer.Lexer
	tok   token.Token
}

// TODO: handle .Lex error
func (p *parser) next() {
	p.tok, _ = p.lexer.Lex()
}

// TODO: handle Lex error
func newParser(l *lexer.Lexer) *parser {
	p := parser{}
	p.lexer = l
	p.tok, _ = p.lexer.Lex()
	return &p
}

func (p *parser) parseExpr() (ast.Expr, error) {
	return p.parseBinaryExpr(1)
}

func (p *parser) parseBinaryExpr(prevPrec int) (ast.Expr, error) {
	left, err := p.parsePrimaryExpr()
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

func (p *parser) parsePrimaryExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.ID, token.STRING, token.NUMBER:
		t := ast.LitExpr{Token: p.tok}
		p.next()
		return t, nil
	case token.LPAREN:
		return p.parseParenExpr()
	}
	return nil, errors.New("unexpected token")
}

func (p *parser) parseParenExpr() (ast.Expr, error) {
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
