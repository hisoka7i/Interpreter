package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"testing"
)

func TestReturnStatement(t *testing.T){
	input := `return 5;
	return 10;
	return 9009;`
	
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 3{
		t.Fatalf("Incorrect return statement syntax found %d variables",len(program.Statements))
	}

	for _, stmt := range program.Statements{
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Stmt not *ast.returnStatement. got=%t", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return"{
			t.Errorf("returnStmt.tokenLiteral() is not return instead got=%q",returnStmt.TokenLiteral())
		}
	}
}

func TestLetStatement(t *testing.T) {
	input := `let x   5;
	let y = 10;
	let foobar = 83830;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	checkParserErrors(t,p)
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

func checkParserErrors(t *testing.T, p *Parser){
	errors := p.Errors()
	if len(errors) == 0{
		return
	}
	t.Errorf("Parser has %d errors", len(errors))

	for _,msg := range errors{
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
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

