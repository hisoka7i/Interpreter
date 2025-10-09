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
	if l.readPosition >= len(l.input) {
		l.ch = 0 //this is ASCII code for NULL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
} //this function will give us the bnext character and advance our position in the input stream

func (l *Lexer) NextToken() token.Token{ //we are getting the current token for the character and then we are shifting the pointer to the next character
	var tok token.Token
	l.skipWhiteSpace()  //we are skipping all the whitespaces, new lines, tab and /r since they are not relevant in making token
	switch l.ch {
	case '=':
		if(l.peekChar() == '='){
			ch := l.ch
			l.readChar() // we are incrementing the position
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		}else{
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.SUBTRACT, l.ch)
	case '!':
		if(l.peekChar() == '='){
			ch := l.ch
			l.readChar() //we are incrementing the position
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		}else{
		tok = newToken(token.BANG, l.ch)}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	// case 'let':
	// 	tok = newToken(token.LET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch){
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return  tok
		}else if isDigit(l.ch){
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}else{
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string{
	position := l.position
	for isLetter(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool{
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'  //we are gonna identify the keywords using this and identifiers with this one
}

func (l *Lexer) skipWhiteSpace(){
	for (l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'){
		l.readChar() //using this we can simply increase the read position
	}
}

//similarly how we are reading the words and we can trying to know, weather it is a keyword or identifier
//we can do the same exact thing, for the 
func (l *Lexer) readNumber() string{
	position := l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9' 
}
//we are creating a peekchar functionality just in case if we want to peek the character ahead without increamenting the positon
func (l *Lexer) peekChar() byte{
	if (l.readPosition >= len(l.input)){
		return  0
	}else{
		return l.input[l.readPosition]
	}
}