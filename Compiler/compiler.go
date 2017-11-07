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

// Stmt is the super type for AST statement node
type Stmt interface {
	stmt()
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
	Operator string
	Left     Expr
	Right    Expr
}

type AssignExpr struct {
	Target Expr
	Source Expr
}

// Statements

type ExprStmt struct {
	Expr Expr
}

type VariableStmt struct {
	Name     string
	Type     string
	Value    Expr
	Optional bool
	Ptr      int
	GCPtr    bool
}

type CodeBlockStmt struct {
	Statements []Stmt
}

func (f IntImmidiateExpr) expr()   {}
func (f IdentImmidiateExpr) expr() {}
func (f StringImmidateExpr) expr() {}
func (f MemberAccessExpr) expr()   {}
func (f FunctionCallExpr) expr()   {}
func (f LHSOperatorExpr) expr()    {}
func (f LROperatorExpr) expr()     {}
func (f AssignExpr) expr()         {}

func (f ExprStmt) stmt()      {}
func (f VariableStmt) stmt()  {}
func (f CodeBlockStmt) stmt() {}

// CompilerContext contains the AST
type CompilerContext struct {
	*parser.BaseAuroraListener

	Statements   []Stmt
	CurrentStmt  Stmt
	CurrentExpr  Expr
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

/* Statements */

// ExitStmt is called when production stmt is exited.
func (s *CompilerContext) ExitStmt(ctx *parser.StmtContext) {
	GetLast(s).Statements = append(GetLast(s).Statements, GetLast(s).CurrentStmt)
	GetLast(s).CurrentStmt = nil
	GetLast(s).CurrentExpr = nil
}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *CompilerContext) ExitExpressionStmt(ctx *parser.ExpressionStmtContext) {
	GetLast(s).CurrentStmt = ExprStmt{
		Expr: GetLast(s).CurrentExpr,
	}
}

