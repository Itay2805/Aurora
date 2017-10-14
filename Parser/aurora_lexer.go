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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 111,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 46, 10,
	7, 13, 7, 14, 7, 47, 3, 8, 3, 8, 7, 8, 52, 10, 8, 12, 8, 14, 8, 55, 11,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 62, 10, 9, 12, 9, 14, 9, 65, 11,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 6, 13, 74, 10, 13,
	13, 13, 14, 13, 75, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 82, 10, 14, 3, 14,
	5, 14, 85, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 91, 10, 15, 12, 15,
	14, 15, 94, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 7, 16, 105, 10, 16, 12, 16, 14, 16, 108, 11, 16, 3, 16, 3, 16,
	3, 92, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	2, 21, 2, 23, 2, 25, 11, 27, 12, 29, 13, 31, 14, 3, 2, 7, 3, 2, 36, 36,
	5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2,
	12, 12, 15, 15, 2, 117, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2, 2, 7,
	37, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 45, 3, 2, 2,
	2, 15, 49, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2, 19, 66, 3, 2, 2, 2, 21, 68,
	3, 2, 2, 2, 23, 70, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2, 27, 84, 3, 2, 2, 2,
	29, 86, 3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33, 34, 7, 61, 2, 2, 34, 4, 3,
	2, 2, 2, 35, 36, 7, 42, 2, 2, 36, 6, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38,
	8, 3, 2, 2, 2, 39, 40, 7, 48, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 46,
	2, 2, 42, 12, 3, 2, 2, 2, 43, 46, 5, 23, 12, 2, 44, 46, 7, 97, 2, 2, 45,
	43, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2,
	2, 47, 48, 3, 2, 2, 2, 48, 14, 3, 2, 2, 2, 49, 53, 7, 36, 2, 2, 50, 52,
	10, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 57, 7,
	36, 2, 2, 57, 16, 3, 2, 2, 2, 58, 63, 5, 19, 10, 2, 59, 62, 5, 19, 10,
	2, 60, 62, 5, 23, 12, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 65,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 18, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 66, 67, 5, 21, 11, 2, 67, 20, 3, 2, 2, 2, 68, 69, 9,
	3, 2, 2, 69, 22, 3, 2, 2, 2, 70, 71, 9, 4, 2, 2, 71, 24, 3, 2, 2, 2, 72,
	74, 9, 5, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2,
	2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 8, 13, 2, 2, 78, 26,
	3, 2, 2, 2, 79, 81, 7, 15, 2, 2, 80, 82, 7, 12, 2, 2, 81, 80, 3, 2, 2,
	2, 81, 82, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 85, 7, 12, 2, 2, 84, 79,
	3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 28, 3, 2, 2, 2, 86, 87, 7, 49, 2, 2,
	87, 88, 7, 44, 2, 2, 88, 92, 3, 2, 2, 2, 89, 91, 11, 2, 2, 2, 90, 89, 3,
	2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93,
	95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96, 7, 44, 2, 2, 96, 97, 7, 49,
	2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 8, 15, 2, 2, 99, 30, 3, 2, 2, 2, 100,
	101, 7, 49, 2, 2, 101, 102, 7, 49, 2, 2, 102, 106, 3, 2, 2, 2, 103, 105,
	10, 6, 2, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2,
	109, 110, 8, 16, 2, 2, 110, 32, 3, 2, 2, 2, 13, 2, 45, 47, 53, 61, 63,
	75, 81, 84, 92, 106, 3, 8, 2, 2,
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
	"", "';'", "'('", "')'", "'.'", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "IntegerLiteral", "StringLiteral", "Identifier",
	"Whitespace", "Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "IntegerLiteral", "StringLiteral",
	"Identifier", "IdentifierNondigit", "Nondigit", "Digit", "Whitespace",
	"Newline", "BlockComment", "LineComment",
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
	AuroraLexerIntegerLiteral = 6
	AuroraLexerStringLiteral  = 7
	AuroraLexerIdentifier     = 8
	AuroraLexerWhitespace     = 9
	AuroraLexerNewline        = 10
	AuroraLexerBlockComment   = 11
	AuroraLexerLineComment    = 12
)

var IdSeen bool = false
