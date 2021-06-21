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

func (c *Codegen) genExpr(e *ast.Expr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) genBinaryExpr(e *ast.BinaryExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) genUnaryExpr(e *ast.UnaryExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) genCallExpr(e *ast.CallExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) genFieldExpr(e *ast.FieldExpr) interface{} {
	return errors.New("not implemented")
}

func (c *Codegen) genLitExpr(e *ast.LitExpr) interface{} {
	return errors.New("not implemented")
}
