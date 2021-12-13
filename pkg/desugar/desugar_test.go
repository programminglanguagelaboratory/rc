package desugar

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestBinaryExpr(t *testing.T) {
	for _, tt := range []struct {
		code     string
		expected ast.Expr
	}{
		{
			code: "a && b",
			expected: ast.CallExpr{
				Func: ast.CallExpr{
					Func: ast.IdentExpr{
						Token: token.Token{Text: "&&", Kind: token.AMPERSANDAMPERSAND},
						Value: "&&",
					},
					Arg: ast.IdentExpr{
						Token: token.Token{Text: "a", Kind: token.ID},
						Value: "a",
					},
				},
				Arg: ast.IdentExpr{
					Token: token.Token{Text: "b", Kind: token.ID},
					Value: "b",
				},
			},
		},
		{
			code: "10 + 20 * 30",
			expected: ast.CallExpr{
				Func: ast.CallExpr{
					Func: ast.IdentExpr{
						Token: token.Token{Text: "+", Kind: token.PLUS},
						Value: "+",
					},
					Arg: ast.NumberExpr{
						Token: token.Token{Text: "10", Kind: token.NUMBER},
						Value: 10,
					},
				},
				Arg: ast.CallExpr{
					Func: ast.CallExpr{
						Func: ast.IdentExpr{
							Token: token.Token{Text: "*", Kind: token.ASTERISK},
							Value: "*",
						},
						Arg: ast.NumberExpr{
							Token: token.Token{Text: "20", Kind: token.NUMBER},
							Value: 20,
						},
					},
					Arg: ast.NumberExpr{
						Token: token.Token{Text: "30", Kind: token.NUMBER},
						Value: 30,
					},
				},
			},
		},
	} {
		expr, err := parser.Parse(tt.code)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v", tt.code, tt.expected, err)
			continue
		}
		actual := Desugar(expr)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, expected %v, but got actual %v", tt.code, tt.expected, actual)
		}
	}
}

func TestUnaryExpr(t *testing.T) {
	code := "!a"
	expr, _ := parser.Parse(code)
	expected := ast.CallExpr{
		Func: ast.IdentExpr{
			Token: token.Token{Text: "!", Kind: token.EXCLAMATION},
			Value: "!",
		},
		Arg: ast.IdentExpr{
			Token: token.Token{Text: "a", Kind: token.ID},
			Value: "a",
		},
	}
	actual := Desugar(expr)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("given %v, expected %v, but got actual %v", expr, expected, actual)
	}
}
