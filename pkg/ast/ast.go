package ast

import "github.com/programminglanguagelaboratory/rc/pkg/token"

type Expr interface{}

type Id string

type BinOpExpr struct {
	Left  Expr
	Right Expr
	Token token.Token
}

type UnOpExpr struct {
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
