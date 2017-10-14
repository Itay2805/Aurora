package compiler

import (
	"Aurora/Parser"
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

// StringImmidateExpr is AST node for immidiate string value
type StringImmidateExpr struct {
	Value string
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

// LHSOperatorExpr is AST node for representing a unary operator
type LHSOperatorExpr struct {
	Operator string
	On       Expr
}

// LROperatorExpr is AST node for representing an operator which has both Left and Right expression
type LROperatorExpr struct {
	Expr
	Operator string
	Left     Expr
	Right    Expr
}

func (f IntImmidiateExpr) expr()   {}
func (f IdentImmidiateExpr) expr() {}
func (f StringImmidateExpr) expr() {}
func (f MemberAccessExpr) expr()   {}
func (f FunctionCallExpr) expr()   {}
func (f LHSOperatorExpr) expr()    {}
func (f LROperatorExpr) expr()     {}

// CompilerContext contains the AST
type CompilerContext struct {
	*parser.BaseAuroraListener

	Expressions  []Expr
	Current      Expr
	ChildContext *CompilerContext
}

func GetLast(ctx *CompilerContext) *CompilerContext {
	for ctx.ChildContext != nil {
		ctx = ctx.ChildContext
	}
	return ctx
}

func GetBeforLast(ctx *CompilerContext) *CompilerContext {
	temp := ctx
	for ctx.ChildContext != nil {
		temp = ctx
		ctx = ctx.ChildContext
	}
	return temp
}

// ExitExpression is called when production expression is exited.
func (s *CompilerContext) ExitExpression(ctx *parser.ExpressionContext) {
	s.Expressions = append(GetLast(s).Expressions, GetLast(s).Current)
	s.Current = nil
}

/* Precedence 0 */

// EnterIntegerImmidiate is called when production integerImmidiate is entered.
func (s *CompilerContext) EnterIntegerImmidiate(ctx *parser.IntegerImmidiateContext) {
	num, _ := strconv.Atoi(ctx.NumberVal.GetText())
	GetLast(s).Current = IntImmidiateExpr{
		Value: num,
	}
}

// EnterIdentifierImmidiate is called when production identifierImmidiate is entered.
func (s *CompilerContext) EnterIdentifierImmidiate(ctx *parser.IdentifierImmidiateContext) {
	GetLast(s).Current = IdentImmidiateExpr{
		Identifier: ctx.Name.GetText(),
	}
}

// EnterStringImmidiate is called when production stringImmidiate is entered.
func (s *CompilerContext) EnterStringImmidiate(ctx *parser.StringImmidiateContext) {
	GetLast(s).Current = StringImmidateExpr{
		Value: ctx.StringVal.GetText()[1 : len(ctx.StringVal.GetText())-1],
	}
}

/* Precedence 1 */

// EnterMemberAccess is called when production memberAccess is entered.
func (s *CompilerContext) EnterMemberAccess(ctx *parser.MemberAccessContext) {
	GetLast(s).Current = MemberAccessExpr{
		AccessFrom: GetLast(s).Current,
		MemberName: ctx.Member.GetText(),
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *CompilerContext) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	GetLast(s).Current = FunctionCallExpr{
		Function:   GetLast(s).Current,
		Parameters: make([]Expr, 0),
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *CompilerContext) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	GetBeforLast(s).ChildContext = nil
}

// ExitFunctionCallParam is called when production functionCallParam is exited.
func (s *CompilerContext) ExitFunctionCallParam(ctx *parser.FunctionCallParamContext) {
	fc := GetBeforLast(s).Current.(FunctionCallExpr)
	fc.Parameters = append(fc.Parameters, GetLast(s).Current)
	GetBeforLast(s).Current = fc
	GetLast(s).Current = nil
}

// /* Precedence 2 */

// // EnterLhsOperator is called when production lhsOperator is entered.
// func (s *CompilerContext) EnterLhsOperator(ctx *parser.LhsOperatorContext) {
// 	// we need to make sure we do not confuse with multipication
// 	GetLast(s).Current = LHSOperatorExpr{
// 		Operator: ctx.Op.GetText(),
// 	}
// 	GetLast(s).ChildContext = NewCompilerContext()
// }

// // ExitLhsOperator is called when production lhsOperator is exited.
// func (s *CompilerContext) ExitLhsOperator(ctx *parser.LhsOperatorContext) {
// 	lhso := GetBeforLast(s).Current.(LHSOperatorExpr)
// 	lhso.On = GetLast(s).Current
// 	GetBeforLast(s).Current = lhso
// 	GetBeforLast(s).ChildContext = nil
// }

// /* Precedence 3 */

// // EnterDivMulOperator is called when production divMulOperator is entered.
// func (s *CompilerContext) EnterDivMulOperator(ctx *parser.DivMulOperatorContext) {
// 	GetLast(s).Current = LROperatorExpr{
// 		Left:     GetLast(s).Current,
// 		Operator: ctx.Op.GetText(),
// 	}
// 	GetLast(s).ChildContext = NewCompilerContext()
// }

// // ExitDivMulOperator is called when production divMulOperator is exited.
// func (s *CompilerContext) ExitDivMulOperator(ctx *parser.DivMulOperatorContext) {
// 	lr := GetBeforLast(s).Current.(LROperatorExpr)
// 	lr.Right = GetLast(s).Current
// 	GetBeforLast(s).Current = lr
// 	GetBeforLast(s).ChildContext = nil
// }

// NewCompilerContext ...
func NewCompilerContext() *CompilerContext {
	return &CompilerContext{
		Expressions:  make([]Expr, 0),
		Current:      nil,
		ChildContext: nil,
	}
}

func BuildAST(source string) *CompilerContext {
	input := antlr.NewInputStream(source)
	lexer := parser.NewAuroraLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAuroraParser(stream)
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	tree := p.Program()
	listener := NewCompilerContext()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener
}
