package lexer

import (
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestToken(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected token.Token
	}{} {
		lexer := Lexer{}
		actual := lexer.lex()
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
