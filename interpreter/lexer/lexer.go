package lexer

import (
	"unicode"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

type lexer struct {
	input      string
	pos        int
	readingPos int
	ch         rune
	tokens     []token.Token
}

// Initialize the lexer
func NewLexer(input string) *lexer {
	l := &lexer{input: input}

	// initializing pointers and first characters
	l.read_char()
	return l
}

// read a character, update reading position and current position
func (l *lexer) read_char() {

	if len(l.input) <= l.readingPos {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readingPos])

	}
	l.pos = l.readingPos
	l.readingPos += 1
}

// peeks and reads next character in sinput
func (l *lexer) peek_char() rune {
	if len(l.input) <= l.readingPos {
		return 0 // Return 0 if at the end of input
	}
	return rune(l.input[l.readingPos])
}

// read an identifier or a keyword
func (l *lexer) read_sequence() string {

	// read one char at a time
	key := ""
	for unicode.IsDigit(l.ch) || unicode.IsLetter(l.ch) || l.ch == '_' {
		key += string(l.ch)
		l.read_char()
	}
	return key
}

/*
Check if a lexeme is a number
*/
// TODO: this
func is_number(ident string) bool {
	for _, ch := range ident {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

/*
Checks if a sequence of characters is a keyword
Otherwise an identifier
*/
func (l *lexer) lookup_ident_type(ident string) token.TokenType {
	if tok, ok := token.Keywords[ident]; ok {
		return tok
	}
	if is_number(ident) {
		return token.INT
	}
	return token.IDENT
}

func (l *lexer) lookup_lexeme(ch rune) token.TokenType {
	if tok, ok := token.Singles[string(ch)]; ok {
		return tok
	}
	return token.ILLEGAL
}

// Token handling
func tokenize(tokenType token.TokenType, ch rune) token.Token {
	token := token.Token{Type: tokenType, Lit: string(ch)}
	return token
}

/*
Sequentially create token from the input
*/
func (l *lexer) get_next_token() token.Token {
	// Skip whitespace
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.read_char()
	}

	// If at end of input, return EOF
	if l.ch == 0 {
		return token.Token{Type: token.EOF}
	}

	// Handle identifiers, keywords, and numbers
	if unicode.IsLetter(l.ch) || unicode.IsNumber(l.ch) {
		ident := l.read_sequence()
		seqType := l.lookup_ident_type(ident)
		return token.Token{Type: seqType, Lit: ident}
	}

	// Handle single lexemes and compound operators
	singleType := l.lookup_lexeme(l.ch)
	switch singleType {
	case token.LT:
		if l.peek_char() == '=' {
			l.read_char() // Advance past '=' for compound token
			l.read_char()
			return token.Token{Type: token.LTEQ, Lit: "<="}
		}
		l.read_char()
		return token.Token{Type: token.LT, Lit: "<"}

	case token.GT:
		if l.peek_char() == '=' {
			l.read_char() // Advance past '=' for compound token
			l.read_char()
			return token.Token{Type: token.GTEQ, Lit: ">="}
		}
		l.read_char()
		return token.Token{Type: token.GT, Lit: ">"}
	}

	// Fallback to single lexeme processing
	t := tokenize(singleType, l.ch)

	// Move to the next character after processing
	l.read_char()
	return t
}
