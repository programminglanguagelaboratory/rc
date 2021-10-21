package codegen

import (
	"errors"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
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
	case ast.DeclExpr:
		return c.genDeclExpr(v)
	case ast.BinaryExpr:
		return c.genBinaryExpr(v)
	case ast.IdentExpr:
		value, ok := c.decls[ast.Id(v.Value)]
		if !ok {
			return nil, errors.New("undefined variable: " + v.Value)
		}
		return value, nil
	case ast.NumberExpr:
		return constant.NewInt(types.I64, v.Value), nil
	}
	return nil, errors.New("not implemented")
}

func (c *Codegen) genDeclExpr(e ast.DeclExpr) (value.Value, error) {
	var err error
	c.decls[e.Name], err = c.genExpr(e.Value)
	if err != nil {
		return nil, err
	}

	return c.genExpr(e.Body)
}

func (c *Codegen) genBinaryExpr(e ast.BinaryExpr) (value.Value, error) {
	left, err := c.genExpr(e.Left)
	if err != nil {
		return nil, err
	}

	right, err := c.genExpr(e.Right)
	if err != nil {
		return nil, err
	}

	t := e.Token
	switch t.Kind {
	case token.PLUS:
		return c.blk.NewAdd(left, right), nil
	case token.MINUS:
		return c.blk.NewSub(left, right), nil
	case token.ASTERISK:
		return c.blk.NewMul(left, right), nil
	case token.SLASH:
		return c.blk.NewSDiv(left, right), nil
	}

	return nil, errors.New("not implemented")
}

func (c *Codegen) genUnaryExpr(e ast.UnaryExpr) (value.Value, error) {
	return nil, errors.New("not implemented")
}

func (c *Codegen) genCallExpr(e ast.CallExpr) (value.Value, error) {
	return nil, errors.New("genCallExpr: not implemented")
}

func (c *Codegen) genFieldExpr(e ast.FieldExpr) (value.Value, error) {
	return nil, errors.New("not implemented")
}
