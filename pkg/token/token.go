package token

// TODO: add position info
type Token struct {
	Text string
	Kind Kind
}

type Kind int

const (
	ID = iota
	STRING
	NUMBER

	PLUS
	MINUS
	ASTERISK
	SLASH

	DOT
	LPAREN
	RPAREN

	EOF
)

var kinds = [...]string{
	ID:     "ID",
	STRING: "STRING",
	NUMBER: "NUMBER",

	PLUS:     "+",
	MINUS:    "-",
	ASTERISK: "*",
	SLASH:    "/",

	DOT:    ".",
	LPAREN: "(",
	RPAREN: ")",

	EOF: "EOF",
}

func (k Kind) String() string {
	return kinds[k]
}

func (k Kind) GetPrec() int {
	switch k {
	case PLUS, MINUS:
		return 10

	case ASTERISK, SLASH:
		return 11
	}

	return 0
}

func (t Token) String() string {
	switch t.Kind {
	case ID, STRING, NUMBER:
		return t.Text
	}
	return t.Kind.String()
}
