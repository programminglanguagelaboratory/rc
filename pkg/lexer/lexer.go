package lexer

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

const EOF = -1

type Lexer struct {
	code  string
	index int
	ch    rune
}

func NewLexer(code string) *Lexer {
	l := Lexer{}
	l.code = code
	l.index = 0

	if l.index >= len(l.code) {
		l.ch = EOF
	} else {
		l.ch = []rune(l.code)[l.index]
	}

	return &l
}

func (l *Lexer) next() {
	l.index++

	if l.index >= len(l.code) {
		l.ch = EOF
		return
	}

	l.ch = []rune(l.code)[l.index]
}

func (l *Lexer) lexIdOrBool() (token.Token, error) {
	start := l.index
	for l.ch != EOF && unicode.IsLetter(l.ch) {
		l.next()
	}
	str := l.code[start:l.index]

	switch str {
	case "false":
		return token.Token{Text: "false", Kind: token.BOOL}, nil
	case "true":
		return token.Token{Text: "true", Kind: token.BOOL}, nil
	}

	return token.Token{Text: l.code[start:l.index], Kind: token.ID}, nil
}

func (l *Lexer) lexString() (token.Token, error) {
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
	return token.Token{Text: l.code[start:l.index], Kind: token.STRING}, nil
}

func (l *Lexer) lexNumber() (token.Token, error) {
	start := l.index

	for l.ch != EOF && unicode.IsDigit(l.ch) {
		l.next()
	}

	return token.Token{Text: l.code[start:l.index], Kind: token.NUMBER}, nil
}

func (l *Lexer) skipSpaces() {
	for unicode.IsSpace(l.ch) {
		l.next()
	}
}

func (l *Lexer) Lex() (token.Token, error) {
	l.skipSpaces()
	switch {

	case unicode.IsDigit(l.ch):
		return l.lexNumber()

	case l.ch == '"':
		return l.lexString()

	case unicode.IsLetter(l.ch):
		return l.lexIdOrBool()

	case l.ch == '+':
		l.next()
		return token.Token{Text: "+", Kind: token.PLUS}, nil

	case l.ch == '-':
		l.next()
		return token.Token{Text: "-", Kind: token.MINUS}, nil

	case l.ch == '*':
		l.next()
		return token.Token{Text: "*", Kind: token.ASTERISK}, nil

	case l.ch == '/':
		l.next()
		return token.Token{Text: "/", Kind: token.SLASH}, nil

	case l.ch == '>':
		l.next()

		if l.ch == '=' {
			l.next()
			return token.Token{Text: ">=", Kind: token.GREATERTHANEQUALS}, nil
		}

		return token.Token{Text: ">", Kind: token.GREATERTHAN}, nil

	case l.ch == '<':
		l.next()

		if l.ch == '=' {
			l.next()
			return token.Token{Text: "<=", Kind: token.LESSTHANEQUALS}, nil
		}

		return token.Token{Text: "<", Kind: token.LESSTHAN}, nil

	case l.ch == '=':
		l.next()

		if l.ch == '=' {
			l.next()
			return token.Token{Text: "==", Kind: token.EQUALSEQUALS}, nil
		}

		return token.Token{}, fmt.Errorf("unexpected rune: %q", l.ch)

	case l.ch == '!':
		l.next()

		if l.ch == '=' {
			l.next()
			return token.Token{Text: "!=", Kind: token.EXCLAMATIONEQUALS}, nil
		}

		return token.Token{Text: "!", Kind: token.EXCLAMATION}, nil

	case l.ch == '&':
		l.next()

		if l.ch == '&' {
			l.next()
			return token.Token{Text: "&&", Kind: token.AMPERSANDAMPERSAND}, nil
		}

		return token.Token{}, fmt.Errorf("unexpected rune: %q", l.ch)

	case l.ch == '|':
		l.next()

		if l.ch == '|' {
			l.next()
			return token.Token{Text: "||", Kind: token.BARBAR}, nil
		}

		return token.Token{}, fmt.Errorf("unexpected rune: %q", l.ch)

	case l.ch == '.':
		l.next()
		return token.Token{Text: ".", Kind: token.DOT}, nil

	case l.ch == ',':
		l.next()
		return token.Token{Text: ",", Kind: token.COMMA}, nil

	case l.ch == '(':
		l.next()
		return token.Token{Text: "(", Kind: token.LPAREN}, nil

	case l.ch == ')':
		l.next()
		return token.Token{Text: ")", Kind: token.RPAREN}, nil

	case l.ch == EOF:
		return token.Token{Kind: token.EOF}, nil
	}

	return token.Token{}, fmt.Errorf("unexpected rune: %q", l.ch)
}
