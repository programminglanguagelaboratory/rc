package table

import (
	"github.com/programminglanguagelaboratory/rc/pkg/token"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type Symbol struct {
	Tok token.Token
}

type Table map[Symbol]typ.Typ

func NewTable() Table {
	return make(Table)
}

func (t *Table) FindType(s Symbol) (typ.Typ, bool) {
	typ, ok := (*t)[s]
	return typ, ok
}

func (t *Table) AddType(s Symbol, typ typ.Typ) {
	(*t)[s] = typ
}
