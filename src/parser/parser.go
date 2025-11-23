package parser

import (
	"fmt"
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"interpreter/src/token"
)

// this iota is for the order preference.
// _ means 0 and as we go down the number increases.
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER //< or >
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	inflixParseFn func(ast.Expression) ast.Expression //because we need to know, what is the left side of the inflix operation
)

type Parser struct {
	l *lexer.Lexer //this is to repeatedly call the next token
	//in order to implement debugging we need to log the error
	errors    []string
	curToken  token.Token
	peekToken token.Token

	//Our parser need to know the correct inflex and preflix, for that we need
	preflixParseFns map[token.TokenType]prefixParseFn
	inflixParseFns  map[token.TokenType]inflixParseFn
	//we need helper functions to add the entries into the map
}

func (p *Parser) registerPreflix(tokenType token.TokenType, fn prefixParseFn) {
	p.preflixParseFns[tokenType] = fn
}

func (p *Parser) registerInflix(tokenType token.TokenType, fn inflixParseFn) {
	p.inflixParseFns[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l,
		errors: []string{}}

	p.nextToken()
	p.nextToken() //we are reading 2 tokens, we setting the current and next token

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) { //this will be used when the peektoken does not have the expected type
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.curToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// now we are parsing Expression Statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixExpression[p.curToken.Type]

	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	//we are skipping all the expression till EOF
	if !p.expectPeek(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken} //this is a pointer composite literal in go. pointer-to-struct here token will be initialized with p.curToken

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	//we need to skip expression until we read the end of the statement

	if !p.expectPeek(token.EOF) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool { //this is one of the "assertion functions: nearly all parsers share.
	//This is to ensure the correctness of the next token
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
