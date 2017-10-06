package main

import (
	"Aurora/Parser"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type FunctionParam struct {
	Name string
	Type string
}

type Function struct {
	Name   string
	Params []FunctionParam
}

type TreeShapeListener struct {
	*parser.BaseAuroraListener

	Functions []Function
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *TreeShapeListener) EnterProgram(ctx *parser.ProgramContext) {
	s.Functions = make([]Function, 0)
}

func (s *TreeShapeListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	functionName := ctx.Name.GetText()

	functionDecl := Function{
		Name:   functionName,
		Params: make([]FunctionParam, 0),
	}

	s.Functions = append(s.Functions, functionDecl)
}

func (s *TreeShapeListener) EnterFunctionParam(ctx *parser.FunctionParamContext) {
	lastFunction := len(s.Functions) - 1

	s.Functions[lastFunction].Params = append(s.Functions[lastFunction].Params, FunctionParam{
		Name: ctx.ParamName.GetText(),
		Type: ctx.ParamType.GetText(),
	})
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAuroraLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAuroraParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Printf("%+v", listener.Functions)
}
