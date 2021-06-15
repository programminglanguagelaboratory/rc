package token

// TODO: add position info
type Token struct {
	Str  string
	Kind int
}

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

func (token Token) String() string {
	return token.Str
}
