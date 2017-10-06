// Generated from c:\Users\Itay Almog\Desktop\GoWorkspace\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

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

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseAuroraListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseAuroraListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionParam is called when production functionParam is entered.
func (s *BaseAuroraListener) EnterFunctionParam(ctx *FunctionParamContext) {}

// ExitFunctionParam is called when production functionParam is exited.
func (s *BaseAuroraListener) ExitFunctionParam(ctx *FunctionParamContext) {}

// EnterFunctionParamList is called when production functionParamList is entered.
func (s *BaseAuroraListener) EnterFunctionParamList(ctx *FunctionParamListContext) {}

// ExitFunctionParamList is called when production functionParamList is exited.
func (s *BaseAuroraListener) ExitFunctionParamList(ctx *FunctionParamListContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseAuroraListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseAuroraListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BaseAuroraListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BaseAuroraListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAuroraListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAuroraListener) ExitExpr(ctx *ExprContext) {}
