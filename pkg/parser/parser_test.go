package parser

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
)

func TestExpr(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected ast.Expr
	}{} {
		actual, _ := newParser().parseExpr()
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
