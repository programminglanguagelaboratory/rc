package typcheck

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/typ"
)

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
