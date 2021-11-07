package typ

import "fmt"

type Typ interface {
	fmt.Stringer
}

type String struct{}

func (t *String) String() string {
	return "str"
}

type Number struct{}

func (t *Number) String() string {
	return "num"
}

type Bool struct{}

func (t *Bool) String() string {
	return "bool"
}

type Func struct{}

func (t *Func) String() string {
	return "fn"
}
