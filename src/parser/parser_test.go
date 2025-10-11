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

}
