package lexer

import (
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

type lexer struct {
	input      string
	pos        int
	readingPos int
	ch         byte
	tokens     []token.Token
}

// Initialize the lexer
func init_lexer(input string) *lexer {
	l := &lexer{input: input}

	// initializing pointers and first characters
	l.read_char()
	return l
}

// read a character, update reading position and current position
func (l *lexer) read_char() {

	if len(l.input) <= l.readingPos {
		l.ch = l.input[l.readingPos]
	} else {
		l.ch = 0
	}
	l.pos = l.readingPos
	l.readingPos++
}

// Token handling
func tokenize(tokenType token.TokenType, ch byte) token.Token {
	token := token.Token{Type: tokenType, Lit: string(ch)}
	return token
}

func (l *lexer) get_next_token() token.Token {
	var t token.Token
	switch c := l.ch; c {

	// handling of single character lexeme

	// single punctuators
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

	// operators
	case '+':
		t = tokenize(token.PLUS, l.ch)
	case '-':
		t = tokenize(token.SUB, l.ch)
	case '/':
		t = tokenize(token.DIV, l.ch)
	case '*':
		t = tokenize(token.MUL, l.ch)
	case '<':
		t = tokenize(token.LT, l.ch)
	case '>':
		t = tokenize(token.GT, l.ch)

	}

	return t
}
