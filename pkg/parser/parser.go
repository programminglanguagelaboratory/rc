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
	expr, err := p.parseBinaryExpr(0)
	if err != nil {
		return nil, err
	}

	if p.tok.Kind == token.COLONEQUALS {
		return p.parseDeclExpr(expr)
	}

	return expr, nil
}

func (p *Parser) parseDeclExpr(name ast.Expr) (ast.Expr, error) {
	var ok bool
	var l ast.LitExpr
	if l, ok = name.(ast.LitExpr); !ok {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	if l.Token.Kind != token.ID {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	id := ast.Id(l.Token.Text)

	if p.tok.Kind != token.COLONEQUALS {
		return nil, fmt.Errorf("expected :=, but got: %v", p.tok.Kind)
	}
	p.next()

	value, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if p.tok.Kind != token.SEMICOLON {
		return nil, fmt.Errorf("expected ;, but got: %v", p.tok.Kind)
	}
	p.next()

	body, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return ast.DeclExpr{Name: id, Value: value, Body: body}, nil
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
	default:
		return p.parseCallExpr()
	}
}

func (p *Parser) parseCallExpr() (ast.Expr, error) {
	x, err := p.parsePrimaryExpr()
	if err != nil {
		return nil, err
	}

	for {
		y, err := p.parsePrimaryExpr()
		if err != nil {
			break
		}

		x = ast.CallExpr{Func: x, Arg: y}
	}

	return x, nil
}

func (p *Parser) parsePrimaryExpr() (ast.Expr, error) {
	expr, err := p.parseLitOrParenExpr()
	if err != nil {
		return nil, err
	}

L:
	for {
		switch p.tok.Kind {
		case token.DOT:
			expr, err = p.parseFieldExpr(expr)
			if err != nil {
				return nil, err
			}
		default:
			break L
		}
	}

	return expr, nil
}

func (p *Parser) parseFieldExpr(left ast.Expr) (ast.Expr, error) {
	if p.tok.Kind != token.DOT {
		return nil, fmt.Errorf("expected dot, but got: %v", p.tok.Kind)
	}
	p.next()

	right := p.tok.Text
	if p.tok.Kind != token.ID {
		return nil, fmt.Errorf("expected id, but got: %v", p.tok.Kind)
	}
	p.next()

	return ast.FieldExpr{Left: left, Right: ast.Id(right)}, nil
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
