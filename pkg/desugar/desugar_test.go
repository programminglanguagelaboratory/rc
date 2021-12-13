package desugar

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/parser"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestBinaryExpr(t *testing.T) {
	code := "a && b"
	expr, _ := parser.Parse(code)
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
		t.Errorf(
			"given %v, expected %v, but got actual %v",
			expr,
			expected,
			actual,
		)
	}
}
