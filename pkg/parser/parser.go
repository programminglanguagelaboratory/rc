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

type parser struct {
	lexer *lexer.Lexer
	tok   token.Token
	table *table.Table
}

func (p *parser) next() {
	p.tok, _ = p.lexer.Lex()
}

func Parse(code string) (ast.Expr, error) {
	p := parser{}
	p.lexer = lexer.NewLexer(code)
	p.tok, _ = p.lexer.Lex()
	p.table = nil
	return p.parseExpr()
}

func (p *parser) parseExpr() (ast.Expr, error) {
	expr, err := p.parseBinaryExpr(0)
	if err != nil {
		return nil, err
	}

	switch p.tok.Kind {
	case token.EQUALSGREATERTHAN:
		return p.parseFuncLitExpr(expr)
	case token.COLONEQUALS:
		return p.parseDeclExpr(expr)
	default:
		return expr, nil
	}
}

func (p *parser) parseFuncLitExpr(e ast.Expr) (ast.Expr, error) {
	x, ok := e.(ast.IdentExpr)
	if !ok {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	if x.Token.Kind != token.ID {
		return nil, fmt.Errorf("expected lhs, but got: %v", p.tok.Kind)
	}
	id := ast.Id(x.Value)

	if p.tok.Kind != token.EQUALSGREATERTHAN {
		return nil, fmt.Errorf("expected =>, but got: %v", p.tok.Kind)
	}
	p.next()

	body, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	return ast.FuncExpr{Name: id, Body: body}, nil
}

func (p *parser) parseDeclExpr(e ast.Expr) (ast.Expr, error) {
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

func (p *parser) parseBinaryExpr(prevPrec int) (ast.Expr, error) {
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

func (p *parser) parseUnaryExpr() (ast.Expr, error) {
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

func (p *parser) parseCallExpr() (ast.Expr, error) {
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

func (p *parser) parsePrimaryExpr() (ast.Expr, error) {
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

func (p *parser) parseFieldExpr(x ast.Expr) (ast.Expr, error) {
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

func (p *parser) parseLitOrParenExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.ID:
		e := ast.IdentExpr{Token: p.tok, Value: p.tok.Text}
		p.next()
		return e, nil
	case token.STRING:
		e := ast.StringExpr{Token: p.tok, Value: p.tok.Text}
		p.next()
		return e, nil
	case token.NUMBER:
		v, _ := strconv.ParseInt(p.tok.Text, 0, 64)
		e := ast.NumberExpr{Token: p.tok, Value: v}
		p.next()
		return e, nil
	case token.BOOL:
		v, _ := strconv.ParseBool(p.tok.Text)
		e := ast.BoolExpr{Token: p.tok, Value: v}
		p.next()
		return e, nil
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
