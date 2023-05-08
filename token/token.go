package token

type TokenType string // Type of token (e.g. IDENT, INT, ASSIGN, PLUS, etc.)

type Token struct {
	Type    TokenType // Type of token (e.g. IDENT, INT, ASSIGN, PLUS, etc.)
	Literal string    // Literal value of token (e.g. "foobar", "123", "+", etc.)
}

// Define constants for the different token types
const (
	ILLEGAL = "ILLEGAL" // Token/character we don't know about
	EOF     = "EOF"     // End of file

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234567890

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION" // fn
	LET      = "LET"      // let x = 5;
)
