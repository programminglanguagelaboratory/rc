package lexer

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

type lexer struct {
	code  string
	index int
	ch    rune
}

func newLexer(code string) *lexer {
	lexer := lexer{}
	lexer.code = code
	lexer.index = 0
	lexer.ch = []rune(lexer.code)[lexer.index]
	return &lexer
}

func (lexer *lexer) Lex() (token.Token, error) {
	return token.Token{}, errors.New("not implemented")
}
