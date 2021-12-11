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
	return &constTyp{c}
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
