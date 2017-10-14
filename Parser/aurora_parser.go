// Generated from d:\Itay\Documents\Coding\Go\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser // Aurora

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 94, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 3, 2, 3, 2, 5, 2, 31, 10, 2, 3, 3, 3, 3, 7, 3, 35, 10,
	3, 12, 3, 14, 3, 38, 11, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 48, 10, 5, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 66, 10, 10, 3,
	10, 5, 10, 69, 10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 76, 10,
	12, 3, 12, 3, 12, 7, 12, 80, 10, 12, 12, 12, 14, 12, 83, 11, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 92, 10, 14, 3, 14, 2,
	2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 4, 4, 3, 3, 3,
	17, 17, 3, 2, 8, 12, 2, 93, 2, 28, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 39,
	3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 56, 3, 2, 2, 2,
	14, 58, 3, 2, 2, 2, 16, 60, 3, 2, 2, 2, 18, 65, 3, 2, 2, 2, 20, 70, 3,
	2, 2, 2, 22, 73, 3, 2, 2, 2, 24, 86, 3, 2, 2, 2, 26, 88, 3, 2, 2, 2, 28,
	30, 5, 4, 3, 2, 29, 31, 7, 2, 2, 3, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2,
	2, 31, 3, 3, 2, 2, 2, 32, 36, 5, 6, 4, 2, 33, 35, 5, 6, 4, 2, 34, 33, 3,
	2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37,
	5, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 40, 5, 8, 5, 2, 40, 41, 9, 2, 2,
	2, 41, 7, 3, 2, 2, 2, 42, 48, 5, 10, 6, 2, 43, 48, 5, 12, 7, 2, 44, 48,
	5, 16, 9, 2, 45, 48, 5, 14, 8, 2, 46, 48, 5, 18, 10, 2, 47, 42, 3, 2, 2,
	2, 47, 43, 3, 2, 2, 2, 47, 44, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46,
	3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 5, 8, 5, 2, 50, 49, 3, 2, 2, 2,
	50, 51, 3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52, 53, 7, 4, 2, 2, 53, 54, 5, 8,
	5, 2, 54, 55, 7, 5, 2, 2, 55, 11, 3, 2, 2, 2, 56, 57, 7, 14, 2, 2, 57,
	13, 3, 2, 2, 2, 58, 59, 7, 15, 2, 2, 59, 15, 3, 2, 2, 2, 60, 61, 7, 13,
	2, 2, 61, 17, 3, 2, 2, 2, 62, 66, 5, 22, 12, 2, 63, 66, 5, 20, 11, 2, 64,
	66, 5, 26, 14, 2, 65, 62, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2,
	2, 2, 66, 68, 3, 2, 2, 2, 67, 69, 5, 18, 10, 2, 68, 67, 3, 2, 2, 2, 68,
	69, 3, 2, 2, 2, 69, 19, 3, 2, 2, 2, 70, 71, 7, 6, 2, 2, 71, 72, 7, 15,
	2, 2, 72, 21, 3, 2, 2, 2, 73, 75, 7, 4, 2, 2, 74, 76, 5, 24, 13, 2, 75,
	74, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 81, 3, 2, 2, 2, 77, 78, 7, 7, 2,
	2, 78, 80, 5, 24, 13, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	84, 85, 7, 5, 2, 2, 85, 23, 3, 2, 2, 2, 86, 87, 5, 8, 5, 2, 87, 25, 3,
	2, 2, 2, 88, 89, 9, 3, 2, 2, 89, 91, 5, 8, 5, 2, 90, 92, 5, 26, 14, 2,
	91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 27, 3, 2, 2, 2, 11, 30, 36,
	47, 50, 65, 68, 75, 81, 91,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "')'", "'.'", "','", "'-'", "'+'", "'!'", "'*'", "'&'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "IntegerLiteral", "StringLiteral",
	"Identifier", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "expressionList", "expression", "expr0", "brackets", "stringImmidiate",
	"identifierImmidiate", "integerImmidiate", "expr1", "memberAccess", "functionCall",
	"functionCallParam", "expr2",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AuroraParser struct {
	*antlr.BaseParser
}

