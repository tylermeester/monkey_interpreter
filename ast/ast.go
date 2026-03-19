package ast

import "monkey/token"

//The abstract syntax tree is the data structure used by the parser to
//represent the hierarchical syntactic structure of the source code in a
//simplified, concrete-syntax-free form.

// The program node is the root node of the AST
type Program struct {
	Statements []Statement
}

// Node represents any AST element that can report the literal
// text of its associated token
type Node interface {
	//
	TokenLiteral() string
}

// Statements do not produce values
type Statement interface {
	//Node is an embedded interface, behaves like 'inheriting' the Node interface,
	// giving access to tokenLiteral()
	Node
	statementNode()
}

// Expressions produce values
type Expression interface {
	//Node is an embedded interface, behaves like 'inheriting' the Node interface,
	// giving access to tokenLiteral()
	Node
	expressionNode()
}

// Returns the
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LET STATEMENT STUFF
type LetStatement struct {
	Token token.Token //the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// IDENTIFIER
type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
