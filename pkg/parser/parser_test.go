package parser

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/lexer"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestExpr(t *testing.T) {
	for _, testcase := range []struct {
		code     string
		expected ast.Expr
	}{
		{
			"10",
			ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
		},
		{
			"10 + a",
			ast.BinaryExpr{
				Left:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Right: ast.IdentExpr{Token: token.Token{Text: "a", Kind: token.ID}, Value: "a"},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"10 + 20 * 30",
			ast.BinaryExpr{
				Left: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Right: ast.BinaryExpr{
					Left:  ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Right: ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
					Token: token.Token{Text: "*", Kind: token.ASTERISK},
				},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"(10 + 20) * 30",
			ast.BinaryExpr{
				Left: ast.BinaryExpr{
					Left:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					Right: ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "+", Kind: token.PLUS},
				},
				Right: ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
				Token: token.Token{Text: "*", Kind: token.ASTERISK},
			},
		},
		{
			"10 - 20 - 30",
			ast.BinaryExpr{
				Left: ast.BinaryExpr{
					Left:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					Right: ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Right: ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"10 - -20",
			ast.BinaryExpr{
				Left: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Right: ast.UnaryExpr{
					Left:  ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"!false",
			ast.UnaryExpr{
				Left:  ast.BoolExpr{Token: token.Token{Text: "false", Kind: token.BOOL}, Value: false},
				Token: token.Token{Text: "!", Kind: token.EXCLAMATION},
			},
		},
		{
			"f 10 20",
			ast.CallExpr{
				Func: ast.CallExpr{
					Func: ast.IdentExpr{Token: token.Token{Text: "f", Kind: token.ID}, Value: "f"},
					Arg:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				},
				Arg: ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
			},
		},
		{
			"- f 10 + g 20",
			ast.BinaryExpr{
				Left: ast.UnaryExpr{
					Left: ast.CallExpr{
						Func: ast.IdentExpr{Token: token.Token{Text: "f", Kind: token.ID}, Value: "f"},
						Arg:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Right: ast.CallExpr{
					Func: ast.IdentExpr{Token: token.Token{Text: "g", Kind: token.ID}, Value: "g"},
					Arg:  ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
				},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"f.x",
			ast.FieldExpr{
				Left:  ast.IdentExpr{Token: token.Token{Text: "f", Kind: token.ID}, Value: "f"},
				Right: "x",
			},
		},
		/* {
			"x . f 10 . g 20",
			ast.CallExpr{
				Func: ast.FieldExpr{
					Left: ast.CallExpr{
						Func: ast.FieldExpr{
							Left:  ast.LitExpr{Token: token.Token{Text: "x", Kind: token.ID}},
							Right: "f",
						},
						Arg: ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
					},
					Right: "g",
				},
				Arg: ast.LitExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}},
			},
		}, */
		{
			"a := 10; a",
			ast.DeclExpr{
				Name:  ast.Id("a"),
				Value: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Body:  ast.IdentExpr{Token: token.Token{Text: "a", Kind: token.ID}, Value: "a"},
			},
		},
		{
			"a := 10; b := 10; a",
			ast.DeclExpr{
				Name:  ast.Id("a"),
				Value: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Body: ast.DeclExpr{
					Name:  ast.Id("b"),
					Value: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					Body:  ast.IdentExpr{Token: token.Token{Text: "a", Kind: token.ID}, Value: "a"},
				},
			},
		},
	} {
		actual, err := NewParser(lexer.NewLexer(testcase.code)).parseExpr()
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v",
				testcase.code,
				testcase.expected,
				err)
		}

		if !reflect.DeepEqual(testcase.expected, actual) {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.code,
				testcase.expected,
				actual,
			)
		}
	}
}
