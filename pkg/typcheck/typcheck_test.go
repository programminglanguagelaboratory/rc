package typcheck

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/desugar"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

func TestInferExpr(t *testing.T) {
	for _, tt := range []struct {
		code     string
		expected typ.Typ
	}{
		{"x => x", typ.NewFunc(typ.NewVar("1"), typ.NewVar("1"))},
		{"\"hello\"", typ.NewString()},
		{"10", typ.NewNumber()},
		{"true", typ.NewBool()},

		{"x := 10; x", typ.NewNumber()},
		{"x := 10; y:= \"hello\"; x", typ.NewNumber()},
		{"x := 10; \"hello\"", typ.NewString()},

		{"f := x => 10; f 20", typ.NewNumber()},
		{"f := x => 10; f false", typ.NewNumber()},
		{"id := x => x; id false", typ.NewBool()},
		{"k := x => y => x; id := a => a; (k id 10) false", typ.NewBool()},
	} {
		expr, err := parser.Parse(tt.code)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v", tt.code, tt.expected, err)
			continue
		}

		desugared := desugar.Desugar(expr)
		actual, err := Infer(desugared)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v", tt.code, tt.expected, err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, expected \n%#v\n, but got actual \n%#v\n", tt.code, tt.expected, actual)
		}
	}
}

func TestSchemeApply(t *testing.T) {
	for _, tt := range []struct {
		s        *scheme
		subst    Subst
		expected *scheme
	}{
		{
			&scheme{nil, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			Subst{"a": &constTyp{typ.NewBool()}},
			&scheme{nil, &funcTyp{from: &constTyp{typ.NewBool()}, to: &varTyp{"b"}}},
		},
		{
			&scheme{[]string{"a"}, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
			Subst{"a": &constTyp{typ.NewBool()}},
			&scheme{[]string{"a"}, &funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}}},
		},
	} {
		actual := tt.s.Apply(tt.subst)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, %v, expected %v, but got actual %v", tt.s, tt.subst, tt.expected, actual)
		}
	}
}

func TestSchemeFreeTypeVars(t *testing.T) {
	for _, tt := range []struct {
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
		for _, tv := range tt.s.FreeTypeVars() {
			actual[tv] = struct{}{}
		}
		expected := make(map[string]struct{})
		for _, tv := range tt.expected {
			expected[tv] = struct{}{}
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("given %v, expected %v, but got actual %v", tt.s, tt.expected, actual)
		}
	}
}
