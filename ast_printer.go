package main

import "fmt"

func printExpr(expr Expr) string {
    switch e := expr.(type) {
    case Binary:
        return "(" + e.Operator.Lexeme + " " + printExpr(e.Left) + " " + printExpr(e.Right) + ")"
    case Grouping:
        return "(group " + printExpr(e.Expression) + ")"
    case Literal:
        if e.Value == nil {
            return "nil"
        }
        return fmt.Sprintf("%v", e.Value)
    case Unary:
        return "(" + e.Operator.Lexeme + " " + printExpr(e.Right) + ")"
    }
    return ""
}