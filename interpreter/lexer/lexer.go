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
func New(input string) *lexer {
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
func (l *lexer) lookup_ident(ident string) token.TokenType {
	if tok, ok := token.Keywords[ident]; ok {
		return tok
	}
	if is_number(ident) {
		return token.INT
	}
	return token.IDENT
}

//TODO: lookup func. for single lexeme

// Token handling
func tokenize(tokenType token.TokenType, ch rune) token.Token {
	token := token.Token{Type: tokenType, Lit: string(ch)}
	return token
}

/*
Sequentially create token from the input
*/
func (l *lexer) get_next_token() token.Token {
	var t token.Token

	// Skip whitespace
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.read_char()
	}

	// If at end of input, return EOF
	if l.ch == 0 {
		return token.Token{Type: token.EOF}
	}

	// handle identifiers, keywords and numbers
	if unicode.IsLetter(l.ch) || unicode.IsNumber(l.ch) {
		ident := l.read_sequence()
		seqType := l.lookup_ident(ident)
		return token.Token{Type: seqType, Lit: ident}
	}

	// Token processing based on current character
	switch l.ch {
	case '=':
		t = tokenize(token.ASS, l.ch)
	case '+':
		t = tokenize(token.PLUS, l.ch)
	case '-':
		t = tokenize(token.SUB, l.ch)
	case '*':
		t = tokenize(token.MUL, l.ch)
	case '/':
		t = tokenize(token.DIV, l.ch)
	case ',':
		t = tokenize(token.COMMA, l.ch)
	case ';':
		t = tokenize(token.SEMICOLON, l.ch)
	case '.':
		t = tokenize(token.DOT, l.ch)
	case '{':
		t = tokenize(token.LBR, l.ch)
	case '}':
		t = tokenize(token.RBR, l.ch)
	case '(':
		t = tokenize(token.LPAR, l.ch)
	case ')':
		t = tokenize(token.RPAR, l.ch)
		// TODO: handle '<='
	case '<':
		t = tokenize(token.LT, l.ch)
		// TODO: handle '>='
	case '>':
		t = tokenize(token.GT, l.ch)
	case '%':
		t = tokenize(token.MOD, l.ch)
	default:
		t = tokenize(token.ILLEGAL, l.ch) // Handle unexpected characters
	}

	// Move to the next character after processing
	l.read_char()
	return t
}