func NewAuroraParser(input antlr.TokenStream) *AuroraParser {
	this := new(AuroraParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Aurora.g4"

	return this
}

// AuroraParser tokens.
const (
	AuroraParserEOF            = antlr.TokenEOF
	AuroraParserT__0           = 1
	AuroraParserT__1           = 2
	AuroraParserT__2           = 3
	AuroraParserT__3           = 4
	AuroraParserT__4           = 5
	AuroraParserT__5           = 6
	AuroraParserT__6           = 7
	AuroraParserT__7           = 8
	AuroraParserT__8           = 9
	AuroraParserT__9           = 10
	AuroraParserIntegerLiteral = 11
	AuroraParserStringLiteral  = 12
	AuroraParserIdentifier     = 13
	AuroraParserWhitespace     = 14
	AuroraParserNewline        = 15
	AuroraParserBlockComment   = 16
	AuroraParserLineComment    = 17
)

// AuroraParser rules.
const (
	AuroraParserRULE_program             = 0
	AuroraParserRULE_expressionList      = 1
	AuroraParserRULE_expression          = 2
	AuroraParserRULE_expr0               = 3
	AuroraParserRULE_brackets            = 4
	AuroraParserRULE_stringImmidiate     = 5
	AuroraParserRULE_identifierImmidiate = 6
	AuroraParserRULE_integerImmidiate    = 7
	AuroraParserRULE_expr1               = 8
	AuroraParserRULE_memberAccess        = 9
	AuroraParserRULE_functionCall        = 10
	AuroraParserRULE_functionCallParam   = 11
	AuroraParserRULE_expr2               = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(AuroraParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *AuroraParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AuroraParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.ExpressionList()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(27)
			p.Match(AuroraParserEOF)
		}

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *AuroraParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AuroraParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Expression()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__1)|(1<<AuroraParserT__3)|(1<<AuroraParserT__5)|(1<<AuroraParserT__6)|(1<<AuroraParserT__7)|(1<<AuroraParserT__8)|(1<<AuroraParserT__9)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(31)
			p.Expression()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *ExpressionContext) Newline() antlr.TerminalNode {
	return s.GetToken(AuroraParserNewline, 0)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AuroraParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AuroraParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AuroraParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Expr0()
	}
	p.SetState(38)
	_la = p.GetTokenStream().LA(1)

	if !(((_la - -1)&-(0x1f+1)) == 0 && ((1<<uint((_la - -1)))&((1<<(AuroraParserEOF - -1))|(1<<(AuroraParserT__0 - -1))|(1<<(AuroraParserNewline - -1)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IExpr0Context is an interface to support dynamic dispatch.
type IExpr0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr0Context differentiates from other interfaces.
	IsExpr0Context()
}

type Expr0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr0Context() *Expr0Context {
	var p = new(Expr0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr0
	return p
}

func (*Expr0Context) IsExpr0Context() {}

func NewExpr0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr0Context {
	var p = new(Expr0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr0

	return p
}

func (s *Expr0Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr0Context) Brackets() IBracketsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketsContext)
}

func (s *Expr0Context) StringImmidiate() IStringImmidiateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringImmidiateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringImmidiateContext)
}

func (s *Expr0Context) IntegerImmidiate() IIntegerImmidiateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerImmidiateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerImmidiateContext)
}

func (s *Expr0Context) IdentifierImmidiate() IIdentifierImmidiateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierImmidiateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierImmidiateContext)
}

func (s *Expr0Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *Expr0Context) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *Expr0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr0(s)
	}
}

func (s *Expr0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr0(s)
	}
}

