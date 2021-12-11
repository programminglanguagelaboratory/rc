package typcheck

import (
	"errors"
	"fmt"
	"strings"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context struct {
	tvs    map[string]*scheme
	lastId int
}

func (c *context) GenId() string {
	c.lastId++
	return fmt.Sprintf("%d", c.lastId)
}

func (c *context) Apply(subst Subst) Substitutable {
	for tv, s := range c.tvs {
		c.tvs[tv] = s.Apply(subst).(*scheme)
	}
	return c
}

func (c *context) FreeTypeVars() []string {
	// panic("not impl")
	return nil
}

type scheme struct {
	tvs []string
	t   inferTyp
}

func (c *context) instantiateScheme(s *scheme) inferTyp {
	// panic("not impl")
	return nil
	// instantiateSubst := Subst{}
	// for _, tv := range s.tvs {
	//   instantiateSubst[tv] = &varTyp{c.GenId()}
	// }
	// return s.Apply(instantiateSubst).(*scheme).t
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
	c.tvs = make(map[string]*scheme)
	inferTyp, _, err := c.inferExpr(e)
	if err != nil {
		return nil, err
	}

	switch inferTyp := inferTyp.(type) {
	case *constTyp:
		return inferTyp.t, nil
	case *funcTyp:
		return typ.NewFunc(), nil
	default:
		return nil, errors.New("not impl")
	}
}

func (c *context) inferExpr(e ast.Expr) (inferTyp, Subst, error) {
	switch e := e.(type) {
	case ast.DeclExpr:
		valueTyp, valueSubst, err := c.inferExpr(e.Value)
		c.tvs[string(e.Name)] = &scheme{t: valueTyp}
		c.Apply(valueSubst)
		bodyTyp, bodySubst, err := c.inferExpr(e.Body)
		if err != nil {
			return nil, nil, err
		}
		return bodyTyp, composeSubst(valueSubst, bodySubst), nil
	case ast.CallExpr:
		return nil, nil, errors.New("not impl")
	case ast.IdentExpr:
		s, ok := c.tvs[e.Value]
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
	case ast.FuncLitExpr:
		nameTypVar := c.GenId()
		nameTyp := inferTyp(&varTyp{nameTypVar})
		c.tvs[nameTypVar] = &scheme{t: nameTyp}
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
