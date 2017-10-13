// Generated from d:\Itay\Documents\Coding\Go\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 6, 6, 42, 10, 6, 13, 6, 14, 6, 43, 3, 7,
	3, 7, 7, 7, 48, 10, 7, 12, 7, 14, 7, 51, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 7, 8, 58, 10, 8, 12, 8, 14, 8, 61, 11, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 6, 12, 70, 10, 12, 13, 12, 14, 12, 71, 3, 12, 3, 12,
	3, 13, 3, 13, 5, 13, 78, 10, 13, 3, 13, 5, 13, 81, 10, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 89, 10, 14, 12, 14, 14, 14, 92, 11,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15,
	103, 10, 15, 12, 15, 14, 15, 106, 11, 15, 3, 15, 3, 15, 3, 90, 2, 16, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 2, 21, 2, 23, 10,
	25, 11, 27, 12, 29, 13, 3, 2, 7, 4, 2, 36, 36, 96, 96, 5, 2, 67, 92, 97,
	97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15,
	2, 115, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 3,
	31, 3, 2, 2, 2, 5, 33, 3, 2, 2, 2, 7, 35, 3, 2, 2, 2, 9, 37, 3, 2, 2, 2,
	11, 41, 3, 2, 2, 2, 13, 45, 3, 2, 2, 2, 15, 54, 3, 2, 2, 2, 17, 62, 3,
	2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 66, 3, 2, 2, 2, 23, 69, 3, 2, 2, 2, 25,
	80, 3, 2, 2, 2, 27, 84, 3, 2, 2, 2, 29, 98, 3, 2, 2, 2, 31, 32, 7, 48,
	2, 2, 32, 4, 3, 2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 6, 3, 2, 2, 2, 35, 36,
	7, 46, 2, 2, 36, 8, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 10, 3, 2, 2, 2,
	39, 42, 5, 21, 11, 2, 40, 42, 7, 97, 2, 2, 41, 39, 3, 2, 2, 2, 41, 40,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2,
	44, 12, 3, 2, 2, 2, 45, 49, 7, 36, 2, 2, 46, 48, 9, 2, 2, 2, 47, 46, 3,
	2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 36, 2, 2, 53, 14, 3, 2,
	2, 2, 54, 59, 5, 17, 9, 2, 55, 58, 5, 17, 9, 2, 56, 58, 5, 21, 11, 2, 57,
	55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2,
	2, 59, 60, 3, 2, 2, 2, 60, 16, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63,
	5, 19, 10, 2, 63, 18, 3, 2, 2, 2, 64, 65, 9, 3, 2, 2, 65, 20, 3, 2, 2,
	2, 66, 67, 9, 4, 2, 2, 67, 22, 3, 2, 2, 2, 68, 70, 9, 5, 2, 2, 69, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 73, 3, 2, 2, 2, 73, 74, 8, 12, 2, 2, 74, 24, 3, 2, 2, 2, 75, 77, 7,
	15, 2, 2, 76, 78, 7, 12, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 81, 3, 2, 2, 2, 79, 81, 7, 12, 2, 2, 80, 75, 3, 2, 2, 2, 80, 79, 3,
	2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 8, 13, 2, 2, 83, 26, 3, 2, 2, 2, 84,
	85, 7, 49, 2, 2, 85, 86, 7, 44, 2, 2, 86, 90, 3, 2, 2, 2, 87, 89, 11, 2,
	2, 2, 88, 87, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 44, 2, 2,
	94, 95, 7, 49, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8, 14, 2, 2, 97, 28, 3,
	2, 2, 2, 98, 99, 7, 49, 2, 2, 99, 100, 7, 49, 2, 2, 100, 104, 3, 2, 2,
	2, 101, 103, 10, 6, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 107, 108, 8, 15, 2, 2, 108, 30, 3, 2, 2, 2, 13, 2, 41, 43,
	49, 57, 59, 71, 77, 80, 90, 104, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'('", "','", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "IntegerLiteral", "StringLiteral", "Identifier", "Whitespace",
	"Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "IntegerLiteral", "StringLiteral", "Identifier",
	"IdentifierNondigit", "Nondigit", "Digit", "Whitespace", "Newline", "BlockComment",
	"LineComment",
}

type AuroraLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAuroraLexer(input antlr.CharStream) *AuroraLexer {

	l := new(AuroraLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Aurora.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AuroraLexer tokens.
const (
	AuroraLexerT__0           = 1
	AuroraLexerT__1           = 2
	AuroraLexerT__2           = 3
	AuroraLexerT__3           = 4
	AuroraLexerIntegerLiteral = 5
	AuroraLexerStringLiteral  = 6
	AuroraLexerIdentifier     = 7
	AuroraLexerWhitespace     = 8
	AuroraLexerNewline        = 9
	AuroraLexerBlockComment   = 10
	AuroraLexerLineComment    = 11
)
