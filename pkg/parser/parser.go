package parser

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
)

type parser struct {
}

func newParser() *parser {
	p := parser{}
	return &p
}

func (p *parser) parseExpr() (ast.Expr, error) {
	return nil, errors.New("not implemented")
}
