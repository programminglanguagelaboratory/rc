package typcheck

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/desugar"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

func TestInferExpr(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected typ.Typ
	}{} {
		expr, err := parser.NewParser(lexer.NewLexer(testcase.code), nil).Parse()
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v",
				testcase.code,
				testcase.expected,
				err)
		}

		desugared := desugar.Desugar(expr)
		actual, nil := Infer(desugared)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v",
				testcase.code,
				testcase.expected,
				err)
		}

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
