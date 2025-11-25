package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

// HELPER FUNCTION
// verify that a given token.Token t is made of
// expected type et and of expected literal el
func testSingleToken(t *testing.T, tok token.Token,
	et token.TokenType, el string) {
	if tok.Type == et {
		if tok.Literal != el {
			t.Fatalf("tokenliteral wrong. expected=%q, got=%q",
				el, tok.Literal)
		}
	} else {
		t.Fatalf("tokentype wrong. expected=%q, got=%q",
			et, tok.Type)
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	// assert that the interface s holds *ast.LetStatement type
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	// test Token field
	testSingleToken(t, letStmt.Token, token.LET, "let")

	// test Name field
	testSingleToken(t, letStmt.Name.Token, token.IDENT, name)

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s",
			name, letStmt.Name.Value)
		return false
	}

	return true
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input) // new lexer
	p := New(l)           // new parser

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
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
