package typcheck

import (
	"errors"
	"fmt"
	"strings"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context struct {
	tvs    map[string]scheme
	lastId int
}

func (c *context) GenId() string {
	c.lastId++
	return fmt.Sprintf("%d", c.lastId)
}

func (c *context) Apply(subst Subst) Substitutable {
	for tv, s := range c.tvs {
		c.tvs[tv] = s.Apply(subst).(scheme)
	}
	return Substitutable(c)
}

func (c *context) FreeTypeVars() []string {
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
	c := context{}
	c.tvs = make(map[string]scheme)
	inferTyp, _, err := c.inferExpr(e)
	if err != nil {
		return nil, err
	}

	typ := inferTyp.(*constTyp)
	return typ.t, nil
}

func (c *context) inferExpr(e ast.Expr) (inferTyp, Subst, error) {
	switch e := e.(type) {
	case ast.DeclExpr:
		return nil, nil, errors.New("not impl")
	case ast.CallExpr:
		return nil, nil, errors.New("not impl")
	case ast.IdentExpr:
		return nil, nil, errors.New("not impl")
	case ast.StringExpr:
		return &constTyp{typ.NewString()}, nil, nil
	case ast.NumberExpr:
		return &constTyp{typ.NewNumber()}, nil, nil
	case ast.BoolExpr:
		return &constTyp{typ.NewBool()}, nil, nil
	case ast.FuncLitExpr:
		nameTypVar := c.GenId()
		nameTyp := inferTyp(&varTyp{nameTypVar})
		c.tvs[nameTypVar] = scheme{t: nameTyp}
		bodyTyp, bodySubst, err := c.inferExpr(e.Body)
		if err != nil {
			return nil, nil, err
		}
		nameTyp = nameTyp.Apply(bodySubst).(inferTyp)
		return &funcTyp{from: nameTyp, to: bodyTyp}, bodySubst, nil
	default:
		return nil, nil, errors.New("unexpected sugared expression")
	}
}
