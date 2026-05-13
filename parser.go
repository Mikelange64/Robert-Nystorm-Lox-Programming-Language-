package main

import "fmt"

type Parser struct {
	Current int
	Tokens  []Token
}

func CreateParser(tokens []Token) Parser {
	return Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

type ParseError struct {
	Token   Token
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[line %d] Error at '%s': %s", e.Token.Line, e.Token.Lexeme, e.Message)
}

func (p *Parser) expression() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) equality() (Expr, error) {
	// The main idea is simple, let's use "5 != 2" as en example
	// Starting at index 0, we call down the precedence chain, to match higher precedence Tokens first
	// In this case 5 will be matched in p.primary() and get us a NUMBER(5) token and consume it
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	// Once the first call is returned (NUMBER(5)) and consumed we're now at index 1 TOKEN(!=).
	// Since p.match() is true, we enter the loop
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		// when we match either tokens we consume it and advance to index 2 TOKEN(2),
		// in order to retrieve the operator, we have to get the token we just consumed using p.previous()
		operator := p.previous()

		// We call down the chain for the right token because equality is left-associative
		// The right side can only be a comparison or anything higher precedence
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}

		// Finally we build the Binary node for the AST. The left side is the original expr retirved at the top (TOKEN(5))
		// The operator from p.previous() (TOKEN(!=)) and the right side from comparison (TOKEN(2))
		expr = Binary{Left: expr, Operator: operator, Right: right}
	}
	// Here we return the Binary node Binary{Left: 5, Operator : !=, Right : 2}
	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(PLUS, MINUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(SLASH, STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(MINUS, BANG) {
		operator := p.previous()
		right, err := p.primary()
		if err != nil {
			return nil, err
		}
		return Unary{Operator: operator, Right: right}, nil
	}
	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(TRUE) {
		return Literal{Value: true}, nil
	}
	if p.match(FALSE) {
		return Literal{Value: false}, nil
	}
	if p.match(NIL) {
		return Literal{Value: nil}, nil
	}
	if p.match(STRING, NUMBER) {
		return Literal{Value: p.previous().Literal}, nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return Grouping{Expression: expr}, nil
	}

	return nil, &ParseError{Token: p.peek(), Message: "expect expression"}
}

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(t TokenType, message string) (Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}
	return Token{}, &ParseError{Token: p.peek(), Message: message}
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) advance() Token {
	p.Current++
	return p.previous()
}

func (p *Parser) previous() Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.peek().Type == SEMICOLON {
			return
		}
		switch p.peek().Type {
		case CLASS, FOR, FUN, IF, PRINT, RETURN, VAR, WHILE:
			return
		}
		p.advance()
	}
}

func (p *Parser) Parse() (Expr, error) {
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}
	return expression, nil
}
