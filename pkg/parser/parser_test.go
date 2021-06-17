package parser

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
)

func TestExpr(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected ast.Expr
	}{} {
		actual, _ := newParser(lexer.NewLexer(testcase.code)).parseExpr()
		if testcase.expected != actual {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.code,
				testcase.expected,
				actual,
			)
		}
	}
}
