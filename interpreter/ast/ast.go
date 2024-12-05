package ast

import (
	"github.com/LinderJonathan/go-projects/interpreter/lexer"
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

/*
PARSER
*/
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

/*
	ABSTRACT SYNTAX TREE
*/

/*
Node interfaces
*/
type Node interface {
	TokenLiteral() string
	String() string
}

type Stm interface {
	Node
	StmNode()
}

type Expr interface {
	Node
	ExprNode()
}

type Ast interface {
	// holds nodes
	*Node
}

/*
program :)
*/
func (p *Program) TokenLiteral() string {
	if len(p.Stms) > 0 {
		return p.Stms[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() {

}

/*
Construct for a program. Main entry point
*/
type Program struct {
	Stms []Stm
}

/*
	Constructs for statements and expressions
*/

// A block/scope. Code executes exclusively in here
type BlockStm struct {
	LBrace token.Token
	Stms   []Stm
	RBrace token.Token
}

// An identifier is some token, followed with a value that it holds
type Identifier struct {
	TokId token.Token
	Value string
}

func (id *Identifier) StmNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.TokId.Lit
}

// a 'let' statement includes the let keyword, an identifer (variable) followed by an expression
type LetStm struct {
	TokenLet token.Token
	Ident    *Identifier
	Expr     Expr
}

func (ls *LetStm) StmNode() {}
func (ls *LetStm) TokenLiteral() string {
	return ls.TokenLet.Lit
}

// a 'while' statement includes the while token, followed by a condition. If it evaluates true, the block should be executed
type WhileStm struct {
	TokenWhile token.Token
	Condition  Expr
	Block      *BlockStm
}

// an "if" statement is followed by an expression, and possibly an "if else" statement and/or an "else" statement
type ifStm struct {
	TokenIf   token.Token
	Condition Expr
	Block     *BlockStm
	ElseIf    *ifStm
	ElseBlock *BlockStm
}

//
