package token

// TODO: add position info
type Token struct {
	Text string
	Kind Kind
}

func (t Token) String() string {
	switch t.Kind {
	case ID, STRING, NUMBER, BOOL:
		return t.Text
	}
	return t.Kind.String()
}

type Kind int

const (
	ID = iota
	STRING
	NUMBER
	BOOL
	PLUS
	MINUS
	ASTERISK
	SLASH
	GREATERTHAN
	GREATERTHANEQUALS
	LESSTHAN
	LESSTHANEQUALS
	EQUALSEQUALS
	EXCLAMATIONEQUALS
	EXCLAMATION
	AMPERSANDAMPERSAND
	BARBAR
	DOT
	LPAREN
	RPAREN
	EOF
)

var kinds = [...]string{
	ID:                 "ID",
	STRING:             "STRING",
	NUMBER:             "NUMBER",
	BOOL:               "BOOL",
	PLUS:               "+",
	MINUS:              "-",
	ASTERISK:           "*",
	SLASH:              "/",
	GREATERTHAN:        ">",
	GREATERTHANEQUALS:  ">=",
	LESSTHAN:           "<",
	LESSTHANEQUALS:     "<=",
	EQUALSEQUALS:       "==",
	EXCLAMATIONEQUALS:  "!=",
	EXCLAMATION:        "!",
	AMPERSANDAMPERSAND: "&&",
	BARBAR:             "||",
	DOT:                ".",
	LPAREN:             "(",
	RPAREN:             ")",
	EOF:                "EOF",
}

func (k Kind) String() string {
	return kinds[k]
}

func (k Kind) GetPrec() int {
	switch k {
	case BARBAR:
		return 1
	case AMPERSANDAMPERSAND:
		return 2
	case
		GREATERTHAN,
		GREATERTHANEQUALS,
		LESSTHAN,
		LESSTHANEQUALS,
		EQUALSEQUALS,
		EXCLAMATIONEQUALS:
		return 3
	case PLUS, MINUS:
		return 4
	case ASTERISK, SLASH:
		return 5
	}
	return -1
}
