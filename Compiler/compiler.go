package main

import (
	"Aurora/Parser"
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Expr is the super type for AST expression node
type Expr interface {
	expr()
}

// IntImmidiateExpr is AST node for immidiate int value
type IntImmidiateExpr struct {
	Value int
}

// IdentImmidiateExpr is AST node for immidiate identifier
type IdentImmidiateExpr struct {
	Identifier string
}

// MemberAccessExpr is AST node for member access
type MemberAccessExpr struct {
	AccessFrom Expr
	MemberName string
}

// FunctionCallExpr is AST node for function call
type FunctionCallExpr struct {
	Function   Expr
	Parameters []Expr
}

// UnaryOperatorExpr is AST node for representing a unary operator
type UnaryOperatorExpr struct {
	Operator string
	On       Expr
}

func (f IntImmidiateExpr) expr()   {}
func (f IdentImmidiateExpr) expr() {}
func (f MemberAccessExpr) expr()   {}
func (f FunctionCallExpr) expr()   {}
func (f UnaryOperatorExpr) expr()  {}

// CompilerContext contains the AST
type CompilerContext struct {
	*parser.BaseAuroraListener

	Expressions []Expr
	Current     []Expr
}

// ExitExpression is called when production expression is exited.
func (s *CompilerContext) ExitExpression(ctx *parser.ExpressionContext) {
	if len(s.Current) == 1 {
		s.Expressions = append(s.Expressions, s.Current[len(s.Current)-1])
		s.Current[0] = nil
	}
}

// ExitExpr0 is called when production expr0 is exited.
func (s *CompilerContext) ExitExpr0(ctx *parser.Expr0Context) {

}

/* Precedence 0 */

// EnterIntegerImmidiate is called when production integerImmidiate is entered.
func (s *CompilerContext) EnterIntegerImmidiate(ctx *parser.IntegerImmidiateContext) {
	num, _ := strconv.Atoi(ctx.NumberVal.GetText())
	s.Current[len(s.Current)-1] = IntImmidiateExpr{
		Value: num,
	}
}

// EnterIdentifierImmidiate is called when production identifierImmidiate is entered.
func (s *CompilerContext) EnterIdentifierImmidiate(ctx *parser.IdentifierImmidiateContext) {
	s.Current[len(s.Current)-1] = IdentImmidiateExpr{
		Identifier: ctx.Name.GetText(),
	}
}

/* Precedence 1 */

// EnterMemberAccess is called when production memberAccess is entered.
func (s *CompilerContext) EnterMemberAccess(ctx *parser.MemberAccessContext) {
	s.Current[len(s.Current)-1] = MemberAccessExpr{
		AccessFrom: s.Current[len(s.Current)-1],
		MemberName: ctx.Member.GetText(),
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *CompilerContext) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	s.Current[len(s.Current)-1] = FunctionCallExpr{
		Function: s.Current[len(s.Current)-1],
	}
	s.Current = append(s.Current, nil)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *CompilerContext) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	s.Current = s.Current[:len(s.Current)-1]
}

// ExitFunctionCallParam is called when production functionCallParam is exited.
func (s *CompilerContext) ExitFunctionCallParam(ctx *parser.FunctionCallParamContext) {
	if x, ok := s.Current[len(s.Current)-2].(FunctionCallExpr); ok {
		param := s.Current[len(s.Current)-1]
		x.Parameters = append(x.Parameters, param)
		s.Current[len(s.Current)-2] = x
	}
}

// EnterExpr2 is called when production expr2 is entered.
func (s *CompilerContext) EnterExpr2(ctx *parser.Expr2Context) {
	s.Current[len(s.Current)-1] = UnaryOperatorExpr{
		Operator: ctx.Op.GetText(),
	}
	s.Current = append(s.Current, nil)
}

// ExitExpr2 is called when production expr2 is exited.
func (s *CompilerContext) ExitExpr2(ctx *parser.Expr2Context) {
	if x, ok := s.Current[len(s.Current)-2].(UnaryOperatorExpr); ok {
		result := s.Current[len(s.Current)-1]
		x.On = result
		s.Current[len(s.Current)-2] = x
	}
	s.Current = s.Current[:len(s.Current)-1]
}

// NewCompilerContext ...
func NewCompilerContext() *CompilerContext {
	return &CompilerContext{
		Expressions: make([]Expr, 0),
		Current:     make([]Expr, 1),
	}
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAuroraLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAuroraParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	listener := NewCompilerContext()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Printf("%+v", listener)
}
