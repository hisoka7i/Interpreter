package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let footbar = 83830;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Let statement does not contains 3 tokens. get= %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Fatalf("Token literal is not let instead it is %q", s.TokenLiteral())
		return false
	} //for the let statement the fist thing used be token which is let
	letStmt, ok := s.(*ast.LetStatement) //this is type check for the interface
	if !ok {
		t.Errorf("Let statment value is incorrect got=%q", letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Did not get the expected value '%s' got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.TokenLiteral())
		return false
	}
	//if everything is correct then simply return true
	return true
}
