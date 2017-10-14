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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 131,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 6, 12, 66, 10, 12, 13,
	12, 14, 12, 67, 3, 13, 3, 13, 7, 13, 72, 10, 13, 12, 13, 14, 13, 75, 11,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 82, 10, 14, 12, 14, 14, 14,
	85, 11, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 6, 18, 94,
	10, 18, 13, 18, 14, 18, 95, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19, 102, 10,
	19, 3, 19, 5, 19, 105, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 111,
	10, 20, 12, 20, 14, 20, 114, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 125, 10, 21, 12, 21, 14, 21, 128, 11,
	21, 3, 21, 3, 21, 3, 112, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 2, 31, 2, 33,
	2, 35, 16, 37, 17, 39, 18, 41, 19, 3, 2, 7, 3, 2, 36, 36, 5, 2, 67, 92,
	97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15,
	15, 2, 137, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45, 3, 2,
	2, 2, 7, 47, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13, 53,
	3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 57, 3, 2, 2, 2, 19, 59, 3, 2, 2, 2,
	21, 61, 3, 2, 2, 2, 23, 65, 3, 2, 2, 2, 25, 69, 3, 2, 2, 2, 27, 78, 3,
	2, 2, 2, 29, 86, 3, 2, 2, 2, 31, 88, 3, 2, 2, 2, 33, 90, 3, 2, 2, 2, 35,
	93, 3, 2, 2, 2, 37, 104, 3, 2, 2, 2, 39, 106, 3, 2, 2, 2, 41, 120, 3, 2,
	2, 2, 43, 44, 7, 61, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 42, 2, 2, 46,
	6, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7, 48, 2,
	2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 46, 2, 2, 52, 12, 3, 2, 2, 2, 53, 54,
	7, 47, 2, 2, 54, 14, 3, 2, 2, 2, 55, 56, 7, 45, 2, 2, 56, 16, 3, 2, 2,
	2, 57, 58, 7, 35, 2, 2, 58, 18, 3, 2, 2, 2, 59, 60, 7, 44, 2, 2, 60, 20,
	3, 2, 2, 2, 61, 62, 7, 40, 2, 2, 62, 22, 3, 2, 2, 2, 63, 66, 5, 33, 17,
	2, 64, 66, 7, 97, 2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 24, 3, 2, 2, 2,
	69, 73, 7, 36, 2, 2, 70, 72, 10, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3,
	2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75,
	73, 3, 2, 2, 2, 76, 77, 7, 36, 2, 2, 77, 26, 3, 2, 2, 2, 78, 83, 5, 29,
	15, 2, 79, 82, 5, 29, 15, 2, 80, 82, 5, 33, 17, 2, 81, 79, 3, 2, 2, 2,
	81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3,
	2, 2, 2, 84, 28, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 5, 31, 16, 2,
	87, 30, 3, 2, 2, 2, 88, 89, 9, 3, 2, 2, 89, 32, 3, 2, 2, 2, 90, 91, 9,
	4, 2, 2, 91, 34, 3, 2, 2, 2, 92, 94, 9, 5, 2, 2, 93, 92, 3, 2, 2, 2, 94,
	95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 98, 8, 18, 2, 2, 98, 36, 3, 2, 2, 2, 99, 101, 7, 15, 2, 2, 100,
	102, 7, 12, 2, 2, 101, 100, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 105,
	3, 2, 2, 2, 103, 105, 7, 12, 2, 2, 104, 99, 3, 2, 2, 2, 104, 103, 3, 2,
	2, 2, 105, 38, 3, 2, 2, 2, 106, 107, 7, 49, 2, 2, 107, 108, 7, 44, 2, 2,
	108, 112, 3, 2, 2, 2, 109, 111, 11, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111,
	114, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 115,
	3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 44, 2, 2, 116, 117, 7, 49,
	2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 8, 20, 2, 2, 119, 40, 3, 2, 2, 2,
	120, 121, 7, 49, 2, 2, 121, 122, 7, 49, 2, 2, 122, 126, 3, 2, 2, 2, 123,
	125, 10, 6, 2, 2, 124, 123, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 126, 3, 2,
	2, 2, 129, 130, 8, 21, 2, 2, 130, 42, 3, 2, 2, 2, 13, 2, 65, 67, 73, 81,
	83, 95, 101, 104, 112, 126, 3, 8, 2, 2,
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
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "IntegerLiteral", "StringLiteral",
	"Identifier", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "IntegerLiteral", "StringLiteral", "Identifier", "IdentifierNondigit",
	"Nondigit", "Digit", "Whitespace", "Newline", "BlockComment", "LineComment",
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
	AuroraLexerIntegerLiteral = 11
	AuroraLexerStringLiteral  = 12
	AuroraLexerIdentifier     = 13
	AuroraLexerWhitespace     = 14
	AuroraLexerNewline        = 15
	AuroraLexerBlockComment   = 16
	AuroraLexerLineComment    = 17
)

var IdSeen bool = false