func (p *AuroraParser) Expr0() (localctx IExpr0Context) {
	localctx = NewExpr0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AuroraParserRULE_expr0)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(40)
			p.Brackets()
		}

	case 2:
		{
			p.SetState(41)
			p.StringImmidiate()
		}

	case 3:
		{
			p.SetState(42)
			p.IntegerImmidiate()
		}

	case 4:
		{
			p.SetState(43)
			p.IdentifierImmidiate()
		}

	case 5:
		{
			p.SetState(44)
			p.Expr1()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(47)
			p.Expr0()
		}

	}

	return localctx
}

// IBracketsContext is an interface to support dynamic dispatch.
type IBracketsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketsContext differentiates from other interfaces.
	IsBracketsContext()
}

type BracketsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketsContext() *BracketsContext {
	var p = new(BracketsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_brackets
	return p
}

func (*BracketsContext) IsBracketsContext() {}

func NewBracketsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketsContext {
	var p = new(BracketsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_brackets

	return p
}

func (s *BracketsContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketsContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *BracketsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterBrackets(s)
	}
}

func (s *BracketsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitBrackets(s)
	}
}

func (p *AuroraParser) Brackets() (localctx IBracketsContext) {
	localctx = NewBracketsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AuroraParserRULE_brackets)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(AuroraParserT__1)
	}
	{
		p.SetState(51)
		p.Expr0()
	}
	{
		p.SetState(52)
		p.Match(AuroraParserT__2)
	}

	return localctx
}

// IStringImmidiateContext is an interface to support dynamic dispatch.
type IStringImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStringVal returns the StringVal token.
	GetStringVal() antlr.Token

	// SetStringVal sets the StringVal token.
	SetStringVal(antlr.Token)

	// IsStringImmidiateContext differentiates from other interfaces.
	IsStringImmidiateContext()
}

type StringImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	StringVal antlr.Token
}

func NewEmptyStringImmidiateContext() *StringImmidiateContext {
	var p = new(StringImmidiateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_stringImmidiate
	return p
}

func (*StringImmidiateContext) IsStringImmidiateContext() {}

func NewStringImmidiateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringImmidiateContext {
	var p = new(StringImmidiateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_stringImmidiate

	return p
}

func (s *StringImmidiateContext) GetParser() antlr.Parser { return s.parser }

func (s *StringImmidiateContext) GetStringVal() antlr.Token { return s.StringVal }

func (s *StringImmidiateContext) SetStringVal(v antlr.Token) { s.StringVal = v }

func (s *StringImmidiateContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(AuroraParserStringLiteral, 0)
}

func (s *StringImmidiateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringImmidiateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringImmidiateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterStringImmidiate(s)
	}
}

func (s *StringImmidiateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitStringImmidiate(s)
	}
}

func (p *AuroraParser) StringImmidiate() (localctx IStringImmidiateContext) {
	localctx = NewStringImmidiateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AuroraParserRULE_stringImmidiate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)

		var _m = p.Match(AuroraParserStringLiteral)

		localctx.(*StringImmidiateContext).StringVal = _m
	}

	return localctx
}

// IIdentifierImmidiateContext is an interface to support dynamic dispatch.
type IIdentifierImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name token.
	GetName() antlr.Token

	// SetName sets the Name token.
	SetName(antlr.Token)

	// IsIdentifierImmidiateContext differentiates from other interfaces.
	IsIdentifierImmidiateContext()
}

type IdentifierImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Name   antlr.Token
}

func NewEmptyIdentifierImmidiateContext() *IdentifierImmidiateContext {
	var p = new(IdentifierImmidiateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_identifierImmidiate
	return p
}

func (*IdentifierImmidiateContext) IsIdentifierImmidiateContext() {}

func NewIdentifierImmidiateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierImmidiateContext {
	var p = new(IdentifierImmidiateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_identifierImmidiate

	return p
}

func (s *IdentifierImmidiateContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierImmidiateContext) GetName() antlr.Token { return s.Name }

func (s *IdentifierImmidiateContext) SetName(v antlr.Token) { s.Name = v }

func (s *IdentifierImmidiateContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AuroraParserIdentifier, 0)
}

func (s *IdentifierImmidiateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierImmidiateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierImmidiateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterIdentifierImmidiate(s)
	}
}

