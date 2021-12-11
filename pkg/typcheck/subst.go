package typcheck

type Subst map[string]inferTyp

func composeSubst(x, y Subst) Subst {
	s := Subst{}
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
	FreeTypeVars() []string
}
