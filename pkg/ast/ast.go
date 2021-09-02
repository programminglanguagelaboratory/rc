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
	Left  Expr
	Right Expr
	Token token.Token
}

type UnaryExpr struct {
	Left  Expr
	Token token.Token
}

type CallExpr struct {
	Func Expr
	Args []Expr
}

type FieldExpr struct {
	Left  Expr
	Right Id
}

type LitExpr struct {
	Token token.Token
}

type FuncLitExpr struct {
	Name  Id
	Param Expr
	Body  Expr
}
