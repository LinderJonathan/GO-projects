package ast

import (
	"github.com/LinderJonathan/go-projects/interpreter/lexer"
	"github.com/LinderJonathan/go-projects/interpreter/token"
)

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

/*
program :)
*/


/*
	A program is a slice of statements
*/
type Program struct {
	Stms []Stm
}

func (p *Program) TokenLiteral() string {
	if len(p.Stms) > 0 {
		return p.Stms[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() {

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

func (bs *BlockStm) StmNode() {}
func (bs *BlockStm) TokenLiteral() string {
	if len(bs.Stms) > 0 {
		return bs.Stms[0].TokenLiteral()
	}
	return bs.Stms.Lit
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

func (ws *WhileStm) StmNode() {}
func (ws *WhileStm) TokenLiteral() string {
	return ws.TokenWhile.Lit
}

// an "if" statement is followed by an expression, and possibly an "if else" statement and/or an "else" statement
type ifStm struct {
	TokenIf   token.Token
	Condition Expr
	Block     *BlockStm
	ElseIf    *ifStm
	ElseBlock *BlockStm
}

func (is *WhileStm) StmNode() {}
func (is *WhileStm) TokenLiteral() string {
	return is.TokenIf.Lit
}

// Return statement
type returnStm struct {
	TokenReturn token.Token
	Value Expr

}

func (rs *returnStm) StmNode() {}
func (rs *returnStm) TokenLiteral() string {
	return rs.TokenReturn.Lit
}