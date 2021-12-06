package typcheck

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

type context struct {
	Substitutable
}

func (c *context) Apply() Substitutable {
	panic("not impl")
}

func (c *context) FreeTypeVars() Substitutable {
	panic("not impl")
}

func Infer(e ast.Expr) (typ.Typ, error) {
	return inferExpr(e)
}

func inferExpr(e ast.Expr) (typ.Typ, error) {
	switch e.(type) {
	case ast.StringExpr:
		return typ.NewString(), nil
	case ast.NumberExpr:
		return typ.NewNumber(), nil
	case ast.BoolExpr:
		return typ.NewBool(), nil
	default:
		return nil, errors.New("not impl")
	}
}
