// Generated from d:\Itay\Documents\Coding\Go\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser // Aurora

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AuroraListener is a complete listener for a parse tree produced by AuroraParser.
type AuroraListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

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

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

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
}
