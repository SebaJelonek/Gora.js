package parser

import (
	"log"

	lexer "github.com/SebaJelonek/Gora.js/Lexer"
)

type Expression interface {
	expressionNode()
	TokenLiteral() string
}

type VarStatement struct {
	Token lexer.Token
}

func Parser() {
	tokens := []lexer.Token{
		{Type: lexer.VAR, Literal: "var"},
		{Type: lexer.IDENT, Literal: "x"},
		{Type: lexer.ASSIGN, Literal: "="},
		{Type: lexer.INT, Literal: "7"},
		{Type: lexer.SEMICOLON, Literal: ";"},
	}
	log.Println(tokens)
}
