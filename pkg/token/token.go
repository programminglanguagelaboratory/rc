package token

// TODO: add position info
type Token struct {
	Str  string
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

func (token Token) String() string {
	return token.Str
}
