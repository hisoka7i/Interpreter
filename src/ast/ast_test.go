package ast

import (
	"interpreter/src/token"
	"testing"
)

func TestString(t *testing.T){
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "MyValue"},
					Value: "MyValue",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "SomeOtherValue"},
					Value: "SomeOtherValue",
				},
			},
		},
	}

	if program.String() != "let MyValue = SomeOtherValue;"{
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}