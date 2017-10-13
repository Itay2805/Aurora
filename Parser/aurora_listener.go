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

	// EnterExpr0 is called when entering the expr0 production.
	EnterExpr0(c *Expr0Context)

	// EnterStringImmidiate is called when entering the stringImmidiate production.
	EnterStringImmidiate(c *StringImmidiateContext)

	// EnterVariableImmidiate is called when entering the variableImmidiate production.
	EnterVariableImmidiate(c *VariableImmidiateContext)

	// EnterIntegerImmidiate is called when entering the integerImmidiate production.
	EnterIntegerImmidiate(c *IntegerImmidiateContext)

	// EnterExpr1 is called when entering the expr1 production.
	EnterExpr1(c *Expr1Context)

	// EnterMemberAccess is called when entering the memberAccess production.
	EnterMemberAccess(c *MemberAccessContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpr0 is called when exiting the expr0 production.
	ExitExpr0(c *Expr0Context)

	// ExitStringImmidiate is called when exiting the stringImmidiate production.
	ExitStringImmidiate(c *StringImmidiateContext)

	// ExitVariableImmidiate is called when exiting the variableImmidiate production.
	ExitVariableImmidiate(c *VariableImmidiateContext)

	// ExitIntegerImmidiate is called when exiting the integerImmidiate production.
	ExitIntegerImmidiate(c *IntegerImmidiateContext)

	// ExitExpr1 is called when exiting the expr1 production.
	ExitExpr1(c *Expr1Context)

	// ExitMemberAccess is called when exiting the memberAccess production.
	ExitMemberAccess(c *MemberAccessContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)
}
