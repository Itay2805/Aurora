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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 139,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 6, 14, 74, 10, 14, 13, 14, 14, 14, 75, 3, 15,
	3, 15, 7, 15, 80, 10, 15, 12, 15, 14, 15, 83, 11, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 90, 10, 16, 12, 16, 14, 16, 93, 11, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 6, 20, 102, 10, 20, 13, 20, 14,
	20, 103, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 110, 10, 21, 3, 21, 5, 21,
	113, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 119, 10, 22, 12, 22, 14,
	22, 122, 11, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 7, 23, 133, 10, 23, 12, 23, 14, 23, 136, 11, 23, 3, 23, 3, 23, 3,
	120, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 2, 35, 2, 37, 2, 39,
	18, 41, 19, 43, 20, 45, 21, 3, 2, 7, 3, 2, 36, 36, 5, 2, 67, 92, 97, 97,
	99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 145,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47,
	3, 2, 2, 2, 5, 49, 3, 2, 2, 2, 7, 51, 3, 2, 2, 2, 9, 53, 3, 2, 2, 2, 11,
	55, 3, 2, 2, 2, 13, 57, 3, 2, 2, 2, 15, 59, 3, 2, 2, 2, 17, 61, 3, 2, 2,
	2, 19, 63, 3, 2, 2, 2, 21, 65, 3, 2, 2, 2, 23, 67, 3, 2, 2, 2, 25, 69,
	3, 2, 2, 2, 27, 73, 3, 2, 2, 2, 29, 77, 3, 2, 2, 2, 31, 86, 3, 2, 2, 2,
	33, 94, 3, 2, 2, 2, 35, 96, 3, 2, 2, 2, 37, 98, 3, 2, 2, 2, 39, 101, 3,
	2, 2, 2, 41, 112, 3, 2, 2, 2, 43, 114, 3, 2, 2, 2, 45, 128, 3, 2, 2, 2,
	47, 48, 7, 61, 2, 2, 48, 4, 3, 2, 2, 2, 49, 50, 7, 42, 2, 2, 50, 6, 3,
	2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 8, 3, 2, 2, 2, 53, 54, 7, 48, 2, 2, 54,
	10, 3, 2, 2, 2, 55, 56, 7, 46, 2, 2, 56, 12, 3, 2, 2, 2, 57, 58, 7, 47,
	2, 2, 58, 14, 3, 2, 2, 2, 59, 60, 7, 45, 2, 2, 60, 16, 3, 2, 2, 2, 61,
	62, 7, 35, 2, 2, 62, 18, 3, 2, 2, 2, 63, 64, 7, 44, 2, 2, 64, 20, 3, 2,
	2, 2, 65, 66, 7, 40, 2, 2, 66, 22, 3, 2, 2, 2, 67, 68, 7, 49, 2, 2, 68,
	24, 3, 2, 2, 2, 69, 70, 7, 39, 2, 2, 70, 26, 3, 2, 2, 2, 71, 74, 5, 37,
	19, 2, 72, 74, 7, 97, 2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74,
	75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 28, 3, 2, 2,
	2, 77, 81, 7, 36, 2, 2, 78, 80, 10, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 83,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2,
	83, 81, 3, 2, 2, 2, 84, 85, 7, 36, 2, 2, 85, 30, 3, 2, 2, 2, 86, 91, 5,
	33, 17, 2, 87, 90, 5, 33, 17, 2, 88, 90, 5, 37, 19, 2, 89, 87, 3, 2, 2,
	2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 32, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 5, 35, 18,
	2, 95, 34, 3, 2, 2, 2, 96, 97, 9, 3, 2, 2, 97, 36, 3, 2, 2, 2, 98, 99,
	9, 4, 2, 2, 99, 38, 3, 2, 2, 2, 100, 102, 9, 5, 2, 2, 101, 100, 3, 2, 2,
	2, 102, 103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	105, 3, 2, 2, 2, 105, 106, 8, 20, 2, 2, 106, 40, 3, 2, 2, 2, 107, 109,
	7, 15, 2, 2, 108, 110, 7, 12, 2, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3,
	2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 113, 7, 12, 2, 2, 112, 107, 3, 2, 2,
	2, 112, 111, 3, 2, 2, 2, 113, 42, 3, 2, 2, 2, 114, 115, 7, 49, 2, 2, 115,
	116, 7, 44, 2, 2, 116, 120, 3, 2, 2, 2, 117, 119, 11, 2, 2, 2, 118, 117,
	3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 44, 2, 2,
	124, 125, 7, 49, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 8, 22, 2, 2, 127,
	44, 3, 2, 2, 2, 128, 129, 7, 49, 2, 2, 129, 130, 7, 49, 2, 2, 130, 134,
	3, 2, 2, 2, 131, 133, 10, 6, 2, 2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2,
	2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2,
	136, 134, 3, 2, 2, 2, 137, 138, 8, 23, 2, 2, 138, 46, 3, 2, 2, 2, 13, 2,
	73, 75, 81, 89, 91, 103, 109, 112, 120, 134, 3, 8, 2, 2,
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
	"", "';'", "'('", "')'", "'.'", "','", "'-'", "'+'", "'!'", "'*'", "'&'",
	"'/'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "IntegerLiteral", "StringLiteral",
	"Identifier", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "IntegerLiteral", "StringLiteral", "Identifier",
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
	AuroraLexerT__4           = 5
	AuroraLexerT__5           = 6
	AuroraLexerT__6           = 7
	AuroraLexerT__7           = 8
	AuroraLexerT__8           = 9
	AuroraLexerT__9           = 10
	AuroraLexerT__10          = 11
	AuroraLexerT__11          = 12
	AuroraLexerIntegerLiteral = 13
	AuroraLexerStringLiteral  = 14
	AuroraLexerIdentifier     = 15
	AuroraLexerWhitespace     = 16
	AuroraLexerNewline        = 17
	AuroraLexerBlockComment   = 18
	AuroraLexerLineComment    = 19
)

var IdSeen bool = false
