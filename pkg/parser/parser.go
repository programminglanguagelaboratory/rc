package parser

import (
	"errors"

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
	return p.parsePrimaryExpr()
}

func (p *parser) parsePrimaryExpr() (ast.Expr, error) {
	switch p.tok.Kind {
	case token.ID, token.STRING, token.NUMBER:
		t := ast.LitExpr{Token: p.tok}
		p.next()
		return t, nil
	case token.LPAREN:
		p.next()
		return nil, errors.New("not implemented")
	}
	return nil, errors.New("unexpected token")
}
