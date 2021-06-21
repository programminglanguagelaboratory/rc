package codegen

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
)

func EmitExpr(e *ast.Expr) interface{} {
	return errors.New("not implemented")
}

func EmitBinaryExpr(e *ast.BinaryExpr) interface{} {
	return errors.New("not implemented")
}

func EmitUnaryExpr(e *ast.UnaryExpr) interface{} {
	return errors.New("not implemented")
}

func EmitCallExpr(e *ast.CallExpr) interface{} {
	return errors.New("not implemented")
}

func EmitFieldExpr(e *ast.FieldExpr) interface{} {
	return errors.New("not implemented")
}

func EmitLitExpr(e *ast.LitExpr) interface{} {
	return errors.New("not implemented")
}
