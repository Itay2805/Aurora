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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 105,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 6, 5, 38, 10, 5, 13, 5, 14, 5, 39, 3, 6, 3, 6, 7, 6, 44, 10, 6, 12,
	6, 14, 6, 47, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 54, 10, 7, 12,
	7, 14, 7, 57, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11,
	66, 10, 11, 13, 11, 14, 11, 67, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 74,
	10, 12, 3, 12, 5, 12, 77, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 7, 13, 85, 10, 13, 12, 13, 14, 13, 88, 11, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 99, 10, 14, 12, 14, 14,
	14, 102, 11, 14, 3, 14, 3, 14, 3, 86, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 2, 17, 2, 19, 2, 21, 9, 23, 10, 25, 11, 27, 12, 3, 2, 7,
	4, 2, 36, 36, 96, 96, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2,
	11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 111, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2,
	9, 37, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 50, 3, 2, 2, 2, 15, 58, 3, 2,
	2, 2, 17, 60, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 65, 3, 2, 2, 2, 23, 76,
	3, 2, 2, 2, 25, 80, 3, 2, 2, 2, 27, 94, 3, 2, 2, 2, 29, 30, 7, 48, 2, 2,
	30, 4, 3, 2, 2, 2, 31, 32, 7, 42, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 43,
	2, 2, 34, 8, 3, 2, 2, 2, 35, 38, 5, 19, 10, 2, 36, 38, 7, 97, 2, 2, 37,
	35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2,
	2, 39, 40, 3, 2, 2, 2, 40, 10, 3, 2, 2, 2, 41, 45, 7, 36, 2, 2, 42, 44,
	9, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2,
	45, 46, 3, 2, 2, 2, 46, 48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7,
	36, 2, 2, 49, 12, 3, 2, 2, 2, 50, 55, 5, 15, 8, 2, 51, 54, 5, 15, 8, 2,
	52, 54, 5, 19, 10, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3,
	2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 14, 3, 2, 2, 2, 57,
	55, 3, 2, 2, 2, 58, 59, 5, 17, 9, 2, 59, 16, 3, 2, 2, 2, 60, 61, 9, 3,
	2, 2, 61, 18, 3, 2, 2, 2, 62, 63, 9, 4, 2, 2, 63, 20, 3, 2, 2, 2, 64, 66,
	9, 5, 2, 2, 65, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 8, 11, 2, 2, 70, 22, 3,
	2, 2, 2, 71, 73, 7, 15, 2, 2, 72, 74, 7, 12, 2, 2, 73, 72, 3, 2, 2, 2,
	73, 74, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 77, 7, 12, 2, 2, 76, 71, 3,
	2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 8, 12, 2, 2, 79,
	24, 3, 2, 2, 2, 80, 81, 7, 49, 2, 2, 81, 82, 7, 44, 2, 2, 82, 86, 3, 2,
	2, 2, 83, 85, 11, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 89, 90, 7, 44, 2, 2, 90, 91, 7, 49, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93,
	8, 13, 2, 2, 93, 26, 3, 2, 2, 2, 94, 95, 7, 49, 2, 2, 95, 96, 7, 49, 2,
	2, 96, 100, 3, 2, 2, 2, 97, 99, 10, 6, 2, 2, 98, 97, 3, 2, 2, 2, 99, 102,
	3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 8, 14, 2, 2, 104, 28, 3, 2, 2, 2,
	13, 2, 37, 39, 45, 53, 55, 67, 73, 76, 86, 100, 3, 8, 2, 2,
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
	"", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "IntegerLiteral", "StringLiteral", "Identifier", "Whitespace",
	"Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "IntegerLiteral", "StringLiteral", "Identifier",
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
	AuroraLexerIntegerLiteral = 4
	AuroraLexerStringLiteral  = 5
	AuroraLexerIdentifier     = 6
	AuroraLexerWhitespace     = 7
	AuroraLexerNewline        = 8
	AuroraLexerBlockComment   = 9
	AuroraLexerLineComment    = 10
)
