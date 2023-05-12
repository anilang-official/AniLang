package test

import (
	"testing"

	"github.com/anilang-official/AniLang/ast"
	"github.com/anilang-official/AniLang/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENTFIER, Literal: "mayVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENTFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
