package typcheck

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

func TestInferTypApply(t *testing.T) {
	for _, tt := range []struct {
		t        inferTyp
		subst    Subst
		expected inferTyp
	}{
		{
			&varTyp{"a"},
			nil,
			&varTyp{"a"},
		},
		{
			&varTyp{"a"},
			Subst{"a": &constTyp{typ.NewBool()}},
			&constTyp{typ.NewBool()},
		},
		{
			&funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}},
			Subst{"a": &constTyp{typ.NewBool()}, "b": &constTyp{typ.NewNumber()}},
			&funcTyp{from: &constTyp{typ.NewBool()}, to: &constTyp{typ.NewNumber()}},
		},
	} {
		actual := tt.t.Apply(tt.subst)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, %v, expected %v, but got actual %v", tt.t, tt.subst, tt.expected, actual)
		}
	}
}

func TestInferTypFreeTypeVars(t *testing.T) {
	for _, tt := range []struct {
		t        inferTyp
		expected []string
	}{
		{
			&varTyp{"a"},
			[]string{"a"},
		},
		{
			&constTyp{typ.NewBool()},
			nil,
		},
		{
			&funcTyp{from: &varTyp{"a"}, to: &varTyp{"b"}},
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

func TestUnify(t *testing.T) {
	for _, tt := range []struct {
		t0       inferTyp
		t1       inferTyp
		expected Subst
	}{
		{
			&constTyp{t: typ.NewBool()},
			&constTyp{t: typ.NewBool()},
			nil,
		},
		{
			&varTyp{"a"},
			&varTyp{"a"},
			nil,
		},
		{
			&varTyp{"a"},
			&constTyp{t: typ.NewBool()},
			Subst{"a": &constTyp{typ.NewBool()}},
		},
		{
			&funcTyp{&varTyp{"a"}, &constTyp{typ.NewBool()}},
			&funcTyp{&varTyp{"b"}, &constTyp{typ.NewBool()}},
			Subst{"a": &varTyp{"b"}},
		},
	} {
		actual, err := unify(tt.t0, tt.t1)
		if err != nil {
			t.Errorf("given %v and %v, expected %v, but got an error %v", tt.t0, tt.t1, tt.expected, err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v and %v, expected %v, but got actual %v", tt.t0, tt.t1, tt.expected, actual)
		}
	}
}
