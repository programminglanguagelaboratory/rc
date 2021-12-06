package typcheck

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

func TestApply(t *testing.T) {
	for _, testcase := range []struct {
		t        inferTyp
		subst    Subst
		expected inferTyp
	}{
		{
			inferTyp(&varTyp{tv: "a"}),
			map[string]typ.Typ{},
			inferTyp(&varTyp{tv: "a"}),
		},
		{
			inferTyp(&varTyp{tv: "a"}),
			map[string]typ.Typ{"a": typ.NewBool()},
			inferTyp(&constTyp{t: typ.NewBool()}),
		},
		{
			inferTyp(&funcTyp{from: &varTyp{tv: "a"}, to: &varTyp{tv: "b"}}),
			map[string]typ.Typ{"a": typ.NewBool(), "b": typ.NewNumber()},
			inferTyp(&funcTyp{from: &constTyp{t: typ.NewBool()}, to: &constTyp{t: typ.NewNumber()}}),
		},
	} {
		actual := testcase.t.Apply(testcase.subst)
		if !reflect.DeepEqual(testcase.expected, actual) {
			t.Errorf(
				"given %v, %v, expected %v, but got actual %v",
				testcase.t,
				testcase.subst,
				testcase.expected,
				actual,
			)
		}
	}
}

func TestFreeTypeVars(t *testing.T) {
	for _, testcase := range []struct {
		t        inferTyp
		expected []string
	}{
		{
			inferTyp(&varTyp{tv: "a"}),
			[]string{"a"},
		},
		{
			inferTyp(&constTyp{t: typ.NewBool()}),
			nil,
		},
		{
			inferTyp(&funcTyp{from: &varTyp{tv: "a"}, to: &varTyp{tv: "b"}}),
			[]string{"a", "b"},
		},
	} {
		actual := make(map[string]struct{})
		for _, tv := range testcase.t.FreeTypeVars() {
			actual[tv] = struct{}{}
		}
		expected := make(map[string]struct{})
		for _, tv := range testcase.expected {
			expected[tv] = struct{}{}
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.t,
				testcase.expected,
				actual,
			)
		}
	}
}