// ExitVariableStmt is called when production variableStmt is exited.
func (s *CompilerContext) ExitVariableStmt(ctx *parser.VariableStmtContext) {
	ptrCount := 0
	if ctx.GetPtr() != nil {
		ptrCount = len(ctx.GetPtr().GetText())
	}
	xtype := ""
	if ctx.GetVariableType() != nil {
		xtype = ctx.GetVariableType().GetText()
	}
	GetLast(s).CurrentStmt = VariableStmt{
		Name:     ctx.GetName().GetText(),
		Type:     xtype,
		Value:    GetLast(s).CurrentExpr,
		Optional: ctx.GetOptional() != nil,
		Ptr:      ptrCount,
		GCPtr:    ctx.GetGcptr() != nil,
	}
}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *CompilerContext) EnterCodeBlock(ctx *parser.CodeBlockContext) {
	GetLast(s).CurrentStmt = CodeBlockStmt{
		Statements: nil,
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *CompilerContext) ExitCodeBlock(ctx *parser.CodeBlockContext) {
	cb := GetBeforLast(s).CurrentStmt.(CodeBlockStmt)
	cb.Statements = GetLast(s).Statements
	GetBeforLast(s).CurrentStmt = cb
	GetBeforLast(s).CurrentExpr = nil
	GetBeforLast(s).ChildContext = nil
}

/* Precedence 0 */

// EnterIntegerImmidiate is called when production integerImmidiate is entered.
func (s *CompilerContext) EnterIntegerImmidiate(ctx *parser.IntegerImmidiateContext) {
	num, _ := strconv.Atoi(ctx.GetNumberVal().GetText())
	GetLast(s).CurrentExpr = IntImmidiateExpr{
		Value: num,
	}
}

// EnterIdentifierImmidiate is called when production identifierImmidiate is entered.
func (s *CompilerContext) EnterIdentifierImmidiate(ctx *parser.IdentifierImmidiateContext) {
	GetLast(s).CurrentExpr = IdentImmidiateExpr{
		Identifier: ctx.GetName().GetText(),
	}
}

// EnterStringImmidiate is called when production stringImmidiate is entered.
func (s *CompilerContext) EnterStringImmidiate(ctx *parser.StringImmidiateContext) {
	GetLast(s).CurrentExpr = StringImmidateExpr{
		Value: ctx.GetStringVal().GetText()[1 : len(ctx.GetStringVal().GetText())-1],
	}
}

/* Precedence 1 */

// EnterMemberAccess is called when production memberAccess is entered.
func (s *CompilerContext) EnterMemberAccess(ctx *parser.MemberAccessContext) {
	GetLast(s).CurrentExpr = MemberAccessExpr{
		AccessFrom: GetLast(s).CurrentExpr,
		MemberName: ctx.GetMember().GetText(),
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *CompilerContext) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	GetLast(s).CurrentExpr = FunctionCallExpr{
		Function:   GetLast(s).CurrentExpr,
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
	fc := GetBeforLast(s).CurrentExpr.(FunctionCallExpr)
	fc.Parameters = append(fc.Parameters, GetLast(s).CurrentExpr)
	GetBeforLast(s).CurrentExpr = fc
	GetLast(s).CurrentExpr = nil
}

// /* Precedence 2 */

// EnterLhsOperator is called when production lhsOperator is entered.
func (s *CompilerContext) EnterLhsOperator(ctx *parser.LhsOperatorContext) {
	GetLast(s).CurrentExpr = LHSOperatorExpr{
		Operator: ctx.GetOp().GetText(),
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitLhsOperator is called when production lhsOperator is exited.
func (s *CompilerContext) ExitLhsOperator(ctx *parser.LhsOperatorContext) {
	lhso := GetBeforLast(s).CurrentExpr.(LHSOperatorExpr)
	lhso.On = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = lhso
	GetBeforLast(s).ChildContext = nil
}

// /* Precedence 3 */

// EnterMulDivMod is called when production mulDivMod is entered.
func (s *CompilerContext) EnterMulDivMod(ctx *parser.MulDivModContext) {
	GetLast(s).CurrentExpr = LROperatorExpr{
		Left:     GetLast(s).CurrentExpr,
		Operator: ctx.GetOp().GetText(),
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitMulDivMod is called when production mulDivMod is exited.
func (s *CompilerContext) ExitMulDivMod(ctx *parser.MulDivModContext) {
	lr := GetBeforLast(s).CurrentExpr.(LROperatorExpr)
	lr.Right = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = lr
	GetBeforLast(s).ChildContext = nil
}

/* Precedence 4 */

// EnterAddSub is called when production addSub is entered.
func (s *CompilerContext) EnterAddSub(ctx *parser.AddSubContext) {
	GetLast(s).CurrentExpr = LROperatorExpr{
		Left:     GetLast(s).CurrentExpr,
		Operator: ctx.GetOp().GetText(),
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitAddSub is called when production addSub is exited.
func (s *CompilerContext) ExitAddSub(ctx *parser.AddSubContext) {
	lr := GetBeforLast(s).CurrentExpr.(LROperatorExpr)
	lr.Right = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = lr
	GetBeforLast(s).ChildContext = nil
}

/* Precedence 5 */

// EnterLogicalAnd is called when production logicalAnd is entered.
func (s *CompilerContext) EnterLogicalAnd(ctx *parser.LogicalAndContext) {
	GetLast(s).CurrentExpr = LROperatorExpr{
		Left:     GetLast(s).CurrentExpr,
		Operator: "&&",
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitLogicalAnd is called when production logicalAnd is exited.
func (s *CompilerContext) ExitLogicalAnd(ctx *parser.LogicalAndContext) {
	lr := GetBeforLast(s).CurrentExpr.(LROperatorExpr)
	lr.Right = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = lr
	GetBeforLast(s).ChildContext = nil
}

/* Precedence 6 */

// EnterLogicalOr is called when production logicalOr is entered.
func (s *CompilerContext) EnterLogicalOr(ctx *parser.LogicalOrContext) {
	GetLast(s).CurrentExpr = LROperatorExpr{
		Left:     GetLast(s).CurrentExpr,
		Operator: "||",
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitLogicalOr is called when production logicalOr is exited.
func (s *CompilerContext) ExitLogicalOr(ctx *parser.LogicalOrContext) {
	lr := GetBeforLast(s).CurrentExpr.(LROperatorExpr)
	lr.Right = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = lr
	GetBeforLast(s).ChildContext = nil
}

/* Precedence 7 */

// EnterAssign is called when production assign is entered.
func (s *CompilerContext) EnterAssign(ctx *parser.AssignContext) {
	GetLast(s).CurrentExpr = AssignExpr{
		Target: GetLast(s).CurrentExpr,
	}
	GetLast(s).ChildContext = NewCompilerContext()
}

// ExitAssign is called when production assign is exited.
func (s *CompilerContext) ExitAssign(ctx *parser.AssignContext) {
	assign := GetBeforLast(s).CurrentExpr.(AssignExpr)
	assign.Source = GetLast(s).CurrentExpr
	GetBeforLast(s).CurrentExpr = assign
	GetBeforLast(s).ChildContext = nil
}

// NewCompilerContext ...
func NewCompilerContext() *CompilerContext {
	return &CompilerContext{
		Statements:   make([]Stmt, 0),
		CurrentStmt:  nil,
		CurrentExpr:  nil,
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
