package parser

import (
	"reflect"
	"testing"

	"github.com/programminglanguagelaboratory/rc/pkg/ast"
	"github.com/programminglanguagelaboratory/rc/pkg/token"
)

func TestExpr(t *testing.T) {
	for _, tt := range []struct {
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
				X:     ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Y:     ast.IdentExpr{Token: token.Token{Text: "a", Kind: token.ID}, Value: "a"},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"10 + 20 * 30",
			ast.BinaryExpr{
				X: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Y: ast.BinaryExpr{
					X:     ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Y:     ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
					Token: token.Token{Text: "*", Kind: token.ASTERISK},
				},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"(10 + 20) * 30",
			ast.BinaryExpr{
				X: ast.BinaryExpr{
					X:     ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					Y:     ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "+", Kind: token.PLUS},
				},
				Y:     ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
				Token: token.Token{Text: "*", Kind: token.ASTERISK},
			},
		},
		{
			"10 - 20 - 30",
			ast.BinaryExpr{
				X: ast.BinaryExpr{
					X:     ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					Y:     ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Y:     ast.NumberExpr{Token: token.Token{Text: "30", Kind: token.NUMBER}, Value: 30},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"10 - -20",
			ast.BinaryExpr{
				X: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				Y: ast.UnaryExpr{
					X:     ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Token: token.Token{Text: "-", Kind: token.MINUS},
			},
		},
		{
			"!false",
			ast.UnaryExpr{
				X:     ast.BoolExpr{Token: token.Token{Text: "false", Kind: token.BOOL}, Value: false},
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
				X: ast.UnaryExpr{
					X: ast.CallExpr{
						Func: ast.IdentExpr{Token: token.Token{Text: "f", Kind: token.ID}, Value: "f"},
						Arg:  ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
					},
					Token: token.Token{Text: "-", Kind: token.MINUS},
				},
				Y: ast.CallExpr{
					Func: ast.IdentExpr{Token: token.Token{Text: "g", Kind: token.ID}, Value: "g"},
					Arg:  ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
				},
				Token: token.Token{Text: "+", Kind: token.PLUS},
			},
		},
		{
			"f.x",
			ast.FieldExpr{
				X: ast.IdentExpr{Token: token.Token{Text: "f", Kind: token.ID}, Value: "f"},
				Y: "x",
			},
		},
		/* {
			"x . f 10 . g 20",
			ast.CallExpr{
				Func: ast.FieldExpr{
					X: ast.CallExpr{
						Func: ast.FieldExpr{
							X: ast.LitExpr{Token: token.Token{Text: "x", Kind: token.ID}},
							Y: "f",
						},
						Arg: ast.LitExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}},
					},
					Y: "g",
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
		{
			"a => 10",
			ast.FuncExpr{
				Name: ast.Id("a"),
				Body: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
			},
		},
		{
			"const10 := x => 10; x 20",
			ast.DeclExpr{
				Name: ast.Id("const10"),
				Value: ast.FuncExpr{
					Name: ast.Id("x"),
					Body: ast.NumberExpr{Token: token.Token{Text: "10", Kind: token.NUMBER}, Value: 10},
				},
				Body: ast.CallExpr{
					Func: ast.IdentExpr{Token: token.Token{Text: "x", Kind: token.ID}, Value: "x"},
					Arg:  ast.NumberExpr{Token: token.Token{Text: "20", Kind: token.NUMBER}, Value: 20},
				},
			},
		},
		{
			"a => b => a + b",
			ast.FuncExpr{
				Name: ast.Id("a"),
				Body: ast.FuncExpr{
					Name: ast.Id("b"),
					Body: ast.BinaryExpr{
						X:     ast.IdentExpr{Token: token.Token{Text: "a", Kind: token.ID}, Value: "a"},
						Y:     ast.IdentExpr{Token: token.Token{Text: "b", Kind: token.ID}, Value: "b"},
						Token: token.Token{Text: "+", Kind: token.PLUS},
					},
				},
			},
		},
	} {
		actual, err := Parse(tt.code)
		if err != nil {
			t.Errorf("given %v, expected %v, but got an error: %v", tt.code, tt.expected, err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("given %v, expected %v, but got actual %v", tt.code, tt.expected, actual)
		}
	}
}
