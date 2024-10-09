package lexer

import (
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

var tok token.TokenType

type Lexer struct {
	input_token      string
	position         int
	current_position int
	char             byte
}

//func tokenize(token.)
