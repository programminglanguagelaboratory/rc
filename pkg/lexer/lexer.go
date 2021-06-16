package lexer

import (
	"errors"
	"unicode"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

const EOF = -1

type lexer struct {
	code  string
	index int
	ch    rune
}

func newLexer(code string) *lexer {
	l := lexer{}
	l.code = code
	l.index = 0

	if l.index >= len(l.code) {
		l.ch = EOF
	} else {
		l.ch = []rune(l.code)[l.index]
	}

	return &l
}

func (l *lexer) next() {
	l.index++

	if l.index >= len(l.code) {
		l.ch = EOF
		return
	}

	l.ch = []rune(l.code)[l.index]
}

func (l *lexer) lexId() (token.Token, error) {
	return token.Token{}, errors.New("not implemented")
}

func (l *lexer) lexString() (token.Token, error) {
	start := l.index
	l.next()

	for {
		if l.ch == EOF {
			return token.Token{}, errors.New("unterminated string literal")
		}
		if l.ch == '"' {
			break
		}
		l.next()
	}

	l.next()
	return token.Token{Str: l.code[start:l.index], Kind: token.STRING}, nil
}

func (l *lexer) lexNumber() (token.Token, error) {
	return token.Token{}, errors.New("not implemented")
}

func (l *lexer) Lex() (token.Token, error) {
	switch {

	case unicode.IsDigit(l.ch):
		return l.lexNumber()

	case l.ch == '"':
		return l.lexString()

	case unicode.IsLetter(l.ch):
		return l.lexId()

	case l.ch == '+':
		l.next()
		return token.Token{Str: "+", Kind: token.PLUS}, nil

	case l.ch == '-':
		l.next()
		return token.Token{Str: "-", Kind: token.MINUS}, nil

	case l.ch == '*':
		l.next()
		return token.Token{Str: "*", Kind: token.ASTERISK}, nil

	case l.ch == '/':
		l.next()
		return token.Token{Str: "/", Kind: token.SLASH}, nil

	case l.ch == '.':
		l.next()
		return token.Token{Str: ".", Kind: token.DOT}, nil

	case l.ch == '(':
		l.next()
		return token.Token{Str: "(", Kind: token.LPAREN}, nil

	case l.ch == ')':
		l.next()
		return token.Token{Str: ")", Kind: token.RPAREN}, nil

	case l.ch == EOF:
		return token.Token{Kind: token.EOF}, nil
	}

	return token.Token{}, errors.New("not implemented")
}
