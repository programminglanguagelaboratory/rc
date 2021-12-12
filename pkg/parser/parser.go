package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/table"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

type Parser struct {
	lexer *lexer.Lexer
	tok   token.Token
	table *table.Table
}

func (p *Parser) next() {
	p.tok, _ = p.lexer.Lex()
}

func NewParser(l *lexer.Lexer, t *table.Table) *Parser {
	p := Parser{}
	p.lexer = l
	p.tok, _ = p.lexer.Lex()
	p.table = t
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

func (p *Parser) parseDeclExpr(e ast.Expr) (ast.Expr, error) {
	x, ok := e.(ast.IdentExpr)
	if !ok {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	if x.Token.Kind != token.ID {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	id := ast.Id(x.Token.Text)

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
	x, err := p.parseUnaryExpr()
	if err != nil {
		return nil, err
	}

	for {
		prec := p.tok.Kind.GetPrec()
		if prec < prevPrec {
			return x, nil
		}

		tok := p.tok
		p.next()

		y, err := p.parseBinaryExpr(prec + 1)
		if err != nil {
			return nil, err
		}

		x = ast.BinaryExpr{X: x, Y: y, Token: tok}
	}
}

func (p *Parser) parseUnaryExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.MINUS, token.EXCLAMATION:
		t := ast.UnaryExpr{Token: p.tok}
		p.next()

		x, err := p.parseUnaryExpr()
		if err != nil {
			return nil, err
		}

		t.X = x
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

func (p *Parser) parseFieldExpr(x ast.Expr) (ast.Expr, error) {
	if p.tok.Kind != token.DOT {
		return nil, fmt.Errorf("expected dot, but got: %v", p.tok.Kind)
	}
	p.next()

	y := p.tok.Text
	if p.tok.Kind != token.ID {
		return nil, fmt.Errorf("expected id, but got: %v", p.tok.Kind)
	}
	p.next()

	return ast.FieldExpr{X: x, Y: ast.Id(y)}, nil
}

func (p *Parser) parseLitOrParenExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.ID:
		t := ast.IdentExpr{Token: p.tok, Value: p.tok.Text}
		p.next()
		return t, nil
	case token.STRING:
		t := ast.StringExpr{Token: p.tok, Value: p.tok.Text}
		p.next()
		return t, nil
	case token.NUMBER:
		n, _ := strconv.ParseInt(p.tok.Text, 0, 64)
		t := ast.NumberExpr{Token: p.tok, Value: n}
		p.next()
		return t, nil
	case token.BOOL:
		b, _ := strconv.ParseBool(p.tok.Text)
		t := ast.BoolExpr{Token: p.tok, Value: b}
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
