package main

import "fmt"

type RuntimeError struct {
    Token   Token
    Message string
}

func (e *RuntimeError) Error() string {
    return fmt.Sprintf("[line %d] Runtime error at '%s': %s", e.Token.Line, e.Token.Lexeme, e.Message)
}