package desugar

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestBinaryExpr(t *testing.T) {
	expr := ast.BinaryExpr{
		X: ast.IdentExpr{
			Token: token.Token{Text: "a", Kind: token.ID},
			Value: "a",
		},
		Y: ast.IdentExpr{
			Token: token.Token{Text: "b", Kind: token.ID},
			Value: "b",
		},
		Token: token.Token{Text: "&&", Kind: token.AMPERSANDAMPERSAND},
	}
	expected := ast.CallExpr{
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
	}
	actual := Desugar(expr)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"given %v, expected %v, but got actual %v",
			expr,
			expected,
			actual,
		)
	}
}

func TestUnaryExpr(t *testing.T) {
	expr := ast.UnaryExpr{
		X: ast.IdentExpr{
			Token: token.Token{Text: "a", Kind: token.ID},
			Value: "a",
		},
		Token: token.Token{Text: "+", Kind: token.PLUS},
	}
	expected := ast.CallExpr{
		Func: ast.IdentExpr{
			Token: token.Token{Text: "+", Kind: token.PLUS},
			Value: "+",
		},
		Arg: ast.IdentExpr{
			Token: token.Token{Text: "a", Kind: token.ID},
			Value: "a",
		},
	}
	actual := Desugar(expr)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"given %v, expected %v, but got actual %v",
			expr,
			expected,
			actual,
		)
	}
}
