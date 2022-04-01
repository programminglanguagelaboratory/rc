package typcheck

import (
	"errors"
	"fmt"
	"strings"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context struct {
	schemes map[string]*scheme
	lastId  int
}

func (c *context) GenId() string {
	c.lastId++
	return fmt.Sprintf("%d", c.lastId)
}

func (c *context) Apply(subst Subst) Substitutable {
	for tv, s := range c.schemes {
		c.schemes[tv] = s.Apply(subst).(*scheme)
	}
	return c
}

func (c *context) FreeTypeVars() []string {
	panic("not impl")
}

type scheme struct {
	tvs []string
	t   inferTyp
}

func (c *context) instantiateScheme(s *scheme) inferTyp {
	instantiateSubst := Subst{}
	for _, tv := range s.tvs {
		instantiateSubst[tv] = &varTyp{c.GenId()}
	}
	return s.Apply(instantiateSubst).(*scheme).t
}

func (s *scheme) Apply(subst Subst) Substitutable {
	freeSubst := Subst{}
	for tv, t := range subst {
		freeSubst[tv] = t
	}
	for _, tv := range s.tvs {
		delete(freeSubst, tv)
	}

	return &scheme{
		tvs: s.tvs,
		t:   s.t.Apply(freeSubst).(inferTyp),
	}
}

func (s *scheme) FreeTypeVars() []string {
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

func (s *scheme) String() string {
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
	c := &context{}
	c.schemes = make(map[string]*scheme)
	inferTyp, _, err := c.inferExpr(e)
	if err != nil {
		return nil, err
	}

	switch inferTyp := inferTyp.(type) {
	case *constTyp:
		return inferTyp.t, nil
	case *varTyp:
		return typ.NewVar(inferTyp.tv), nil
	case *funcTyp:
		return typ.NewFunc(inferTyp.from, inferTyp.to), nil
	default:
		return nil, errors.New("unreachable")
	}
}

func (c *context) inferExpr(e ast.Expr) (inferTyp, Subst, error) {
	switch e := e.(type) {
	case ast.DeclExpr:
		valueTyp, valueSubst, err := c.inferExpr(e.Value)
		if err != nil {
			return nil, nil, err
		}

		c.schemes[string(e.Name)] = &scheme{t: valueTyp}
		c.Apply(valueSubst)
		bodyTyp, bodySubst, err := c.inferExpr(e.Body)
		if err != nil {
			return nil, nil, err
		}
		return bodyTyp, composeSubst(valueSubst, bodySubst), nil
	case ast.CallExpr:
		t0, s0, err := c.inferExpr(e.Func)
		if err != nil {
			return nil, nil, err
		}

		t1, s1, err := c.inferExpr(e.Arg)
		if err != nil {
			return nil, nil, err
		}

		tv := c.GenId()
		t := inferTyp(&varTyp{tv})
		s2, err := unify(t0.Apply(s1).(inferTyp), &funcTyp{from: t1, to: t})
		if err != nil {
			return nil, nil, err
		}

		return t.Apply(s2).(inferTyp), composeSubst(s2, composeSubst(s1, s0)), nil
	case ast.IdentExpr:
		s, ok := c.schemes[e.Value]
		if !ok {
			return nil, nil, fmt.Errorf("undefined variable: %s", s)
		}
		return c.instantiateScheme(s), nil, nil
	case ast.StringExpr:
		return &constTyp{typ.NewString()}, nil, nil
	case ast.NumberExpr:
		return &constTyp{typ.NewNumber()}, nil, nil
	case ast.BoolExpr:
		return &constTyp{typ.NewBool()}, nil, nil
	case ast.FuncExpr:
		nameTypVar := c.GenId()
		nameTyp := &varTyp{nameTypVar}
		c.schemes[string(e.Name)] = &scheme{t: nameTyp}
		bodyTyp, bodySubst, err := c.inferExpr(e.Body)
		if err != nil {
			return nil, nil, err
		}
		return &funcTyp{nameTyp.Apply(bodySubst).(inferTyp), bodyTyp}, bodySubst, nil
	default:
		return nil, nil, errors.New("unexpected sugared expression")
	}
}
