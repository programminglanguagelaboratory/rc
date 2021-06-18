package lexer

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestToken(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected token.Token
	}{
		{"\n\r id", token.Token{Text: "id", Kind: token.ID}},
		{"id", token.Token{Text: "id", Kind: token.ID}},
		{"\"string\"", token.Token{Text: "\"string\"", Kind: token.STRING}},
		{"10", token.Token{Text: "10", Kind: token.NUMBER}},

		{"+", token.Token{Text: "+", Kind: token.PLUS}},
		{"-", token.Token{Text: "-", Kind: token.MINUS}},
		{"*", token.Token{Text: "*", Kind: token.ASTERISK}},
		{"/", token.Token{Text: "/", Kind: token.SLASH}},

		{".", token.Token{Text: ".", Kind: token.DOT}},
		{"(", token.Token{Text: "(", Kind: token.LPAREN}},
		{")", token.Token{Text: ")", Kind: token.RPAREN}},
	} {
		actual, _ := NewLexer(testcase.code).Lex()
		if testcase.expected != actual {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.code,
				testcase.expected,
				actual,
			)
		}
	}
}
