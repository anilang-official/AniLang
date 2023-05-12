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
	IDENTFIER = "IDENT"  // add, foobar, x, y, ...
	INT       = "INT"    // 1234567890
	STRING    = "STRING" // "foobar"

	// Operators
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"
	BITWISEAND = "&"
	BITWISEOR  = "|"

	EQUAL              = "=="
	NOTEQUAL           = "!="
	LESSTHAN           = "<"
	GREATERTHAN        = ">"
	OR                 = "||"
	AND                = "&&"
	GREATERTHANOREQUAL = ">="
	LESSTHANOREQUAL    = "<="

	LEFTSHIFT  = "<<"
	RIGHTSHIFT = ">>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION" // fn
	LET      = "LET"      // let x = 5;
	TRUE     = "TRUE"     // true
	FALSE    = "FALSE"    // false
	IF       = "IF"       // if (x < y) { return true; }
	ELSEIF   = "ELIF"     // elif (x > y) { return false; }
	ELSE     = "ELSE"     // else { return false; }
	RETURN   = "RETURN"   // return 5;
	BREAK    = "BREAK"    // break;
	CONTINUE = "CONTINUE" // continue;
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"elif":     ELSEIF,
	"sayonara": RETURN,
	"yamete":   BREAK,
	"continue": CONTINUE,
}

func LookupIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTFIER
}

func LookupTokenIdentifier(tokenType TokenType) string {
	for identifier, tok := range keywords {
		if tok == tokenType {
			return identifier
		}
	}
	return ""
}
