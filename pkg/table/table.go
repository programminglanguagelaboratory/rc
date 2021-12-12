package table

import (
	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type Table struct {
	Outer *Table
	Inner *Table
	Id    ast.Id
	Typ   typ.Typ
}

func (t *Table) Resolve(id ast.Id) (typ.Typ, bool) {
	if t == nil {
		return nil, false
	}

	if t.Id == id {
		return t.Typ, true
	}

	return t.Outer.Resolve(id)
}

func (t *Table) Define(id ast.Id, typ typ.Typ) *Table {
	newTable := &Table{
		Outer: t,
		Inner: nil,
		Id:    id,
		Typ:   typ,
	}
	newTable.Outer.Inner = newTable
	return newTable
}
