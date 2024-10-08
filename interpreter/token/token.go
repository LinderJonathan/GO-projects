package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Number  int
}

const (
	// Identifiers
	IDENT = "IDENT"
	INT   = "INT"
	// Operators
	PLUS = "+"
	SUB  = "-"
	DIV  = "/"
	MUL  = "*"
	LT   = "<"
	GT   = ">"
	LTEQ = "<="
	GTEQ = ">="
	// punctuations
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."
	LBR       = "{"
	RBR       = "}"
	LPAR      = "("
	RPAR      = ")"
	// keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	FOR      = "FOR"
	WHILE    = "WHILE"
)
