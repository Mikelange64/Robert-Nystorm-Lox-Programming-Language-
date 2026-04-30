package main

import "strconv"

var keywords = map[string]TokenType{
	"and" :   AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
}

func CreateScanner(source string) Scanner {
	return Scanner{
		source  : source,
		tokens  : []Token{},
		start   : 0,
		current : 0,
		line    : 0,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
    switch(c) {
    case '(' :
       	s.addToken(LEFT_PAREN, nil)
    case ')' :
       	s.addToken(RIGHT_PAREN, nil)
    case '{' :
       	s.addToken(LEFT_BRACE, nil)
    case '}' :
       	s.addToken(RIGHT_BRACE, nil)
    case ',' :
       	s.addToken(COMMA, nil)
    case '.' :
       	s.addToken(DOT, nil)
    case '-' :
       	s.addToken(MINUS, nil)
    case '+' :
       	s.addToken(PLUS, nil)
    case ';' :
       	s.addToken(SEMICOLON, nil)
    case '*' :
       	s.addToken(STAR, nil)
    case '!' :
       	if s.match('=') {
      		s.addToken(BANG_EQUAL, nil)
       	} else {
      		s.addToken(BANG, nil)
       	}
    case '=' :
       	if s.match('=') {
      		s.addToken(EQUAL_EQUAL, nil)
       	} else {
      		s.addToken(EQUAL, nil)
       	}
    case '<' :
    	if s.match('=') {
      		s.addToken(LESS_EQUAL, nil)
       	} else {
      		s.addToken(LESS, nil)
       	}
    case '>' :
    	if s.match('=') {
      		s.addToken(GREATER_EQUAL, nil)
       	} else {
      		s.addToken(GREATER, nil)
       	}
    case '/' :
        if s.match('/') {
            // A comment goes until the end of the line
            for s.peek() != '\n' && !s.isAtEnd() {
               	s.advance()
            }
        } else {
            s.addToken(SLASH, nil)
        }
    // We ignore white space, tabs, new lines, and
    case ' ' :
    case '\r' :
    case '\t' :
    case '\n' :
        s.line++
    case '"' :
    	s.string()
    default  :
        if s.isDigit(c) {
            s.number()
        } else if s.isAlpha(c) { // if it starts with a letter or an underscore
            s.identifier()
        } else {
            showError(s.line, "Unexpected character.")
            // if we run into a character we do not recognize, we report an error
            // we still consume the token so we're not stuck in an infinite loop
            // also we keep scanning, so we can detect as many errors in one go
            // since hadError gets set to true, we never try to execute any of the code
        }
    }
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start: s.current]
	lexType, ok := keywords[text]
	if !ok{
		lexType = IDENTIFIER
	}
	s.addToken(lexType, nil)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	num, _ := strconv.ParseFloat(s.source[s.start: s.current], 64)
	s.addToken(NUMBER, num)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !(s.isAtEnd()) {
		if s.peek() == '\n' { // we accept multi-line comments
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		showError(s.line, "Unterminated string")
	}

	s.advance() // we consume the closing "
	s.addToken(STRING, s.source[s.start+1: s.current-1]) // Trim surrounding quotes
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.peek() == expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '0'
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current + 1 >= len(s.source) {
		return '0'
	}
	return s.source[s.current + 1]
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) isAlpha(c byte) bool {
	min   := c >= 'a' && c <= 'z'
	maj   := c >= 'A' && c <= 'Z'
	under := c == '_'
	return min || maj || under
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	curr := s.source[s.current]
	s.current++
	return curr
}

func (s *Scanner) addToken(t TokenType, literal any) {
	s.tokens = append(s.tokens, Token{
		Type: t,
		Lexeme: s.source[s.start: s.current],
		Literal: literal,
		Line: s.line,
	})
}
