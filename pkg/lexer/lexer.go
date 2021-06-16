package lexer

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

type Lexer struct {
}

func (lexer *Lexer) Lex() (token.Token, error) {
	return token.Token{}, errors.New("not implemented")
}
