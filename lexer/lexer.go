package lexer

import (
	"github.com/anilang-official/AniLang/token"
)

type Lexer struct {
	input        string // Input string to lex
	position     int    // Current position in input (points to current char)
	nextPosition int    // Current reading position in input (after current char)
	ch           byte   // Current char under examination
}

// New returns a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Read the first char in the input
	return l
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.EQUAL, Literal: "="},
			},
			newToken(token.ASSIGN, l.ch),
		)
	case '+':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.INCREMENT, Literal: "+"},
				{Type: token.PLUSEQUAL, Literal: "="},
			},
			newToken(token.PLUS, l.ch),
		)
	case '-':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.DECREMENT, Literal: "-"},
				{Type: token.MINUSEQUAL, Literal: "="},
			},
			newToken(token.MINUS, l.ch),
		)
	case '!':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.NOTEQUAL, Literal: "="},
			},
			newToken(token.BANG, l.ch),
		)
	case '*':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.POW, Literal: "*"},
				{Type: token.MULTIPLYEQUAL, Literal: "="},
			},
			newToken(token.ASTERISK, l.ch),
		)
	case '/':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.DIVIDEEQUAL, Literal: "="},
			},
			newToken(token.SLASH, l.ch),
		)
	case '<':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.LESSTHANOREQUAL, Literal: "="},
				{Type: token.LEFTSHIFT, Literal: "<"},
			},
			newToken(token.LESSTHAN, l.ch),
		)
	case '>':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.GREATERTHANOREQUAL, Literal: "="},
				{Type: token.RIGHTSHIFT, Literal: ">"},
			},
			newToken(token.GREATERTHAN, l.ch),
		)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '&':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.AND, Literal: "&"},
				{Type: token.BITWISEANDEQUAL, Literal: "="},
			},
			newToken(token.BITWISEAND, l.ch),
		)
	case '|':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.OR, Literal: "|"},
				{Type: token.BITWISEOREQUAL, Literal: "="},
			},
			newToken(token.BITWISEOR, l.ch),
		)
	case '%':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.MODULOEQUAL, Literal: "="},
			},
			newToken(token.MODULO, l.ch),
		)
	case '^':
		tok = l.extraTokenCheck(
			[]token.Token{
				{Type: token.BITWISEXOREQUAL, Literal: "="},
			},
			newToken(token.BITWISEXOR, l.ch),
		)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifierType(tok.Literal) // Check if the identifier is a keyword
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// accepts array of token.TokenType and returns true if the current token type is in the array
func (l *Lexer) extraTokenCheck(types []token.Token, _default token.Token) token.Token {
	ch := string(l.peekChar())
	for _, t := range types {
		if ch == t.Literal {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: t.Type, Literal: literal}
		}
	}
	return _default
}

// newToken returns a new token.Token instance
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

// isLetter returns true if the byte is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit returns true if the byte is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}
