package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	//EOF will tell our parser later that it needs to stop.
	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	//Operators
	ASSIGN = "="
	ADD    = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
	//we are using string also has the advantage of being easy to debug, String is slower in comparison to byte and int.
}