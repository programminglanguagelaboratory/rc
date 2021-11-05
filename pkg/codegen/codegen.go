package codegen

import (
	"errors"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
)

type Codegen struct {
	mod   *ir.Module
	fun   *ir.Func
	blk   *ir.Block
	decls map[ast.Id]value.Value
}

func NewCodegen() *Codegen {
	c := Codegen{}
	c.mod = ir.NewModule()
	c.fun = c.mod.NewFunc("main", types.I64)
	c.blk = c.fun.NewBlock("entry")
	c.decls = make(map[ast.Id]value.Value)
	return &c
}

func (c *Codegen) Gen(e ast.Expr) (*ir.Module, error) {
	ret, err := c.genExpr(e)
	if err != nil {
		return nil, err
	}

	c.blk.NewRet(ret)
	return c.mod, nil
}

func (c *Codegen) genExpr(e ast.Expr) (value.Value, error) {
	switch v := interface{}(e).(type) {
	case ast.BinaryExpr:
	case ast.UnaryExpr:
		return nil, errors.New("unexpected sugared expression")
	case ast.DeclExpr:
		return c.genDeclExpr(v)
	case ast.IdentExpr:
		p, ok := c.decls[ast.Id(v.Value)]
		if !ok {
			return nil, errors.New("undefined variable: " + v.Value)
		}
		value := c.blk.NewLoad(types.I64, p)
		return value, nil
	case ast.NumberExpr:
		return constant.NewInt(types.I64, v.Value), nil
	}
	return nil, errors.New("not implemented")
}

func (c *Codegen) genDeclExpr(e ast.DeclExpr) (value.Value, error) {
	value, err := c.genExpr(e.Value)
	if err != nil {
		return nil, err
	}

	p := c.blk.NewAlloca(types.I64)
	c.blk.NewStore(value, p)
	c.decls[e.Name] = p
	return c.genExpr(e.Body)
}

func (c *Codegen) genCallExpr(e ast.CallExpr) (value.Value, error) {
	return nil, errors.New("genCallExpr: not implemented")
}

func (c *Codegen) genFieldExpr(e ast.FieldExpr) (value.Value, error) {
	return nil, errors.New("not implemented")
}
