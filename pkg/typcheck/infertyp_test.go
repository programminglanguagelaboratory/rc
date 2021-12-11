package typcheck

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

func TestApply(t *testing.T) {
	for _, tt := range []struct {
		t        inferTyp
		subst    Subst
		expected inferTyp
	}{
		{
			inferTyp(&varTyp{tv: "a"}),
			nil,
			inferTyp(&varTyp{tv: "a"}),
		},
		{
			inferTyp(&varTyp{tv: "a"}),
			Subst{"a": &constTyp{typ.NewBool()}},
			inferTyp(&constTyp{t: typ.NewBool()}),
		},
		{
			inferTyp(&funcTyp{from: &varTyp{tv: "a"}, to: &varTyp{tv: "b"}}),
			Subst{"a": &constTyp{typ.NewBool()}, "b": &constTyp{typ.NewNumber()}},
			inferTyp(&funcTyp{from: &constTyp{t: typ.NewBool()}, to: &constTyp{t: typ.NewNumber()}}),
		},
	} {
		actual := tt.t.Apply(tt.subst)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, %v, expected %v, but got actual %v", tt.t, tt.subst, tt.expected, actual)
		}
	}
}

func TestFreeTypeVars(t *testing.T) {
	for _, tt := range []struct {
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
		for _, tv := range tt.t.FreeTypeVars() {
			actual[tv] = struct{}{}
		}
		expected := make(map[string]struct{})
		for _, tv := range tt.expected {
			expected[tv] = struct{}{}
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("given %v, expected %v, but got actual %v", tt.t, tt.expected, actual)
		}
	}
}
