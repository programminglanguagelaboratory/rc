package typcheck

import (
	"fmt"

	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type inferTyp interface {
	Substitutable
	fmt.Stringer
	inferType()
}

type constTyp struct {
	t typ.Typ
}

func (t *constTyp) Apply(s Subst) Substitutable { return t }
func (t *constTyp) FreeTypeVars() []string      { return nil }
func (t *constTyp) String() string              { return t.t.String() }
func (t *constTyp) inferType()                  {}

type varTyp struct {
	tv string
}

func (t *varTyp) Apply(s Subst) Substitutable {
	c, ok := s[t.tv]
	if !ok {
		return t
	}
	return c
}
func (t *varTyp) FreeTypeVars() []string { return []string{t.tv} }
func (t *varTyp) String() string         { return t.tv }
func (t *varTyp) inferType()             {}

type funcTyp struct {
	from inferTyp
	to   inferTyp
}

func (t *funcTyp) Apply(s Subst) Substitutable {
	return &funcTyp{
		from: t.from.Apply(s).(inferTyp),
		to:   t.to.Apply(s).(inferTyp),
	}
}
func (t *funcTyp) FreeTypeVars() []string {
	tvs := make(map[string]struct{})
	for _, tv := range t.from.FreeTypeVars() {
		tvs[tv] = struct{}{}
	}
	for _, tv := range t.to.FreeTypeVars() {
		tvs[tv] = struct{}{}
	}

	ret := make([]string, len(tvs))
	i := 0
	for tv := range tvs {
		ret[i] = tv
		i++
	}
	return ret
}
func (t *funcTyp) inferType()     {}
func (t *funcTyp) String() string { return t.from.String() + " -> " + t.to.String() }

func unify(t0, t1 inferTyp) (Subst, error) {
	c0, ok0 := t0.(*constTyp)
	c1, ok1 := t1.(*constTyp)
	if ok0 && ok1 && c0.t == c1.t {
		return nil, nil
	}

	f0, ok0 := t0.(*funcTyp)
	f1, ok1 := t1.(*funcTyp)
	if ok0 && ok1 {
		s0, err := unify(f0.from, f1.from)
		if err != nil {
			return nil, err
		}

		s1, err := unify(f0.to.Apply(s0).(inferTyp), f1.to.Apply(s0).(inferTyp))
		if err != nil {
			return nil, err
		}

		return composeSubst(s1, s0), nil
	}

	v0, ok0 := t0.(*varTyp)
	v1, ok1 := t1.(*varTyp)
	if ok0 && ok1 && v0.tv == v1.tv {
		return nil, nil
	}

	if !(ok0 || ok1) {
		return nil, fmt.Errorf("unification failed: %v and %v", t0, t1)
	}

	var v *varTyp
	var t inferTyp
	if ok0 {
		v = v0
		t = t1
	} else {
		v = v1
		t = t0
	}

	for _, ftv := range t.FreeTypeVars() {
		if v.tv == ftv {
			return nil, fmt.Errorf("occurs check failed - infinite type: %v and %v", t0, t1)
		}
	}

	return Subst{v.tv: t}, nil
}
