package typcheck

import (
	"reflect"
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

		{"10 + 20", typ.NewNumber()},
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

func TestSchemeApply(t *testing.T) {
	for _, testcase := range []struct {
		s        *scheme
		subst    Subst
		expected *scheme
	}{
		{
			&scheme{nil, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			map[string]typ.Typ{"a": typ.NewBool()},
			&scheme{nil, &funcTyp{from: &constTyp{typ.NewBool()}, to: &varTyp{"b"}}},
		},
		{
			&scheme{[]string{"a"}, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			map[string]typ.Typ{"a": typ.NewBool()},
			&scheme{[]string{"a"}, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
		},
	} {
		actual := testcase.s.Apply(testcase.subst)
		if !reflect.DeepEqual(testcase.expected, actual) {
			t.Errorf(
				"given %v, %v, expected %v, but got actual %v",
				testcase.s,
				testcase.subst,
				testcase.expected,
				actual,
			)
		}
	}
}

func TestSchemeFreeTypeVars(t *testing.T) {
	for _, testcase := range []struct {
		s        *scheme
		expected []string
	}{
		{
			&scheme{nil, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			[]string{"a", "b"},
		},
		{
			&scheme{[]string{"a"}, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			[]string{"b"},
		},
		{
			&scheme{[]string{"b"}, &funcTyp{from: &varTyp{"a"}, to: &funcTyp{from: &varTyp{"b"}, to: &varTyp{"a"}}}},
			[]string{"a"},
		},
	} {
		actual := make(map[string]struct{})
		for _, tv := range testcase.s.FreeTypeVars() {
			actual[tv] = struct{}{}
		}
		expected := make(map[string]struct{})
		for _, tv := range testcase.expected {
			expected[tv] = struct{}{}
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.s,
				testcase.expected,
				actual,
			)
		}
	}
}
