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
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
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
	checkParserErrors(t, p)
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

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 999434;
`
	l := lexer.New(input) // new lexer
	p := New(l)           // new parser

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {

		// assert that the interface stmt holds *ast.ReturnStatement type
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		// test Token field
		testSingleToken(t, returnStmt.Token, token.RETURN, "return")
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input) // new lexer
	p := New(l)           // new parser
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	// test Token field
	testSingleToken(t, ident.Token, token.IDENT, "foobar")
	// test Name field
	if ident.Value != "foobar" {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s",
			"foobar", ident.Value)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input) // new lexer
	p := New(l)           // new parser
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	// test Token field
	testSingleToken(t, literal.Token, token.INT, "5")
	// test Name field
	if literal.Value != 5 {
		t.Errorf("literal.Value not %s. got=%d", "5", literal.Value)
	}
}