func (s *IdentifierImmidiateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitIdentifierImmidiate(s)
	}
}

func (p *AuroraParser) IdentifierImmidiate() (localctx IIdentifierImmidiateContext) {
	localctx = NewIdentifierImmidiateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AuroraParserRULE_identifierImmidiate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*IdentifierImmidiateContext).Name = _m
	}

	return localctx
}

// IIntegerImmidiateContext is an interface to support dynamic dispatch.
type IIntegerImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNumberVal returns the NumberVal token.
	GetNumberVal() antlr.Token

	// SetNumberVal sets the NumberVal token.
	SetNumberVal(antlr.Token)

	// IsIntegerImmidiateContext differentiates from other interfaces.
	IsIntegerImmidiateContext()
}

type IntegerImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	NumberVal antlr.Token
}

func NewEmptyIntegerImmidiateContext() *IntegerImmidiateContext {
	var p = new(IntegerImmidiateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_integerImmidiate
	return p
}

func (*IntegerImmidiateContext) IsIntegerImmidiateContext() {}

func NewIntegerImmidiateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerImmidiateContext {
	var p = new(IntegerImmidiateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_integerImmidiate

	return p
}

func (s *IntegerImmidiateContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerImmidiateContext) GetNumberVal() antlr.Token { return s.NumberVal }

func (s *IntegerImmidiateContext) SetNumberVal(v antlr.Token) { s.NumberVal = v }

func (s *IntegerImmidiateContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(AuroraParserIntegerLiteral, 0)
}

func (s *IntegerImmidiateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerImmidiateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerImmidiateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterIntegerImmidiate(s)
	}
}

func (s *IntegerImmidiateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitIntegerImmidiate(s)
	}
}

func (p *AuroraParser) IntegerImmidiate() (localctx IIntegerImmidiateContext) {
	localctx = NewIntegerImmidiateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AuroraParserRULE_integerImmidiate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)

		var _m = p.Match(AuroraParserIntegerLiteral)

		localctx.(*IntegerImmidiateContext).NumberVal = _m
	}

	return localctx
}

// IExpr1Context is an interface to support dynamic dispatch.
type IExpr1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr1Context differentiates from other interfaces.
	IsExpr1Context()
}

type Expr1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr1Context() *Expr1Context {
	var p = new(Expr1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr1
	return p
}

func (*Expr1Context) IsExpr1Context() {}

func NewExpr1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr1Context {
	var p = new(Expr1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr1

	return p
}

func (s *Expr1Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr1Context) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *Expr1Context) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *Expr1Context) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *Expr1Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *Expr1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr1(s)
	}
}

func (s *Expr1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr1(s)
	}
}

func (p *AuroraParser) Expr1() (localctx IExpr1Context) {
	localctx = NewExpr1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AuroraParserRULE_expr1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__1:
		{
			p.SetState(60)
			p.FunctionCall()
		}

	case AuroraParserT__3:
		{
			p.SetState(61)
			p.MemberAccess()
		}

	case AuroraParserT__5, AuroraParserT__6, AuroraParserT__7, AuroraParserT__8, AuroraParserT__9:
		{
			p.SetState(62)
			p.Expr2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(65)
			p.Expr1()
		}

	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMember returns the Member token.
	GetMember() antlr.Token

	// SetMember sets the Member token.
	SetMember(antlr.Token)

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Member antlr.Token
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) GetMember() antlr.Token { return s.Member }

func (s *MemberAccessContext) SetMember(v antlr.Token) { s.Member = v }

