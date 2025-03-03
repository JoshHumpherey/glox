package scanner

import (
	"fmt"
	"glox/lox_error"
	"glox/token"
	"strconv"
)

type Scanner struct {
	Source       string
	errorHandler lox_error.ErrorHandler
	tokens       []token.Token
	start        int
	current      int
	line         int
}

var keywords = map[string]int{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func (s *Scanner) ScanTokens() []token.Token {
	tokens := []token.Token{}
	s.line = 1

	for !s.isAtEnd() {
		s.scanToken()
	}
	return tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.errorHandler.Report(lox_error.Error{
				Message: fmt.Sprintf("Unexpected Character: '%v'", c),
			})
		}
	}
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.Source[s.current])
}

func (s *Scanner) addToken(tokenType int) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType int, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    0,
	})
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	} else if rune(s.Source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0x0
	}
	return rune(s.Source[s.current])
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.Source) {
		return 0x0
	}
	return rune(s.Source[s.current+1])
}

func (s *Scanner) string() {
	for s.peek() != '"' {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errorHandler.Report(lox_error.Error{
			Line:    s.line,
			Message: "Unterminated String",
		})
	}

	// handle the closing " character
	s.advance()

	value := s.Source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(token.STRING, value)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
	}

	floatVal, _ := strconv.ParseFloat(s.Source[s.start:s.current], 64)
	s.addTokenWithLiteral(token.NUMBER, floatVal)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.Source[s.start:s.current]
	tokenType, isReservedWord := keywords[text]
	if isReservedWord {
		s.addToken(tokenType)
	} else {
		s.addTokenWithLiteral(token.IDENTIFIER, text)
	}
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
