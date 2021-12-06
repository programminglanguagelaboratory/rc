package typcheck

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context struct{}

func (c *context) Apply(s Subst) Substitutable {
	panic("not impl")
}

func (c *context) FreeTypeVars() Substitutable {
	panic("not impl")
}

type scheme struct {
	tvs []string
	t   inferTyp
}

func (s scheme) Apply(subst Subst) Substitutable {
	var freeSubst Subst
	for tv, t := range subst {
		freeSubst[tv] = t
	}
	for _, tv := range s.tvs {
		delete(freeSubst, tv)
	}

	return Substitutable(&scheme{
		tvs: s.tvs,
		t:   s.t.Apply(freeSubst).(inferTyp),
	})
}

func (s scheme) FreeTypeVars() []string {
	tvs := make(map[string]struct{})
	for _, tv := range s.t.FreeTypeVars() {
		tvs[tv] = struct{}{}
	}
	for _, tv := range s.tvs {
		delete(tvs, tv)
	}

	ret := make([]string, len(tvs))
	i := 0
	for tv := range tvs {
		ret[i] = tv
		i++
	}
	return ret
}

func Infer(e ast.Expr) (typ.Typ, error) {
	return inferExpr(e)
}

func inferExpr(e ast.Expr) (typ.Typ, error) {
	switch e.(type) {
	case ast.StringExpr:
		return typ.NewString(), nil
	case ast.NumberExpr:
		return typ.NewNumber(), nil
	case ast.BoolExpr:
		return typ.NewBool(), nil
	default:
		return nil, errors.New("not impl")
	}
}
