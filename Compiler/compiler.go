package main

import (
	"Aurora/Parser"
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Expr interface {
	expr()
}

type IntImmidiateExpr struct {
	Value int
}

type MemberAccessExpr struct {
	AccessFrom Expr
	MemberName string
}

type FunctionCallExpr struct {
	Function Expr
}

func (f IntImmidiateExpr) expr() {}
func (f MemberAccessExpr) expr() {}
func (f FunctionCallExpr) expr() {}

type CompilerContext struct {
	*parser.BaseAuroraListener

	Expressions []Expr
	Current     Expr
}

// once we exited the expression we can continue onward
func (s *CompilerContext) ExitExpr0(ctx *parser.Expr0Context) {
	if s.Current != nil {
		s.Expressions = append(s.Expressions, s.Current)
		s.Current = nil
	}
}

/* Precedence 0 */

func (s *CompilerContext) EnterIntegerImmidiate(ctx *parser.IntegerImmidiateContext) {
	num, _ := strconv.Atoi(ctx.NumberVal.GetText())
	s.Current = IntImmidiateExpr{
		Value: num,
	}
}

/* Precedence 1 */

func (s *CompilerContext) EnterMemberAccess(ctx *parser.MemberAccessContext) {
	s.Current = MemberAccessExpr{
		AccessFrom: s.Current,
		MemberName: ctx.Member.GetText(),
	}
}

func (s *CompilerContext) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	s.Current = FunctionCallExpr{
		Function: s.Current,
	}
}

func NewCompilerContext() *CompilerContext {
	return &CompilerContext{
		Expressions: make([]Expr, 0),
		Current:     nil,
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
