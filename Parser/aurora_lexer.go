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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 167,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 6, 20, 102, 10, 20, 13, 20, 14, 20, 103, 3, 21, 3, 21, 7, 21, 108,
	10, 21, 12, 21, 14, 21, 111, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	7, 22, 118, 10, 22, 12, 22, 14, 22, 121, 11, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 6, 26, 130, 10, 26, 13, 26, 14, 26, 131, 3, 26,
	3, 26, 3, 27, 3, 27, 5, 27, 138, 10, 27, 3, 27, 5, 27, 141, 10, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 7, 28, 147, 10, 28, 12, 28, 14, 28, 150, 11, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 161,
	10, 29, 12, 29, 14, 29, 164, 11, 29, 3, 29, 3, 29, 3, 148, 2, 30, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 2, 47, 2, 49, 2, 51, 24, 53, 25, 55, 26, 57, 27, 3, 2, 7, 3,
	2, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34,
	34, 4, 2, 12, 12, 15, 15, 2, 173, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 3,
	59, 3, 2, 2, 2, 5, 61, 3, 2, 2, 2, 7, 63, 3, 2, 2, 2, 9, 67, 3, 2, 2, 2,
	11, 69, 3, 2, 2, 2, 13, 71, 3, 2, 2, 2, 15, 73, 3, 2, 2, 2, 17, 75, 3,
	2, 2, 2, 19, 77, 3, 2, 2, 2, 21, 79, 3, 2, 2, 2, 23, 81, 3, 2, 2, 2, 25,
	83, 3, 2, 2, 2, 27, 85, 3, 2, 2, 2, 29, 87, 3, 2, 2, 2, 31, 89, 3, 2, 2,
	2, 33, 91, 3, 2, 2, 2, 35, 93, 3, 2, 2, 2, 37, 96, 3, 2, 2, 2, 39, 101,
	3, 2, 2, 2, 41, 105, 3, 2, 2, 2, 43, 114, 3, 2, 2, 2, 45, 122, 3, 2, 2,
	2, 47, 124, 3, 2, 2, 2, 49, 126, 3, 2, 2, 2, 51, 129, 3, 2, 2, 2, 53, 140,
	3, 2, 2, 2, 55, 142, 3, 2, 2, 2, 57, 156, 3, 2, 2, 2, 59, 60, 7, 125, 2,
	2, 60, 4, 3, 2, 2, 2, 61, 62, 7, 127, 2, 2, 62, 6, 3, 2, 2, 2, 63, 64,
	7, 120, 2, 2, 64, 65, 7, 99, 2, 2, 65, 66, 7, 116, 2, 2, 66, 8, 3, 2, 2,
	2, 67, 68, 7, 63, 2, 2, 68, 10, 3, 2, 2, 2, 69, 70, 7, 61, 2, 2, 70, 12,
	3, 2, 2, 2, 71, 72, 7, 42, 2, 2, 72, 14, 3, 2, 2, 2, 73, 74, 7, 43, 2,
	2, 74, 16, 3, 2, 2, 2, 75, 76, 7, 48, 2, 2, 76, 18, 3, 2, 2, 2, 77, 78,
	7, 46, 2, 2, 78, 20, 3, 2, 2, 2, 79, 80, 7, 47, 2, 2, 80, 22, 3, 2, 2,
	2, 81, 82, 7, 45, 2, 2, 82, 24, 3, 2, 2, 2, 83, 84, 7, 35, 2, 2, 84, 26,
	3, 2, 2, 2, 85, 86, 7, 44, 2, 2, 86, 28, 3, 2, 2, 2, 87, 88, 7, 40, 2,
	2, 88, 30, 3, 2, 2, 2, 89, 90, 7, 49, 2, 2, 90, 32, 3, 2, 2, 2, 91, 92,
	7, 39, 2, 2, 92, 34, 3, 2, 2, 2, 93, 94, 7, 40, 2, 2, 94, 95, 7, 40, 2,
	2, 95, 36, 3, 2, 2, 2, 96, 97, 7, 126, 2, 2, 97, 98, 7, 126, 2, 2, 98,
	38, 3, 2, 2, 2, 99, 102, 5, 49, 25, 2, 100, 102, 7, 97, 2, 2, 101, 99,
	3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 101, 3, 2,
	2, 2, 103, 104, 3, 2, 2, 2, 104, 40, 3, 2, 2, 2, 105, 109, 7, 36, 2, 2,
	106, 108, 10, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109,
	107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 112, 113, 7, 36, 2, 2, 113, 42, 3, 2, 2, 2, 114, 119, 5, 45,
	23, 2, 115, 118, 5, 45, 23, 2, 116, 118, 5, 49, 25, 2, 117, 115, 3, 2,
	2, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2,
	119, 120, 3, 2, 2, 2, 120, 44, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123,
	5, 47, 24, 2, 123, 46, 3, 2, 2, 2, 124, 125, 9, 3, 2, 2, 125, 48, 3, 2,
	2, 2, 126, 127, 9, 4, 2, 2, 127, 50, 3, 2, 2, 2, 128, 130, 9, 5, 2, 2,
	129, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131,
	132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 8, 26, 2, 2, 134, 52,
	3, 2, 2, 2, 135, 137, 7, 15, 2, 2, 136, 138, 7, 12, 2, 2, 137, 136, 3,
	2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139, 141, 7, 12, 2,
	2, 140, 135, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 54, 3, 2, 2, 2, 142,
	143, 7, 49, 2, 2, 143, 144, 7, 44, 2, 2, 144, 148, 3, 2, 2, 2, 145, 147,
	11, 2, 2, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2,
	151, 152, 7, 44, 2, 2, 152, 153, 7, 49, 2, 2, 153, 154, 3, 2, 2, 2, 154,
	155, 8, 28, 2, 2, 155, 56, 3, 2, 2, 2, 156, 157, 7, 49, 2, 2, 157, 158,
	7, 49, 2, 2, 158, 162, 3, 2, 2, 2, 159, 161, 10, 6, 2, 2, 160, 159, 3,
	2, 2, 2, 161, 164, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2,
	2, 163, 165, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 166, 8, 29, 2, 2, 166,
	58, 3, 2, 2, 2, 13, 2, 101, 103, 109, 117, 119, 131, 137, 140, 148, 162,
	3, 8, 2, 2,
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
	"", "'{'", "'}'", "'var'", "'='", "';'", "'('", "')'", "'.'", "','", "'-'",
	"'+'", "'!'", "'*'", "'&'", "'/'", "'%'", "'&&'", "'||'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "IntegerLiteral", "StringLiteral", "Identifier", "Whitespace", "Newline",
	"BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "IntegerLiteral", "StringLiteral", "Identifier", "IdentifierNondigit",
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
	AuroraLexerT__10          = 11
	AuroraLexerT__11          = 12
	AuroraLexerT__12          = 13
	AuroraLexerT__13          = 14
	AuroraLexerT__14          = 15
	AuroraLexerT__15          = 16
	AuroraLexerT__16          = 17
	AuroraLexerT__17          = 18
	AuroraLexerIntegerLiteral = 19
	AuroraLexerStringLiteral  = 20
	AuroraLexerIdentifier     = 21
	AuroraLexerWhitespace     = 22
	AuroraLexerNewline        = 23
	AuroraLexerBlockComment   = 24
	AuroraLexerLineComment    = 25
)
