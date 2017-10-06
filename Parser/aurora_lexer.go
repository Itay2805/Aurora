// Generated from c:\Users\Itay Almog\Desktop\GoWorkspace\src\Aurora\Parser\Aurora.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 58,
	10, 10, 12, 10, 14, 10, 61, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 6, 14, 70, 10, 14, 13, 14, 14, 14, 71, 3, 14, 3, 14, 3, 15,
	3, 15, 5, 15, 78, 10, 15, 3, 15, 5, 15, 81, 10, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 89, 10, 16, 12, 16, 14, 16, 92, 11, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 103,
	10, 17, 12, 17, 14, 17, 106, 11, 17, 3, 17, 3, 17, 3, 90, 2, 18, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 2, 23, 2, 25,
	2, 27, 12, 29, 13, 31, 14, 33, 15, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99, 124,
	3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 112, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 40, 3, 2, 2, 2, 7, 42, 3, 2, 2,
	2, 9, 44, 3, 2, 2, 2, 11, 46, 3, 2, 2, 2, 13, 48, 3, 2, 2, 2, 15, 50, 3,
	2, 2, 2, 17, 52, 3, 2, 2, 2, 19, 54, 3, 2, 2, 2, 21, 62, 3, 2, 2, 2, 23,
	64, 3, 2, 2, 2, 25, 66, 3, 2, 2, 2, 27, 69, 3, 2, 2, 2, 29, 80, 3, 2, 2,
	2, 31, 84, 3, 2, 2, 2, 33, 98, 3, 2, 2, 2, 35, 36, 7, 104, 2, 2, 36, 37,
	7, 119, 2, 2, 37, 38, 7, 112, 2, 2, 38, 39, 7, 101, 2, 2, 39, 4, 3, 2,
	2, 2, 40, 41, 7, 42, 2, 2, 41, 6, 3, 2, 2, 2, 42, 43, 7, 43, 2, 2, 43,
	8, 3, 2, 2, 2, 44, 45, 7, 60, 2, 2, 45, 10, 3, 2, 2, 2, 46, 47, 7, 46,
	2, 2, 47, 12, 3, 2, 2, 2, 48, 49, 7, 125, 2, 2, 49, 14, 3, 2, 2, 2, 50,
	51, 7, 127, 2, 2, 51, 16, 3, 2, 2, 2, 52, 53, 7, 61, 2, 2, 53, 18, 3, 2,
	2, 2, 54, 59, 5, 21, 11, 2, 55, 58, 5, 21, 11, 2, 56, 58, 5, 25, 13, 2,
	57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3,
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 20, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62,
	63, 5, 23, 12, 2, 63, 22, 3, 2, 2, 2, 64, 65, 9, 2, 2, 2, 65, 24, 3, 2,
	2, 2, 66, 67, 9, 3, 2, 2, 67, 26, 3, 2, 2, 2, 68, 70, 9, 4, 2, 2, 69, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 73, 3, 2, 2, 2, 73, 74, 8, 14, 2, 2, 74, 28, 3, 2, 2, 2, 75, 77, 7,
	15, 2, 2, 76, 78, 7, 12, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 81, 3, 2, 2, 2, 79, 81, 7, 12, 2, 2, 80, 75, 3, 2, 2, 2, 80, 79, 3,
	2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 8, 15, 2, 2, 83, 30, 3, 2, 2, 2, 84,
	85, 7, 49, 2, 2, 85, 86, 7, 44, 2, 2, 86, 90, 3, 2, 2, 2, 87, 89, 11, 2,
	2, 2, 88, 87, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 44, 2, 2,
	94, 95, 7, 49, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8, 16, 2, 2, 97, 32, 3,
	2, 2, 2, 98, 99, 7, 49, 2, 2, 99, 100, 7, 49, 2, 2, 100, 104, 3, 2, 2,
	2, 101, 103, 10, 5, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 107, 108, 8, 17, 2, 2, 108, 34, 3, 2, 2, 2, 10, 2, 57, 59,
	71, 77, 80, 90, 104, 3, 8, 2, 2,
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
	"", "'func'", "'('", "')'", "':'", "','", "'{'", "'}'", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "Identifier", "Whitespace", "Newline",
	"BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "Identifier",
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
	AuroraLexerT__0         = 1
	AuroraLexerT__1         = 2
	AuroraLexerT__2         = 3
	AuroraLexerT__3         = 4
	AuroraLexerT__4         = 5
	AuroraLexerT__5         = 6
	AuroraLexerT__6         = 7
	AuroraLexerT__7         = 8
	AuroraLexerIdentifier   = 9
	AuroraLexerWhitespace   = 10
	AuroraLexerNewline      = 11
	AuroraLexerBlockComment = 12
	AuroraLexerLineComment  = 13
)
