package typ

import "fmt"

type Typ interface {
	fmt.Stringer
}

type StringTyp struct{}

func NewString() *StringTyp {
	return &StringTyp{}
}

func (t *StringTyp) String() string {
	return "str"
}

type NumberTyp struct{}

func NewNumber() *NumberTyp {
	return &NumberTyp{}
}

func (t *NumberTyp) String() string {
	return "num"
}

type BoolTyp struct{}

func NewBool() *BoolTyp {
	return &BoolTyp{}
}

func (t *BoolTyp) String() string {
	return "bool"
}

type FuncTyp struct {
	From Typ
	To   Typ
}

func NewFunc(from, to Typ) *FuncTyp {
	return &FuncTyp{From: from, To: to}
}

func (t *FuncTyp) String() string {
	return "fn"
}

type VarTyp struct {
	Ident string
}

func NewVar(ident string) *VarTyp {
	return &VarTyp{Ident: ident}
}

func (t *VarTyp) String() string {
	return "var"
}
