package lexer

import (
	"errors"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

const EOF = -1

type lexer struct {
	code  string
	index int
	ch    rune
}

func newLexer(code string) *lexer {
	lexer := lexer{}
	lexer.code = code
	lexer.index = 0

	if lexer.index >= len(lexer.code) {
		lexer.ch = EOF
	} else {
		lexer.ch = []rune(lexer.code)[lexer.index]
	}

	return &lexer
}

func (lexer *lexer) next() {
	lexer.index++

	if lexer.index >= len(lexer.code) {
		lexer.ch = EOF
		return
	}

	lexer.ch = []rune(lexer.code)[lexer.index]
}

func (lexer *lexer) lexString() (token.Token, error) {
	start := lexer.index
	lexer.next()

	for {
		if lexer.ch == EOF {
			return token.Token{}, errors.New("unterminated string literal")
		}
		if lexer.ch == '"' {
			break
		}
		lexer.next()
	}

	lexer.next()
	return token.Token{
		Str:  lexer.code[start:lexer.index],
		Kind: token.STRING,
	}, nil
}

func (lexer *lexer) Lex() (token.Token, error) {
	switch {

	case lexer.ch == '"':
		return lexer.lexString()

	case lexer.ch == '+':
		lexer.next()
		return token.Token{Str: "+", Kind: token.PLUS}, nil

	case lexer.ch == '-':
		lexer.next()
		return token.Token{Str: "-", Kind: token.MINUS}, nil

	case lexer.ch == '*':
		lexer.next()
		return token.Token{Str: "*", Kind: token.ASTERISK}, nil

	case lexer.ch == '/':
		lexer.next()
		return token.Token{Str: "/", Kind: token.SLASH}, nil

	case lexer.ch == '.':
		lexer.next()
		return token.Token{Str: ".", Kind: token.DOT}, nil

	case lexer.ch == '(':
		lexer.next()
		return token.Token{Str: "(", Kind: token.LPAREN}, nil

	case lexer.ch == ')':
		lexer.next()
		return token.Token{Str: ")", Kind: token.RPAREN}, nil

	case lexer.ch == EOF:
		return token.Token{Kind: token.EOF}, nil
	}

	return token.Token{}, errors.New("not implemented")
}
