package ast

import (
	"bytes"
	"interpreter/src/token"
)

type Node interface {
	TokenLiteral() string //this will be used for debugging and testing
	String() string       //.This is to print our ast, which in turn will make our life easier
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
} //this is going to be the binding statement and we are going to need, token. identifier and the expression

type ReturnStatement struct { //here we are defining the ast for the return statement
	Token       token.Token
	ReturnValue Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
} //this is the ast for the expressions

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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value }

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

//ast.ExpressionStatement fullfills that as.Statement interface.

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal } //this is function overloading for the return statement

//statement node and expression node is such that it is only to differentiate between the two nodes

func (ls *Identifier) expressionNode()      {}
func (ls *Identifier) TokenLiteral() string { return ls.Token.Literal } //this is going to be used to hold the name of the identifier
