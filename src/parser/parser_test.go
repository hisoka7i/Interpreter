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
		t.Fatalf("Did not get the expected token instead got= %q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not ast Statement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStatement is not '%s'. got= %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got=%s", name, letStmt.TokenLiteral())
		return false
	}
	return true
}
