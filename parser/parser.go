package parser

import (
	"github.com/anilang-official/AniLang/ast"
	"github.com/anilang-official/AniLang/lexer"
	"github.com/anilang-official/AniLang/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}
	if !p.expectPeek(token.IDENTFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}