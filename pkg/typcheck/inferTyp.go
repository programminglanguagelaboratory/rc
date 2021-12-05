package typcheck

import (
	"fmt"

	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type inferTyp interface {
	inferType()
	Substitutable
	fmt.Stringer
}

type varType struct {
	tv string
}

func (t *varType) inferType()     {}
func (t *varType) String() string { return t.tv }

type constType struct {
	t typ.Typ
}

func (t *constType) inferType()     {}
func (t *constType) String() string { return t.t.String() }

type funcTyp struct {
	from inferTyp
	to   inferTyp
}

func (t *funcTyp) inferType()     {}
func (t *funcTyp) String() string { return t.from.String() + " -> " + t.to.String() }
