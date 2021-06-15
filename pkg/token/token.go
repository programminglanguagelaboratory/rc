package token

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