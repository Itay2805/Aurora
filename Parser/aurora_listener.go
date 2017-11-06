// Generated from d:\Itay\Documents\Coding\Go\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser // Aurora

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AuroraListener is a complete listener for a parse tree produced by AuroraParser.
type AuroraListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterVariableStmt is called when entering the variableStmt production.
	EnterVariableStmt(c *VariableStmtContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterExpressionStmt is called when entering the expressionStmt production.
	EnterExpressionStmt(c *ExpressionStmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr0 is called when entering the expr0 production.
	EnterExpr0(c *Expr0Context)

	// EnterBrackets is called when entering the brackets production.
	EnterBrackets(c *BracketsContext)

	// EnterStringImmidiate is called when entering the stringImmidiate production.
	EnterStringImmidiate(c *StringImmidiateContext)

	// EnterIdentifierImmidiate is called when entering the identifierImmidiate production.
	EnterIdentifierImmidiate(c *IdentifierImmidiateContext)

	// EnterIntegerImmidiate is called when entering the integerImmidiate production.
	EnterIntegerImmidiate(c *IntegerImmidiateContext)

	// EnterExpr1 is called when entering the expr1 production.
	EnterExpr1(c *Expr1Context)

	// EnterMemberAccess is called when entering the memberAccess production.
	EnterMemberAccess(c *MemberAccessContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionCallParam is called when entering the functionCallParam production.
	EnterFunctionCallParam(c *FunctionCallParamContext)

	// EnterExpr2 is called when entering the expr2 production.
	EnterExpr2(c *Expr2Context)

	// EnterLhsOperator is called when entering the lhsOperator production.
	EnterLhsOperator(c *LhsOperatorContext)

	// EnterExpr3 is called when entering the expr3 production.
	EnterExpr3(c *Expr3Context)

	// EnterMulDivMod is called when entering the mulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterExpr4 is called when entering the expr4 production.
	EnterExpr4(c *Expr4Context)

	// EnterAddSub is called when entering the addSub production.
	EnterAddSub(c *AddSubContext)

	// EnterExpr5 is called when entering the expr5 production.
	EnterExpr5(c *Expr5Context)

	// EnterLogicalAnd is called when entering the logicalAnd production.
	EnterLogicalAnd(c *LogicalAndContext)

	// EnterExpr6 is called when entering the expr6 production.
	EnterExpr6(c *Expr6Context)

	// EnterLogicalOr is called when entering the logicalOr production.
	EnterLogicalOr(c *LogicalOrContext)

	// EnterExpr7 is called when entering the expr7 production.
	EnterExpr7(c *Expr7Context)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitVariableStmt is called when exiting the variableStmt production.
	ExitVariableStmt(c *VariableStmtContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitExpressionStmt is called when exiting the expressionStmt production.
	ExitExpressionStmt(c *ExpressionStmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr0 is called when exiting the expr0 production.
	ExitExpr0(c *Expr0Context)

	// ExitBrackets is called when exiting the brackets production.
	ExitBrackets(c *BracketsContext)

	// ExitStringImmidiate is called when exiting the stringImmidiate production.
	ExitStringImmidiate(c *StringImmidiateContext)

	// ExitIdentifierImmidiate is called when exiting the identifierImmidiate production.
	ExitIdentifierImmidiate(c *IdentifierImmidiateContext)

	// ExitIntegerImmidiate is called when exiting the integerImmidiate production.
	ExitIntegerImmidiate(c *IntegerImmidiateContext)

	// ExitExpr1 is called when exiting the expr1 production.
	ExitExpr1(c *Expr1Context)

	// ExitMemberAccess is called when exiting the memberAccess production.
	ExitMemberAccess(c *MemberAccessContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionCallParam is called when exiting the functionCallParam production.
	ExitFunctionCallParam(c *FunctionCallParamContext)

	// ExitExpr2 is called when exiting the expr2 production.
	ExitExpr2(c *Expr2Context)

	// ExitLhsOperator is called when exiting the lhsOperator production.
	ExitLhsOperator(c *LhsOperatorContext)

	// ExitExpr3 is called when exiting the expr3 production.
	ExitExpr3(c *Expr3Context)

	// ExitMulDivMod is called when exiting the mulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitExpr4 is called when exiting the expr4 production.
	ExitExpr4(c *Expr4Context)

	// ExitAddSub is called when exiting the addSub production.
	ExitAddSub(c *AddSubContext)

	// ExitExpr5 is called when exiting the expr5 production.
	ExitExpr5(c *Expr5Context)

	// ExitLogicalAnd is called when exiting the logicalAnd production.
	ExitLogicalAnd(c *LogicalAndContext)

	// ExitExpr6 is called when exiting the expr6 production.
	ExitExpr6(c *Expr6Context)

	// ExitLogicalOr is called when exiting the logicalOr production.
	ExitLogicalOr(c *LogicalOrContext)

	// ExitExpr7 is called when exiting the expr7 production.
	ExitExpr7(c *Expr7Context)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)
}
