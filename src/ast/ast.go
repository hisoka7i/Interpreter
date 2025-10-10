package ast

import (
	"interpreter/src/token"
)

type Node interface {
	TokenLiteral() string //this will be used for debugging and testing
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct { //this is going to be the root node of every AST
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
} //this is going to be the root of the AST

type LetStatement struct{
	Token token.Token
	Name *Identifier
	Value Expression
} //this is going to be the binding statement and we are going to need, token. identifier and the expression

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct{
	Token token.Token
	Value string
}
//statement node and expression node is such that it is only to differentiate between the two nodes

func (ls *Identifier) expressionNode(){}
func (ls *Identifier) TokenLiteral() string {return  ls.Token.Literal} //this is going to be used to hold the name of the identifier