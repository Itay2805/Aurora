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

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseAuroraListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseAuroraListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpr0 is called when production expr0 is entered.
func (s *BaseAuroraListener) EnterExpr0(ctx *Expr0Context) {}

// ExitExpr0 is called when production expr0 is exited.
func (s *BaseAuroraListener) ExitExpr0(ctx *Expr0Context) {}

// EnterStringImmidiate is called when production stringImmidiate is entered.
func (s *BaseAuroraListener) EnterStringImmidiate(ctx *StringImmidiateContext) {}

// ExitStringImmidiate is called when production stringImmidiate is exited.
func (s *BaseAuroraListener) ExitStringImmidiate(ctx *StringImmidiateContext) {}

// EnterVariableImmidiate is called when production variableImmidiate is entered.
func (s *BaseAuroraListener) EnterVariableImmidiate(ctx *VariableImmidiateContext) {}

// ExitVariableImmidiate is called when production variableImmidiate is exited.
func (s *BaseAuroraListener) ExitVariableImmidiate(ctx *VariableImmidiateContext) {}

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
