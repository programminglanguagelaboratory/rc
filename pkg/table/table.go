package table

import (
	"github.com/programminglanguagelaboratory/rc/pkg/symbol"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type Table map[symbol.Symbol]typ.Typ

func (t *Table) FindType(s symbol.Symbol) (typ.Typ, bool) {
	typ, ok := (*t)[s]
	return typ, ok
}
