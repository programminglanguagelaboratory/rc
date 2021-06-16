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
		{"\n\r id", token.Token{Str: "id", Kind: token.ID}},
		{"id", token.Token{Str: "id", Kind: token.ID}},
		{"\"string\"", token.Token{Str: "\"string\"", Kind: token.STRING}},
		{"10", token.Token{Str: "10", Kind: token.NUMBER}},

		{"+", token.Token{Str: "+", Kind: token.PLUS}},
		{"-", token.Token{Str: "-", Kind: token.MINUS}},
		{"*", token.Token{Str: "*", Kind: token.ASTERISK}},
		{"/", token.Token{Str: "/", Kind: token.SLASH}},

		{".", token.Token{Str: ".", Kind: token.DOT}},
		{"(", token.Token{Str: "(", Kind: token.LPAREN}},
		{")", token.Token{Str: ")", Kind: token.RPAREN}},
	} {
		actual, _ := newLexer(testcase.code).Lex()
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
