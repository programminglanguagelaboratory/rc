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

type varTyp struct {
	tv string
}

func (t *varTyp) Apply(s Subst) Substitutable {
	c, ok := map[string]typ.Typ(s)[t.tv]
	if !ok {
		return Substitutable(t)
	}
	return Substitutable(&constTyp{t: c})
}
func (t *varTyp) String() string { return t.tv }
func (t *varTyp) inferType()     {}

type constTyp struct {
	t typ.Typ
}

func (t *constTyp) Apply(s Subst) Substitutable {
	return Substitutable(t)
}
func (t *constTyp) String() string { return t.t.String() }
func (t *constTyp) inferType()     {}

type funcTyp struct {
	from inferTyp
	to   inferTyp
}

func (t *funcTyp) Apply(s Subst) Substitutable {
	t.from = t.from.Apply(s).(inferTyp)
	t.to = t.to.Apply(s).(inferTyp)
	return Substitutable(t)
}
func (t *funcTyp) inferType()     {}
func (t *funcTyp) String() string { return t.from.String() + " -> " + t.to.String() }
