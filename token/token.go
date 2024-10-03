package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores y literales
	IDENT = "IDENT" // add, x, y, etc.
	INT   = "INT"   // 1, 2, 34253, etc.

	// Operadores
	ASSIGN = "="
	PLUS   = "+"

	// Delimitadores
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
