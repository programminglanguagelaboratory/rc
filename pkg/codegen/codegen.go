package codegen

import (
	"errors"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
)

type Codegen struct {
	mod *ir.Module
	fun *ir.Func
	blk *ir.Block
}

func NewCodegen() *Codegen {
	c := Codegen{}
	c.mod = ir.NewModule()
	c.fun = c.mod.NewFunc("", types.I64)
	c.blk = c.fun.NewBlock("")
	return &c
}

func (c *Codegen) EmitExpr(e *ast.Expr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) EmitBinaryExpr(e *ast.BinaryExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) EmitUnaryExpr(e *ast.UnaryExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) EmitCallExpr(e *ast.CallExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) EmitFieldExpr(e *ast.FieldExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) EmitLitExpr(e *ast.LitExpr) interface{} {
	return errors.New("not implemented")
}
