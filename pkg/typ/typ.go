package typ

import "fmt"

type Typ interface {
	fmt.Stringer
}

type StringTyp struct{}

func (t *StringTyp) String() string {
	return "str"
}

type NumberTyp struct{}

func (t *NumberTyp) String() string {
	return "num"
}

type BoolTyp struct{}

func (t *BoolTyp) String() string {
	return "bool"
}

type FuncTyp struct{}

func (t *FuncTyp) String() string {
	return "fn"
}
