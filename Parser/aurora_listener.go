// Generated from c:\Users\Itay Almog\Desktop\GoWorkspace\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser // Aurora

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AuroraListener is a complete listener for a parse tree produced by AuroraParser.
type AuroraListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionParam is called when entering the functionParam production.
	EnterFunctionParam(c *FunctionParamContext)

	// EnterFunctionParamList is called when entering the functionParamList production.
	EnterFunctionParamList(c *FunctionParamListContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionParam is called when exiting the functionParam production.
	ExitFunctionParam(c *FunctionParamContext)

	// ExitFunctionParamList is called when exiting the functionParamList production.
	ExitFunctionParamList(c *FunctionParamListContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
