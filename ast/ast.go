package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program satisfies the Node interface
type Program struct {
	// slice of AST nodes that implement the Statement interface
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// STRUCTS THAT SATISFY THE STATEMENT INTERFACE

// LetStatement satisfies the Statement interface
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement satisfies the Statement interface
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// ExpressionStatement satisfies the Statement interface
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// STRUCTS THAT SATISFY THE EXPRESSION INTERFACE

// Identifier satisfies the Expression interface
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string {
	return i.Value
}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// IntegerLiteral satisfies the Expression interface
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
