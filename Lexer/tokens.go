package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF       = "EOF"
	VAR       = "VAR"
	ASSIGN    = "="
	PLUS      = "+"
	IDENT     = "IDENT"
	SEMICOLON = ";"
	IF        = "IF"
	INT       = "INT"
)
