package main

import "fmt"

// here we will only note which line the token appears on, but normally we'd want to also
// display the column and length as well

// Shortcut: storing only line number, not offset + length (gives us column number)
// For Okapi: store offset and lexeme length instead, This gives better erro
// line and column on demand when generating error messages.
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func (t Token) String() string {
	// this statifies the Stinger interface of the fmt package
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}

func (t TokenType) String() string {
    switch t {
    case LEFT_PAREN:
        return "LEFT_PAREN"
    case RIGHT_PAREN:
        return "RIGHT_PAREN"
    case LEFT_BRACE:
        return "LEFT_BRACE"
    case RIGHT_BRACE:
        return "RIGHT_BRACE"
    case COMMA:
        return "COMMA"
    case DOT:
        return "DOT"
    case MINUS:
        return "MINUS"
    case PLUS:
        return "PLUS"
    case SEMICOLON:
        return "SEMICOLON"
    case SLASH:
        return "SLASH"
    case STAR:
        return "STAR"
    case BANG:
        return "BANG"
    case BANG_EQUAL:
        return "BANG_EQUAL"
    case EQUAL:
        return "EQUAL"
    case EQUAL_EQUAL:
        return "EQUAL_EQUAL"
    case GREATER:
        return "GREATER"
    case GREATER_EQUAL:
        return "GREATER_EQUAL"
    case LESS:
        return "LESS"
    case LESS_EQUAL:
        return "LESS_EQUAL"
    case IDENTIFIER:
        return "IDENTIFIER"
    case STRING:
        return "STRING"
    case NUMBER:
        return "NUMBER"
    case AND:
        return "AND"
    case CLASS:
        return "CLASS"
    case ELSE:
        return "ELSE"
    case FALSE:
        return "FALSE"
    case FUN:
        return "FUN"
    case FOR:
        return "FOR"
    case IF:
        return "IF"
    case NIL:
        return "NIL"
    case OR:
        return "OR"
    case PRINT:
        return "PRINT"
    case RETURN:
        return "RETURN"
    case SUPER:
        return "SUPER"
    case THIS:
        return "THIS"
    case TRUE:
        return "TRUE"
    case VAR:
        return "VAR"
    case WHILE:
        return "WHILE"
    case EOF:
        return "EOF"
    default:
        return "UNKNOWN"
    }
}
