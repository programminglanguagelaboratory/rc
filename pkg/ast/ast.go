package ast

import "github.com/programminglanguagelaboratory/rc/pkg/token"

type Expr interface{}

type Id string

type Field struct {
	Typ  []Id
	Name Id
}

type DeclExpr struct {
	Name  Id
	Value Expr
	Body  Expr
}

type BinaryExpr struct {
	X     Expr
	Y     Expr
	Token token.Token
}

type UnaryExpr struct {
	X     Expr
	Token token.Token
}

type CallExpr struct {
	Func Expr
	Arg  Expr
}

type FieldExpr struct {
	X Expr
	Y Id
}

type IdentExpr struct {
	Token token.Token
	Value string
}

type StringExpr struct {
	Token token.Token
	Value string
}

type NumberExpr struct {
	Token token.Token
	Value int64
}

type BoolExpr struct {
	Token token.Token
	Value bool
}

type FuncLitExpr struct {
	Name  Id
	Param Expr
	Body  Expr
}
