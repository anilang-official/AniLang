package lexer

import (
	"anilang/interpreter/token"
)

type Lexer struct {
	input        string // Input string to lex
	position     int    // Current position in input (points to current char)
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // Current char under examination
}

// New returns a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Read the first char in the input
	return l
}

// NextToken returns the next token in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // If we've reached the end of the input...
		l.ch = 0 // Set the current char to 0 (ASCII for "NUL")
	} else {
		l.ch = l.input[l.readPosition] // Otherwise, set the current char to the next char in the input
	}

	l.position = l.readPosition // Set the current position to the current reading position
	l.readPosition += 1         // Increment the reading position
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// newToken returns a new token.Token instance
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
