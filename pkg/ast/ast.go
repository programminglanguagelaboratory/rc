package ast

type Expr interface{}

type Id string

type BinOpExpr struct {
	Left  Expr
	Right Expr
	Op    string
}

type UnOpExpr struct {
	Left Expr
	Op   string
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
	Lit string
}
