package lexer

import "interpreter/src/token"

type Lexer struct {
	input        string
	position     int  //current postion in input(points to current char)
	readPosition int  //current reading position in input
	ch           byte //current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() //so that the lexer is in reading mode
	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0 //this is ASCII code for NULL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
} //this function will give us the bnext character and advance our position in the input stream

func (l *Lexer) NextToken() token.Token{ //we are getting the current token for the character and then we are shifting the pointer to the next character
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '(':
		tok = newToken(token.RPAREN, l.ch)
	case ')':
		tok = newToken(token.LPAREN, l.ch)
	case '{':
		tok = newToken(token.RBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}