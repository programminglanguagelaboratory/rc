package typcheck

import "github.com/programminglanguagelaboratory/rc/pkg/typ"

type Subst map[string]typ.Typ

func composeSubst(x, y Subst) Subst {
	s := make(Subst)
	for k, v := range y {
		s[k] = v
	}
	for k, v := range x {
		s[k] = v
	}
	return s
}

type Substitutable interface {
	Apply(Subst) Substitutable
}
