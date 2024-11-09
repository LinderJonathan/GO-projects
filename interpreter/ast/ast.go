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

// An identifier is some token, followed with a value that it holds
type Identifier struct {
	token token.Token
	value string
}

// a 'let' statement includes the let keyword, an identifer (variable) followed by an expression
type LetStm struct {
	token token.Token
	ident *Identifier
	expr  Expr
}
