package desugar

import "github.com/programminglanguagelaboratory/rc/pkg/ast"

func Desugar(e ast.Expr) ast.Expr {
	switch v := interface{}(e).(type) {
	case ast.BinaryExpr:
		return desugarBinaryExpr(v)
	case ast.UnaryExpr:
		return desugarUnaryExpr(v)
	default:
		return v
	}
}

func desugarBinaryExpr(e ast.BinaryExpr) ast.CallExpr {
	return ast.CallExpr{
		Func: ast.CallExpr{
			Func: ast.IdentExpr{
				Token: e.Token,
				Value: e.Token.Text,
			},
			Arg: e.X,
		},
		Arg: e.Y,
	}
}

func desugarUnaryExpr(e ast.UnaryExpr) ast.CallExpr {
	return ast.CallExpr{
		Func: ast.IdentExpr{
			Token: e.Token,
			Value: e.Token.Text,
		},
		Arg: e.X,
	}
}
