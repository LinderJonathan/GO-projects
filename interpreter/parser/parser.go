package Parser

import (
	"github.com/LinderJonathan/go-projects/interpreter/token"
	"github.com/LinderJonathan/go-projects/interpreter/lexer"
	"github.com/LinderJonathan/go-projects/interpreter/ast"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// init token positions
	p.currentToken = l.GetNextToken()
	p.peekToken = l.GetNextToken()

	return p
}

/*
Update token position
*/
func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.GetNextToken()
}