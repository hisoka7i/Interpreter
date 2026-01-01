package parser

import (
	"fmt"
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"interpreter/src/token"
	"strconv"
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

// adding precendence map or table
var precedence = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.ADD:      SUM,
	token.SUBTRACT: SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

// this function will return the precedence of the peek token(next token)
func (p *Parser) peekPrecedence() int {
	if p, ok := precedence[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedence[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

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

// this function is for Integer literal
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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
	p.preflixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPreflix(token.BANG, p.parsePrefixExpression)
	p.registerPreflix(token.SUBTRACT, p.parsePrefixExpression)
	p.registerPreflix(token.IDENT, p.parseIdentifier)
	p.registerPreflix(token.INT, p.parseIntegerLiteral)
	//we need to register the infix parse functions
	p.inflixParseFns = make(map[token.TokenType]inflixParseFn)
	p.registerInflix(token.ADD, p.parseInfixExpression)
	p.registerInflix(token.SUBTRACT, p.parseInfixExpression)
	p.registerInflix(token.SLASH, p.parseInfixExpression)
	p.registerInflix(token.ASTERISK, p.parseInfixExpression)
	p.registerInflix(token.EQ, p.parseInfixExpression)
	p.registerInflix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInflix(token.LT, p.parseInfixExpression)
	p.registerInflix(token.GT, p.parseInfixExpression)
	return p
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InflixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

/*
	Here we are adding the prefix expression parsing

we are adding expression token and operator
then we are moving to the next token and setting the right hand expression with prefix precedence
*/
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	//since more then one token needs to be parsed we are calling parse expression
	expression.Right = p.parseExpression(PREFIX)
	//current token is the open after prefix operator argument
	return expression
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
// There is no precendence for this function it just uses the LOWEST
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// We are using the precedence to parse the expression correctly
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.preflixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.inflixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
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

// we are adding parse Expression method to give us better error message
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("No prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
