package lexer

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestToken(t *testing.T) {
	for _, tt := range []struct {
		code     string
		expected token.Token
	}{
		{"\n\r id", token.Token{Text: "id", Kind: token.ID}},
		{"id", token.Token{Text: "id", Kind: token.ID}},
		{"\"string\"", token.Token{Text: "\"string\"", Kind: token.STRING}},
		{"10", token.Token{Text: "10", Kind: token.NUMBER}},
		{"false", token.Token{Text: "false", Kind: token.BOOL}},

		{"+", token.Token{Text: "+", Kind: token.PLUS}},
		{"-", token.Token{Text: "-", Kind: token.MINUS}},
		{"*", token.Token{Text: "*", Kind: token.ASTERISK}},
		{"/", token.Token{Text: "/", Kind: token.SLASH}},

		{">", token.Token{Text: ">", Kind: token.GREATERTHAN}},
		{">=", token.Token{Text: ">=", Kind: token.GREATERTHANEQUALS}},
		{"<", token.Token{Text: "<", Kind: token.LESSTHAN}},
		{"<=", token.Token{Text: "<=", Kind: token.LESSTHANEQUALS}},

		{"==", token.Token{Text: "==", Kind: token.EQUALSEQUALS}},
		{"!=", token.Token{Text: "!=", Kind: token.EXCLAMATIONEQUALS}},

		{"!", token.Token{Text: "!", Kind: token.EXCLAMATION}},
		{"&&", token.Token{Text: "&&", Kind: token.AMPERSANDAMPERSAND}},
		{"||", token.Token{Text: "||", Kind: token.BARBAR}},

		{".", token.Token{Text: ".", Kind: token.DOT}},
		{",", token.Token{Text: ",", Kind: token.COMMA}},
		{"(", token.Token{Text: "(", Kind: token.LPAREN}},
		{")", token.Token{Text: ")", Kind: token.RPAREN}},

		{"=>", token.Token{Text: "=>", Kind: token.EQUALSGREATERTHAN}},
		{"fn", token.Token{Text: "fn", Kind: token.FN}},
		{":=", token.Token{Text: ":=", Kind: token.COLONEQUALS}},
		{";", token.Token{Text: ";", Kind: token.SEMICOLON}},

		{"", token.Token{Kind: token.EOF}},
	} {
		actual, err := NewLexer(tt.code).Lex()
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v", tt.code, tt.expected, err)
			continue
		}
		if tt.expected != actual {
			t.Errorf("given %v, expected %v, but got actual %v", tt.code, tt.expected, actual)
		}
	}
}
