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
	} {
		actual, _ := newParser(lexer.NewLexer(testcase.code)).parseExpr()
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
