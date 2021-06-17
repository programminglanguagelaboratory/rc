package parser

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
)

type parser struct {
	lexer lexer.Lexer
}

func newParser(l *lexer.Lexer) *parser {
	p := parser{}
	return &p
}

func (p *parser) parseExpr() (ast.Expr, error) {
	return nil, errors.New("not implemented")
}
