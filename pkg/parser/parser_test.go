package parser

import (
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
			ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
		},
		{
			"10 + a",
			ast.BinaryExpr{
				Left:  ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
				Right: ast.LitExpr{Token: token.Token{Text: "a", Kind: token.ID}},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"10 + 20 * 30",
			ast.BinaryExpr{
				Left: ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
				Right: ast.BinaryExpr{
					Left:  ast.LitExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}},
					Right: ast.LitExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}},
					Token: token.Token{Text: "*", Kind: token.ASTERISK},
				},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"(10 + 20) * 30",
			ast.BinaryExpr{
				Left: ast.BinaryExpr{
					Left:  ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
					Right: ast.LitExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}},
					Token: token.Token{Text: "+", Kind: token.PLUS},
				},
				Right: ast.LitExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}},
				Token: token.Token{Text: "*", Kind: token.ASTERISK},
			},
		},
		{
			"10 - 20 - 30",
			ast.BinaryExpr{
				Left: ast.BinaryExpr{
					Left:  ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
					Right: ast.LitExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Right: ast.LitExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"10 - -20",
			ast.BinaryExpr{
				Left: ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
				Right: ast.UnaryExpr{
					Left:  ast.LitExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"!false",
			ast.UnaryExpr{
				Left:  ast.LitExpr{Token: token.Token{Text: "false", Kind: token.BOOL}},
				Token: token.Token{Text: "!", Kind: token.EXCLAMATION},
			},
		},
		{
			"f(10)",
			ast.CallExpr{
				Func: ast.LitExpr{Token: token.Token{Text: "f", Kind: token.ID}},
				Args: []ast.Expr{
					ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
				},
			},
		},
		{
			"f.x",
			ast.FieldExpr{
				Left:  ast.LitExpr{Token: token.Token{Text: "f", Kind: token.ID}},
				Right: "x",
			},
		},
		{
			"x.f().g()",
			ast.CallExpr{
				Func: ast.FieldExpr{
					Left: ast.CallExpr{
						Func: ast.FieldExpr{
							Left:  ast.LitExpr{Token: token.Token{Text: "x", Kind: token.ID}},
							Right: "f",
						},
					},
					Right: "g",
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
		} else if testcase.expected != actual {
			t.Errorf(
				"given %v, expected %v, but got actual %v",
				testcase.code,
				testcase.expected,
				actual,
			)
		}
	}
}
