package table

import (
	"github.com/programminglanguagelaboratory/rc/pkg/symbol"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type Table map[symbol.Symbol]typ.Typ

func NewTable() Table {
	return make(Table)
}

func (t *Table) FindType(s symbol.Symbol) (typ.Typ, bool) {
	typ, ok := (*t)[s]
	return typ, ok
}

func (t *Table) AddType(s symbol.Symbol, typ typ.Typ) {
	(*t)[s] = typ
}
