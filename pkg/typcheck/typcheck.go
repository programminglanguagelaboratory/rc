package typcheck

import (
	"errors"
	"strings"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context map[string]scheme

func (c context) Apply(subst Subst) Substitutable {
	for tv, s := range c {
		c[tv] = s.Apply(subst).(scheme)
	}
	return Substitutable(c)
}

func (c context) FreeTypeVars() []string {
	panic("not impl")
}

type scheme struct {
	tvs []string
	t   inferTyp
}

func (s scheme) Apply(subst Subst) Substitutable {
	freeSubst := Subst{}
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

func (s scheme) String() string {
	var out strings.Builder
	if s.tvs != nil {
		out.WriteString("forall")
		for _, tv := range s.tvs {
			out.WriteString(" ")
			out.WriteString(tv)
		}
		out.WriteString(".")
	}
	out.WriteString(s.t.String())
	return out.String()
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
