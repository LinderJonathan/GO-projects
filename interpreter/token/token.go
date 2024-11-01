package token

type TokenType string

type Token struct {
	Type TokenType
	Lit  string
}

var Keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"for":    FOR,
	"while":  WHILE,
	"return": RETURN,
	"false":  FALSE,
	"true":   TRUE,
}

const (
	// Misc.
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

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
	MOD  = "%"
	ASS  = "="

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
	RETURN   = "RETURN"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
)

//func create
