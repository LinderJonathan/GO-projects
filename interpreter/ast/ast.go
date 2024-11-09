package ast

import (
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

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

/*
	Constructs for statements and expressions
*/

// A block/scope. Code executes exclusively in here
type BlockStm struct {
	lBrace token.Token
	stms   []Stm
	rBrace token.Token
}

// An identifier is some token, followed with a value that it holds
type Identifier struct {
	token token.Token
	value string
}

// a 'let' statement includes the let keyword, an identifer (variable) followed by an expression
type LetStm struct {
	tokenLet token.Token
	ident    *Identifier
	expr     Expr
}

// a 'while' statement includes the while token, followed by a condition. If it evaluates true, the block should be executed
type WhileStm struct {
	tokenWhile token.Token
	condition  Expr
	block      *BlockStm
}
