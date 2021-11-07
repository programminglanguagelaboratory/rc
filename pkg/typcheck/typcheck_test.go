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
	}{
		{"\"hello\"", typ.NewString()},
		{"10", typ.NewNumber()},
		{"true", typ.NewBool()},

		{"x := 10; x", typ.NewNumber()},
		{"x := 10; y:= \"hello\"; x", typ.NewNumber()},
		{"x := 10; \"hello\"", typ.NewString()},

		{"10 + 20", typ.NewString()},
	} {
		expr, err := parser.NewParser(lexer.NewLexer(testcase.code), nil).Parse()
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v",
				testcase.code,
				testcase.expected,
				err,
			)
			continue
		}

		desugared := desugar.Desugar(expr)
		actual, err := Infer(desugared)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v",
				testcase.code,
				testcase.expected,
				err,
			)
			continue
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
