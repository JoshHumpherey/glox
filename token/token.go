package token

import "fmt"

const (
	_ = iota

	// Single-character tokens
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	Type    int
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s %s %s", getTypeStringRepresentation(t.Type), t.Lexeme, getLiteralStringRepresentation(t.Literal))
}

func getTypeStringRepresentation(tokenType int) string {
	switch tokenType {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case STAR:
		return "*"
	case EOF:
		return "EOF"
	case BANG_EQUAL:
		return "!="
	case BANG:
		return "!"
	case EQUAL_EQUAL:
		return "=="
	case EQUAL:
		return "="
	case LESS_EQUAL:
		return "<="
	case LESS:
		return "<"
	case GREATER_EQUAL:
		return ">="
	case GREATER:
		return ">"
	case SLASH:
		return "/"
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
	default:
		return "N/A"
	}
}

func getLiteralStringRepresentation(literal interface{}) string {
	switch literal := literal.(type) {
	case string:
		return literal
	case []rune:
		return string(literal)
	default:
		return fmt.Sprintf("%v", literal)
	}
}
