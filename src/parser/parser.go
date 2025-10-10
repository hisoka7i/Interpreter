package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"interpreter/src/token"
)

type Parser struct{
	l *lexer.Lexer  //this is to repeatedly call the next token

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{l :l}

	p.nextToken()
	p.nextToken() //we are reading 2 tokens, we setting the current and next token 

	return p
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program{
	return nil
}

