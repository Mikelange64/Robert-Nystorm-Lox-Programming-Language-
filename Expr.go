package main

type Expr interface {
    exprNode()
}

type Binary struct {
    Left     Expr
    Operator Token
    Right    Expr
}

type Grouping struct {
    Expression Expr
}

type Literal struct {
    Value any
}

type Unary struct {
    Operator Token
    Right    Expr
}

func (b Binary) exprNode()   {}
func (g Grouping) exprNode() {}
func (l Literal) exprNode()  {}
func (u Unary) exprNode()    {}