func (s *MemberAccessContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AuroraParserIdentifier, 0)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (p *AuroraParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AuroraParserRULE_memberAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(AuroraParserT__3)
	}
	{
		p.SetState(69)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*MemberAccessContext).Member = _m
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AllFunctionCallParam() []IFunctionCallParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallParamContext)(nil)).Elem())
	var tst = make([]IFunctionCallParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallParamContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) FunctionCallParam(i int) IFunctionCallParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallParamContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *AuroraParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AuroraParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(AuroraParserT__1)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__1)|(1<<AuroraParserT__3)|(1<<AuroraParserT__5)|(1<<AuroraParserT__6)|(1<<AuroraParserT__7)|(1<<AuroraParserT__8)|(1<<AuroraParserT__9)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(72)
			p.FunctionCallParam()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__4 {
		{
			p.SetState(75)
			p.Match(AuroraParserT__4)
		}
		{
			p.SetState(76)
			p.FunctionCallParam()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(AuroraParserT__2)
	}

	return localctx
}

// IFunctionCallParamContext is an interface to support dynamic dispatch.
type IFunctionCallParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParamExpr returns the ParamExpr rule contexts.
	GetParamExpr() IExpr0Context

	// SetParamExpr sets the ParamExpr rule contexts.
	SetParamExpr(IExpr0Context)

	// IsFunctionCallParamContext differentiates from other interfaces.
	IsFunctionCallParamContext()
}

type FunctionCallParamContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ParamExpr IExpr0Context
}

func NewEmptyFunctionCallParamContext() *FunctionCallParamContext {
	var p = new(FunctionCallParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_functionCallParam
	return p
}

func (*FunctionCallParamContext) IsFunctionCallParamContext() {}

func NewFunctionCallParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallParamContext {
	var p = new(FunctionCallParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_functionCallParam

	return p
}

func (s *FunctionCallParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallParamContext) GetParamExpr() IExpr0Context { return s.ParamExpr }

func (s *FunctionCallParamContext) SetParamExpr(v IExpr0Context) { s.ParamExpr = v }

func (s *FunctionCallParamContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *FunctionCallParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterFunctionCallParam(s)
	}
}

func (s *FunctionCallParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitFunctionCallParam(s)
	}
}

func (p *AuroraParser) FunctionCallParam() (localctx IFunctionCallParamContext) {
	localctx = NewFunctionCallParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AuroraParserRULE_functionCallParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)

		var _x = p.Expr0()

		localctx.(*FunctionCallParamContext).ParamExpr = _x
	}

	return localctx
}

// IExpr2Context is an interface to support dynamic dispatch.
type IExpr2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the Op token.
	GetOp() antlr.Token

	// SetOp sets the Op token.
	SetOp(antlr.Token)

	// GetExpr returns the Expr rule contexts.
	GetExpr() IExpr0Context

	// SetExpr sets the Expr rule contexts.
	SetExpr(IExpr0Context)

	// IsExpr2Context differentiates from other interfaces.
	IsExpr2Context()
}

type Expr2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Op     antlr.Token
	Expr   IExpr0Context
}

func NewEmptyExpr2Context() *Expr2Context {
	var p = new(Expr2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr2
	return p
}

func (*Expr2Context) IsExpr2Context() {}

func NewExpr2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr2Context {
	var p = new(Expr2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr2

	return p
}

func (s *Expr2Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr2Context) GetOp() antlr.Token { return s.Op }

func (s *Expr2Context) SetOp(v antlr.Token) { s.Op = v }

func (s *Expr2Context) GetExpr() IExpr0Context { return s.Expr }

func (s *Expr2Context) SetExpr(v IExpr0Context) { s.Expr = v }

func (s *Expr2Context) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *Expr2Context) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *Expr2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr2(s)
	}
}

func (s *Expr2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr2(s)
	}
}

func (p *AuroraParser) Expr2() (localctx IExpr2Context) {
	localctx = NewExpr2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AuroraParserRULE_expr2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*Expr2Context).Op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__5)|(1<<AuroraParserT__6)|(1<<AuroraParserT__7)|(1<<AuroraParserT__8)|(1<<AuroraParserT__9))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*Expr2Context).Op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(87)

		var _x = p.Expr0()

		localctx.(*Expr2Context).Expr = _x
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(88)
			p.Expr2()
		}

	}

	return localctx
}
