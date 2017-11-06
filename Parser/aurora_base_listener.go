// Generated from d:\Itay\Documents\Coding\Go\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser // Aurora

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAuroraListener is a complete listener for a parse tree produced by AuroraParser.
type BaseAuroraListener struct{}

var _ AuroraListener = &BaseAuroraListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAuroraListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAuroraListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAuroraListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAuroraListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAuroraListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAuroraListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseAuroraListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseAuroraListener) ExitStatementList(ctx *StatementListContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseAuroraListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseAuroraListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterVariableStmt is called when production variableStmt is entered.
func (s *BaseAuroraListener) EnterVariableStmt(ctx *VariableStmtContext) {}

// ExitVariableStmt is called when production variableStmt is exited.
func (s *BaseAuroraListener) ExitVariableStmt(ctx *VariableStmtContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseAuroraListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseAuroraListener) ExitStmt(ctx *StmtContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BaseAuroraListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BaseAuroraListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAuroraListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAuroraListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr0 is called when production expr0 is entered.
func (s *BaseAuroraListener) EnterExpr0(ctx *Expr0Context) {}

// ExitExpr0 is called when production expr0 is exited.
func (s *BaseAuroraListener) ExitExpr0(ctx *Expr0Context) {}

// EnterBrackets is called when production brackets is entered.
func (s *BaseAuroraListener) EnterBrackets(ctx *BracketsContext) {}

// ExitBrackets is called when production brackets is exited.
func (s *BaseAuroraListener) ExitBrackets(ctx *BracketsContext) {}

// EnterStringImmidiate is called when production stringImmidiate is entered.
func (s *BaseAuroraListener) EnterStringImmidiate(ctx *StringImmidiateContext) {}

// ExitStringImmidiate is called when production stringImmidiate is exited.
func (s *BaseAuroraListener) ExitStringImmidiate(ctx *StringImmidiateContext) {}

// EnterIdentifierImmidiate is called when production identifierImmidiate is entered.
func (s *BaseAuroraListener) EnterIdentifierImmidiate(ctx *IdentifierImmidiateContext) {}

// ExitIdentifierImmidiate is called when production identifierImmidiate is exited.
func (s *BaseAuroraListener) ExitIdentifierImmidiate(ctx *IdentifierImmidiateContext) {}

// EnterIntegerImmidiate is called when production integerImmidiate is entered.
func (s *BaseAuroraListener) EnterIntegerImmidiate(ctx *IntegerImmidiateContext) {}

// ExitIntegerImmidiate is called when production integerImmidiate is exited.
func (s *BaseAuroraListener) ExitIntegerImmidiate(ctx *IntegerImmidiateContext) {}

// EnterExpr1 is called when production expr1 is entered.
func (s *BaseAuroraListener) EnterExpr1(ctx *Expr1Context) {}

// ExitExpr1 is called when production expr1 is exited.
func (s *BaseAuroraListener) ExitExpr1(ctx *Expr1Context) {}

// EnterMemberAccess is called when production memberAccess is entered.
func (s *BaseAuroraListener) EnterMemberAccess(ctx *MemberAccessContext) {}

// ExitMemberAccess is called when production memberAccess is exited.
func (s *BaseAuroraListener) ExitMemberAccess(ctx *MemberAccessContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseAuroraListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseAuroraListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionCallParam is called when production functionCallParam is entered.
func (s *BaseAuroraListener) EnterFunctionCallParam(ctx *FunctionCallParamContext) {}

// ExitFunctionCallParam is called when production functionCallParam is exited.
func (s *BaseAuroraListener) ExitFunctionCallParam(ctx *FunctionCallParamContext) {}

// EnterExpr2 is called when production expr2 is entered.
func (s *BaseAuroraListener) EnterExpr2(ctx *Expr2Context) {}

// ExitExpr2 is called when production expr2 is exited.
func (s *BaseAuroraListener) ExitExpr2(ctx *Expr2Context) {}

// EnterLhsOperator is called when production lhsOperator is entered.
func (s *BaseAuroraListener) EnterLhsOperator(ctx *LhsOperatorContext) {}

// ExitLhsOperator is called when production lhsOperator is exited.
func (s *BaseAuroraListener) ExitLhsOperator(ctx *LhsOperatorContext) {}

// EnterExpr3 is called when production expr3 is entered.
func (s *BaseAuroraListener) EnterExpr3(ctx *Expr3Context) {}

// ExitExpr3 is called when production expr3 is exited.
func (s *BaseAuroraListener) ExitExpr3(ctx *Expr3Context) {}

// EnterMulDivMod is called when production mulDivMod is entered.
func (s *BaseAuroraListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production mulDivMod is exited.
func (s *BaseAuroraListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterExpr4 is called when production expr4 is entered.
func (s *BaseAuroraListener) EnterExpr4(ctx *Expr4Context) {}

// ExitExpr4 is called when production expr4 is exited.
func (s *BaseAuroraListener) ExitExpr4(ctx *Expr4Context) {}

// EnterAddSub is called when production addSub is entered.
func (s *BaseAuroraListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production addSub is exited.
func (s *BaseAuroraListener) ExitAddSub(ctx *AddSubContext) {}

// EnterExpr5 is called when production expr5 is entered.
func (s *BaseAuroraListener) EnterExpr5(ctx *Expr5Context) {}

// ExitExpr5 is called when production expr5 is exited.
func (s *BaseAuroraListener) ExitExpr5(ctx *Expr5Context) {}

// EnterLogicalAnd is called when production logicalAnd is entered.
func (s *BaseAuroraListener) EnterLogicalAnd(ctx *LogicalAndContext) {}

// ExitLogicalAnd is called when production logicalAnd is exited.
func (s *BaseAuroraListener) ExitLogicalAnd(ctx *LogicalAndContext) {}

// EnterExpr6 is called when production expr6 is entered.
func (s *BaseAuroraListener) EnterExpr6(ctx *Expr6Context) {}

// ExitExpr6 is called when production expr6 is exited.
func (s *BaseAuroraListener) ExitExpr6(ctx *Expr6Context) {}

// EnterLogicalOr is called when production logicalOr is entered.
func (s *BaseAuroraListener) EnterLogicalOr(ctx *LogicalOrContext) {}

// ExitLogicalOr is called when production logicalOr is exited.
func (s *BaseAuroraListener) ExitLogicalOr(ctx *LogicalOrContext) {}

// EnterExpr7 is called when production expr7 is entered.
func (s *BaseAuroraListener) EnterExpr7(ctx *Expr7Context) {}

// ExitExpr7 is called when production expr7 is exited.
func (s *BaseAuroraListener) ExitExpr7(ctx *Expr7Context) {}

// EnterAssign is called when production assign is entered.
func (s *BaseAuroraListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseAuroraListener) ExitAssign(ctx *AssignContext) {}
