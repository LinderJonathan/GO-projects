package lexer

import (
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

type lexer struct {
	input           string
	position        int
	currentPosition int
	ch              byte
	tokens          []token.Token
}

func (l *lexer) lexInit(input string) {
	l.currentPosition = 0
	l.position = 0
	l.input = input
	l.ch = 0
}

func (l *lexer) read_char(input string) {
	l.input = input
	l.position = 0
	l.currentPosition = 0

	if len(input) > 0 {
		l.ch = input[0]
	} else {
		l.ch = 0
	}
}

/*
Function Name: advance
Description: advance the input position pointer

Params:
  - input: string - input code

Returns:
  - nil
*/
func (l *lexer) advance_pos(input string) {
	l.position++
	if l.position >= len(input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
}

// Token handling
func (l *lexer) add_token(tokenType token.TokenType, literal string, number int) {
	token := token.Token{Type: tokenType, Literal: literal, Number: number}
	l.tokens = append(l.tokens, token)
}

// TODO: restructure token
func (l *lexer) tokenize(input string) {
	switch c := l.ch; c {
	// punctuations
	case ',':
		l.add_token(token.COMMA, string(l.ch), nil)
	}
}
