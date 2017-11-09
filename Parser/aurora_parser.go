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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 298,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 5, 2, 74, 10, 2, 3, 2, 5, 2, 77,
	10, 2, 3, 3, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 3, 4, 3, 4,
	5, 4, 88, 10, 4, 3, 5, 3, 5, 7, 5, 92, 10, 5, 12, 5, 14, 5, 95, 11, 5,
	3, 6, 3, 6, 5, 6, 99, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8,
	107, 10, 8, 3, 8, 5, 8, 110, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 116,
	10, 8, 3, 8, 5, 8, 119, 10, 8, 5, 8, 121, 10, 8, 3, 8, 3, 8, 5, 8, 125,
	10, 8, 3, 9, 3, 9, 5, 9, 129, 10, 9, 3, 9, 5, 9, 132, 10, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 144, 10, 10,
	3, 10, 5, 10, 147, 10, 10, 3, 10, 5, 10, 150, 10, 10, 3, 11, 3, 11, 3,
	11, 7, 11, 155, 10, 11, 12, 11, 14, 11, 158, 11, 11, 3, 12, 6, 12, 161,
	10, 12, 13, 12, 14, 12, 162, 3, 13, 3, 13, 5, 13, 167, 10, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 172, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 182, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	5, 21, 200, 10, 21, 7, 21, 202, 10, 21, 12, 21, 14, 21, 205, 11, 21, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 212, 10, 23, 3, 23, 3, 23, 7, 23,
	216, 10, 23, 12, 23, 14, 23, 219, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 5, 25, 227, 10, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 7, 27, 237, 10, 27, 12, 27, 14, 27, 240, 11, 27, 3, 28, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 250, 10, 29, 12, 29,
	14, 29, 253, 11, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 7, 31, 263, 10, 31, 12, 31, 14, 31, 266, 11, 31, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 276, 10, 33, 12, 33, 14, 33,
	279, 11, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 288,
	10, 35, 3, 35, 3, 35, 3, 35, 5, 35, 293, 10, 35, 3, 36, 3, 36, 3, 36, 3,
	36, 2, 7, 40, 52, 56, 60, 64, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 2, 6, 4, 3, 18, 18, 31, 31, 5, 2, 15, 15, 17, 17,
	20, 22, 4, 2, 17, 17, 23, 24, 3, 2, 20, 21, 2, 297, 2, 73, 3, 2, 2, 2,
	4, 78, 3, 2, 2, 2, 6, 87, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10, 96, 3, 2,
	2, 2, 12, 102, 3, 2, 2, 2, 14, 104, 3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18,
	149, 3, 2, 2, 2, 20, 151, 3, 2, 2, 2, 22, 160, 3, 2, 2, 2, 24, 171, 3,
	2, 2, 2, 26, 173, 3, 2, 2, 2, 28, 175, 3, 2, 2, 2, 30, 181, 3, 2, 2, 2,
	32, 183, 3, 2, 2, 2, 34, 187, 3, 2, 2, 2, 36, 189, 3, 2, 2, 2, 38, 191,
	3, 2, 2, 2, 40, 193, 3, 2, 2, 2, 42, 206, 3, 2, 2, 2, 44, 209, 3, 2, 2,
	2, 46, 222, 3, 2, 2, 2, 48, 226, 3, 2, 2, 2, 50, 228, 3, 2, 2, 2, 52, 231,
	3, 2, 2, 2, 54, 241, 3, 2, 2, 2, 56, 244, 3, 2, 2, 2, 58, 254, 3, 2, 2,
	2, 60, 257, 3, 2, 2, 2, 62, 267, 3, 2, 2, 2, 64, 270, 3, 2, 2, 2, 66, 280,
	3, 2, 2, 2, 68, 292, 3, 2, 2, 2, 70, 294, 3, 2, 2, 2, 72, 74, 5, 4, 3,
	2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77,
	7, 2, 2, 3, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 3, 3, 2, 2, 2,
	78, 82, 5, 6, 4, 2, 79, 81, 5, 6, 4, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3,
	2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84,
	82, 3, 2, 2, 2, 85, 88, 5, 12, 7, 2, 86, 88, 5, 16, 9, 2, 87, 85, 3, 2,
	2, 2, 87, 86, 3, 2, 2, 2, 88, 7, 3, 2, 2, 2, 89, 93, 5, 24, 13, 2, 90,
	92, 5, 24, 13, 2, 91, 90, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2,
	2, 2, 93, 94, 3, 2, 2, 2, 94, 9, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 98,
	7, 3, 2, 2, 97, 99, 5, 8, 5, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2,
	99, 100, 3, 2, 2, 2, 100, 101, 7, 4, 2, 2, 101, 11, 3, 2, 2, 2, 102, 103,
	5, 14, 8, 2, 103, 13, 3, 2, 2, 2, 104, 106, 7, 5, 2, 2, 105, 107, 7, 6,
	2, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2,
	108, 110, 7, 7, 2, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 112, 7, 29, 2, 2, 112, 120, 7, 8, 2, 2, 113, 115,
	7, 29, 2, 2, 114, 116, 5, 22, 12, 2, 115, 114, 3, 2, 2, 2, 115, 116, 3,
	2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 119, 7, 9, 2, 2, 118, 117, 3, 2, 2,
	2, 118, 119, 3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 113, 3, 2, 2, 2, 120,
	121, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 123, 7, 10, 2, 2, 123, 125,
	5, 28, 15, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 15, 3, 2,
	2, 2, 126, 128, 7, 11, 2, 2, 127, 129, 7, 6, 2, 2, 128, 127, 3, 2, 2, 2,
	128, 129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 132, 7, 12, 2, 2, 131,
	130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134,
	7, 29, 2, 2, 134, 135, 7, 13, 2, 2, 135, 136, 5, 20, 11, 2, 136, 137, 7,
	14, 2, 2, 137, 138, 5, 10, 6, 2, 138, 17, 3, 2, 2, 2, 139, 140, 7, 29,
	2, 2, 140, 141, 7, 8, 2, 2, 141, 143, 7, 29, 2, 2, 142, 144, 5, 22, 12,
	2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2, 145,
	147, 7, 9, 2, 2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 150,
	3, 2, 2, 2, 148, 150, 7, 15, 2, 2, 149, 139, 3, 2, 2, 2, 149, 148, 3, 2,
	2, 2, 150, 19, 3, 2, 2, 2, 151, 156, 5, 18, 10, 2, 152, 153, 7, 16, 2,
	2, 153, 155, 5, 18, 10, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2,
	156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 21, 3, 2, 2, 2, 158, 156,
	3, 2, 2, 2, 159, 161, 7, 17, 2, 2, 160, 159, 3, 2, 2, 2, 161, 162, 3, 2,
	2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 23, 3, 2, 2, 2,
	164, 167, 5, 14, 8, 2, 165, 167, 5, 26, 14, 2, 166, 164, 3, 2, 2, 2, 166,
	165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 9, 2, 2, 2, 169, 172,
	3, 2, 2, 2, 170, 172, 5, 10, 6, 2, 171, 166, 3, 2, 2, 2, 171, 170, 3, 2,
	2, 2, 172, 25, 3, 2, 2, 2, 173, 174, 5, 28, 15, 2, 174, 27, 3, 2, 2, 2,
	175, 176, 5, 68, 35, 2, 176, 29, 3, 2, 2, 2, 177, 182, 5, 32, 17, 2, 178,
	182, 5, 34, 18, 2, 179, 182, 5, 38, 20, 2, 180, 182, 5, 36, 19, 2, 181,
	177, 3, 2, 2, 2, 181, 178, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 180,
	3, 2, 2, 2, 182, 31, 3, 2, 2, 2, 183, 184, 7, 13, 2, 2, 184, 185, 5, 28,
	15, 2, 185, 186, 7, 14, 2, 2, 186, 33, 3, 2, 2, 2, 187, 188, 7, 28, 2,
	2, 188, 35, 3, 2, 2, 2, 189, 190, 7, 29, 2, 2, 190, 37, 3, 2, 2, 2, 191,
	192, 7, 27, 2, 2, 192, 39, 3, 2, 2, 2, 193, 194, 8, 21, 1, 2, 194, 195,
	5, 30, 16, 2, 195, 203, 3, 2, 2, 2, 196, 199, 12, 4, 2, 2, 197, 200, 5,
	42, 22, 2, 198, 200, 5, 44, 23, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2,
	2, 2, 200, 202, 3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2,
	203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 41, 3, 2, 2, 2, 205, 203,
	3, 2, 2, 2, 206, 207, 7, 19, 2, 2, 207, 208, 7, 29, 2, 2, 208, 43, 3, 2,
	2, 2, 209, 211, 7, 13, 2, 2, 210, 212, 5, 46, 24, 2, 211, 210, 3, 2, 2,
	2, 211, 212, 3, 2, 2, 2, 212, 217, 3, 2, 2, 2, 213, 214, 7, 16, 2, 2, 214,
	216, 5, 46, 24, 2, 215, 213, 3, 2, 2, 2, 216, 219, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 220, 3, 2, 2, 2, 219, 217, 3, 2,
	2, 2, 220, 221, 7, 14, 2, 2, 221, 45, 3, 2, 2, 2, 222, 223, 5, 28, 15,
	2, 223, 47, 3, 2, 2, 2, 224, 227, 5, 50, 26, 2, 225, 227, 5, 40, 21, 2,
	226, 224, 3, 2, 2, 2, 226, 225, 3, 2, 2, 2, 227, 49, 3, 2, 2, 2, 228, 229,
	9, 3, 2, 2, 229, 230, 5, 28, 15, 2, 230, 51, 3, 2, 2, 2, 231, 232, 8, 27,
	1, 2, 232, 233, 5, 48, 25, 2, 233, 238, 3, 2, 2, 2, 234, 235, 12, 4, 2,
	2, 235, 237, 5, 54, 28, 2, 236, 234, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2,
	238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 53, 3, 2, 2, 2, 240, 238,
	3, 2, 2, 2, 241, 242, 9, 4, 2, 2, 242, 243, 5, 52, 27, 2, 243, 55, 3, 2,
	2, 2, 244, 245, 8, 29, 1, 2, 245, 246, 5, 52, 27, 2, 246, 251, 3, 2, 2,
	2, 247, 248, 12, 4, 2, 2, 248, 250, 5, 58, 30, 2, 249, 247, 3, 2, 2, 2,
	250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252,
	57, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 255, 9, 5, 2, 2, 255, 256, 5,
	56, 29, 2, 256, 59, 3, 2, 2, 2, 257, 258, 8, 31, 1, 2, 258, 259, 5, 56,
	29, 2, 259, 264, 3, 2, 2, 2, 260, 261, 12, 4, 2, 2, 261, 263, 5, 62, 32,
	2, 262, 260, 3, 2, 2, 2, 263, 266, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264,
	265, 3, 2, 2, 2, 265, 61, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 267, 268, 7,
	25, 2, 2, 268, 269, 5, 60, 31, 2, 269, 63, 3, 2, 2, 2, 270, 271, 8, 33,
	1, 2, 271, 272, 5, 60, 31, 2, 272, 277, 3, 2, 2, 2, 273, 274, 12, 4, 2,
	2, 274, 276, 5, 66, 34, 2, 275, 273, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2,
	277, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 65, 3, 2, 2, 2, 279, 277,
	3, 2, 2, 2, 280, 281, 7, 26, 2, 2, 281, 282, 5, 64, 33, 2, 282, 67, 3,
	2, 2, 2, 283, 284, 5, 40, 21, 2, 284, 285, 5, 42, 22, 2, 285, 288, 3, 2,
	2, 2, 286, 288, 5, 36, 19, 2, 287, 283, 3, 2, 2, 2, 287, 286, 3, 2, 2,
	2, 288, 289, 3, 2, 2, 2, 289, 290, 5, 70, 36, 2, 290, 293, 3, 2, 2, 2,
	291, 293, 5, 64, 33, 2, 292, 287, 3, 2, 2, 2, 292, 291, 3, 2, 2, 2, 293,
	69, 3, 2, 2, 2, 294, 295, 7, 10, 2, 2, 295, 296, 5, 28, 15, 2, 296, 71,
	3, 2, 2, 2, 35, 73, 76, 82, 87, 93, 98, 106, 109, 115, 118, 120, 124, 128,
	131, 143, 146, 149, 156, 162, 166, 171, 181, 199, 203, 211, 217, 226, 238,
	251, 264, 277, 287, 292,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'var'", "'export'", "'optional'", "':'", "'^'", "'='",
	"'func'", "'native'", "'('", "')'", "'&'", "','", "'*'", "';'", "'.'",
	"'-'", "'+'", "'!'", "'/'", "'%'", "'&&'", "'||'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "IntegerLiteral", "StringLiteral", "Identifier",
	"Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "declarationList", "declaration", "statementList", "codeBlock",
	"variableDeclaration", "variableStmt", "functionDeclaration", "functionParameter",
	"functionParameterList", "pointer", "stmt", "expressionStmt", "expr", "expr0",
	"brackets", "stringImmidiate", "identifierImmidiate", "integerImmidiate",
	"expr1", "memberAccess", "functionCall", "functionCallParam", "expr2",
	"lhsOperator", "expr3", "mulDivMod", "expr4", "addSub", "expr5", "logicalAnd",
	"expr6", "logicalOr", "expr7", "assign",
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
	AuroraParserT__10          = 11
	AuroraParserT__11          = 12
	AuroraParserT__12          = 13
	AuroraParserT__13          = 14
	AuroraParserT__14          = 15
	AuroraParserT__15          = 16
	AuroraParserT__16          = 17
	AuroraParserT__17          = 18
	AuroraParserT__18          = 19
	AuroraParserT__19          = 20
	AuroraParserT__20          = 21
	AuroraParserT__21          = 22
	AuroraParserT__22          = 23
	AuroraParserT__23          = 24
	AuroraParserIntegerLiteral = 25
	AuroraParserStringLiteral  = 26
	AuroraParserIdentifier     = 27
	AuroraParserWhitespace     = 28
	AuroraParserNewline        = 29
	AuroraParserBlockComment   = 30
	AuroraParserLineComment    = 31
)

// AuroraParser rules.
const (
	AuroraParserRULE_program               = 0
	AuroraParserRULE_declarationList       = 1
	AuroraParserRULE_declaration           = 2
	AuroraParserRULE_statementList         = 3
	AuroraParserRULE_codeBlock             = 4
	AuroraParserRULE_variableDeclaration   = 5
	AuroraParserRULE_variableStmt          = 6
	AuroraParserRULE_functionDeclaration   = 7
	AuroraParserRULE_functionParameter     = 8
	AuroraParserRULE_functionParameterList = 9
	AuroraParserRULE_pointer               = 10
	AuroraParserRULE_stmt                  = 11
	AuroraParserRULE_expressionStmt        = 12
	AuroraParserRULE_expr                  = 13
	AuroraParserRULE_expr0                 = 14
	AuroraParserRULE_brackets              = 15
	AuroraParserRULE_stringImmidiate       = 16
	AuroraParserRULE_identifierImmidiate   = 17
	AuroraParserRULE_integerImmidiate      = 18
	AuroraParserRULE_expr1                 = 19
	AuroraParserRULE_memberAccess          = 20
	AuroraParserRULE_functionCall          = 21
	AuroraParserRULE_functionCallParam     = 22
	AuroraParserRULE_expr2                 = 23
	AuroraParserRULE_lhsOperator           = 24
	AuroraParserRULE_expr3                 = 25
	AuroraParserRULE_mulDivMod             = 26
	AuroraParserRULE_expr4                 = 27
	AuroraParserRULE_addSub                = 28
	AuroraParserRULE_expr5                 = 29
	AuroraParserRULE_logicalAnd            = 30
	AuroraParserRULE_expr6                 = 31
	AuroraParserRULE_logicalOr             = 32
	AuroraParserRULE_expr7                 = 33
	AuroraParserRULE_assign                = 34
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

func (s *ProgramContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__2 || _la == AuroraParserT__8 {
		{
			p.SetState(70)
			p.DeclarationList()
		}

	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(73)
			p.Match(AuroraParserEOF)
		}

	}

	return localctx
}

// IDeclarationListContext is an interface to support dynamic dispatch.
type IDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationListContext differentiates from other interfaces.
	IsDeclarationListContext()
}

type DeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationListContext() *DeclarationListContext {
	var p = new(DeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_declarationList
	return p
}

func (*DeclarationListContext) IsDeclarationListContext() {}

func NewDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationListContext {
	var p = new(DeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_declarationList

	return p
}

func (s *DeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationListContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *DeclarationListContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterDeclarationList(s)
	}
}

func (s *DeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitDeclarationList(s)
	}
}

func (p *AuroraParser) DeclarationList() (localctx IDeclarationListContext) {
	localctx = NewDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AuroraParserRULE_declarationList)
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
		p.SetState(76)
		p.Declaration()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__2 || _la == AuroraParserT__8 {
		{
			p.SetState(77)
			p.Declaration()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *AuroraParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AuroraParserRULE_declaration)

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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.VariableDeclaration()
		}

	case AuroraParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.FunctionDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *StatementListContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *AuroraParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AuroraParserRULE_statementList)
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
		p.SetState(87)
		p.Stmt()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__0)|(1<<AuroraParserT__2)|(1<<AuroraParserT__10)|(1<<AuroraParserT__12)|(1<<AuroraParserT__14)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(88)
			p.Stmt()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_codeBlock
	return p
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (p *AuroraParser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AuroraParserRULE_codeBlock)
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
		p.SetState(94)
		p.Match(AuroraParserT__0)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__0)|(1<<AuroraParserT__2)|(1<<AuroraParserT__10)|(1<<AuroraParserT__12)|(1<<AuroraParserT__14)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(95)
			p.StatementList()
		}

	}
	{
		p.SetState(98)
		p.Match(AuroraParserT__1)
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableStmt() IVariableStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStmtContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *AuroraParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AuroraParserRULE_variableDeclaration)

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
		p.SetState(100)
		p.VariableStmt()
	}

	return localctx
}

// IVariableStmtContext is an interface to support dynamic dispatch.
type IVariableStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExport returns the export token.
	GetExport() antlr.Token

	// GetOptional returns the optional token.
	GetOptional() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// GetVariableType returns the variableType token.
	GetVariableType() antlr.Token

	// GetGcptr returns the gcptr token.
	GetGcptr() antlr.Token

	// SetExport sets the export token.
	SetExport(antlr.Token)

	// SetOptional sets the optional token.
	SetOptional(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetVariableType sets the variableType token.
	SetVariableType(antlr.Token)

	// SetGcptr sets the gcptr token.
	SetGcptr(antlr.Token)

	// GetPtr returns the ptr rule contexts.
	GetPtr() IPointerContext

	// SetPtr sets the ptr rule contexts.
	SetPtr(IPointerContext)

	// IsVariableStmtContext differentiates from other interfaces.
	IsVariableStmtContext()
}

type VariableStmtContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	export       antlr.Token
	optional     antlr.Token
	name         antlr.Token
	variableType antlr.Token
	ptr          IPointerContext
	gcptr        antlr.Token
}

func NewEmptyVariableStmtContext() *VariableStmtContext {
	var p = new(VariableStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_variableStmt
	return p
}

func (*VariableStmtContext) IsVariableStmtContext() {}

func NewVariableStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStmtContext {
	var p = new(VariableStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_variableStmt

	return p
}

func (s *VariableStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStmtContext) GetExport() antlr.Token { return s.export }

func (s *VariableStmtContext) GetOptional() antlr.Token { return s.optional }

func (s *VariableStmtContext) GetName() antlr.Token { return s.name }

func (s *VariableStmtContext) GetVariableType() antlr.Token { return s.variableType }

func (s *VariableStmtContext) GetGcptr() antlr.Token { return s.gcptr }

func (s *VariableStmtContext) SetExport(v antlr.Token) { s.export = v }

func (s *VariableStmtContext) SetOptional(v antlr.Token) { s.optional = v }

func (s *VariableStmtContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableStmtContext) SetVariableType(v antlr.Token) { s.variableType = v }

func (s *VariableStmtContext) SetGcptr(v antlr.Token) { s.gcptr = v }

func (s *VariableStmtContext) GetPtr() IPointerContext { return s.ptr }

func (s *VariableStmtContext) SetPtr(v IPointerContext) { s.ptr = v }

func (s *VariableStmtContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(AuroraParserIdentifier)
}

func (s *VariableStmtContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(AuroraParserIdentifier, i)
}

func (s *VariableStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariableStmtContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *VariableStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterVariableStmt(s)
	}
}

func (s *VariableStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitVariableStmt(s)
	}
}

func (p *AuroraParser) VariableStmt() (localctx IVariableStmtContext) {
	localctx = NewVariableStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AuroraParserRULE_variableStmt)
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
		p.SetState(102)
		p.Match(AuroraParserT__2)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__3 {
		{
			p.SetState(103)

			var _m = p.Match(AuroraParserT__3)

			localctx.(*VariableStmtContext).export = _m
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__4 {
		{
			p.SetState(106)

			var _m = p.Match(AuroraParserT__4)

			localctx.(*VariableStmtContext).optional = _m
		}

	}
	{
		p.SetState(109)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*VariableStmtContext).name = _m
	}
	{
		p.SetState(110)
		p.Match(AuroraParserT__5)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserIdentifier {
		{
			p.SetState(111)

			var _m = p.Match(AuroraParserIdentifier)

			localctx.(*VariableStmtContext).variableType = _m
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__14 {
			{
				p.SetState(112)

				var _x = p.Pointer()

				localctx.(*VariableStmtContext).ptr = _x
			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__6 {
			{
				p.SetState(115)

				var _m = p.Match(AuroraParserT__6)

				localctx.(*VariableStmtContext).gcptr = _m
			}

		}

	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__7 {
		{
			p.SetState(120)
			p.Match(AuroraParserT__7)
		}
		{
			p.SetState(121)
			p.Expr()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExport returns the export token.
	GetExport() antlr.Token

	// GetNative returns the native token.
	GetNative() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetExport sets the export token.
	SetExport(antlr.Token)

	// SetNative sets the native token.
	SetNative(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetCode returns the code rule contexts.
	GetCode() ICodeBlockContext

	// SetCode sets the code rule contexts.
	SetCode(ICodeBlockContext)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	export antlr.Token
	native antlr.Token
	name   antlr.Token
	code   ICodeBlockContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetExport() antlr.Token { return s.export }

func (s *FunctionDeclarationContext) GetNative() antlr.Token { return s.native }

func (s *FunctionDeclarationContext) GetName() antlr.Token { return s.name }

func (s *FunctionDeclarationContext) SetExport(v antlr.Token) { s.export = v }

func (s *FunctionDeclarationContext) SetNative(v antlr.Token) { s.native = v }

func (s *FunctionDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDeclarationContext) GetCode() ICodeBlockContext { return s.code }

func (s *FunctionDeclarationContext) SetCode(v ICodeBlockContext) { s.code = v }

func (s *FunctionDeclarationContext) FunctionParameterList() IFunctionParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterListContext)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AuroraParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *AuroraParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AuroraParserRULE_functionDeclaration)
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
		p.SetState(124)
		p.Match(AuroraParserT__8)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__3 {
		{
			p.SetState(125)

			var _m = p.Match(AuroraParserT__3)

			localctx.(*FunctionDeclarationContext).export = _m
		}

	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__9 {
		{
			p.SetState(128)

			var _m = p.Match(AuroraParserT__9)

			localctx.(*FunctionDeclarationContext).native = _m
		}

	}
	{
		p.SetState(131)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*FunctionDeclarationContext).name = _m
	}
	{
		p.SetState(132)
		p.Match(AuroraParserT__10)
	}
	{
		p.SetState(133)
		p.FunctionParameterList()
	}
	{
		p.SetState(134)
		p.Match(AuroraParserT__11)
	}
	{
		p.SetState(135)

		var _x = p.CodeBlock()

		localctx.(*FunctionDeclarationContext).code = _x
	}

	return localctx
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParameterName returns the parameterName token.
	GetParameterName() antlr.Token

	// GetParameterType returns the parameterType token.
	GetParameterType() antlr.Token

	// GetGcptr returns the gcptr token.
	GetGcptr() antlr.Token

	// GetRef returns the ref token.
	GetRef() antlr.Token

	// SetParameterName sets the parameterName token.
	SetParameterName(antlr.Token)

	// SetParameterType sets the parameterType token.
	SetParameterType(antlr.Token)

	// SetGcptr sets the gcptr token.
	SetGcptr(antlr.Token)

	// SetRef sets the ref token.
	SetRef(antlr.Token)

	// GetPtr returns the ptr rule contexts.
	GetPtr() IPointerContext

	// SetPtr sets the ptr rule contexts.
	SetPtr(IPointerContext)

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	parameterName antlr.Token
	parameterType antlr.Token
	ptr           IPointerContext
	gcptr         antlr.Token
	ref           antlr.Token
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_functionParameter
	return p
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) GetParameterName() antlr.Token { return s.parameterName }

func (s *FunctionParameterContext) GetParameterType() antlr.Token { return s.parameterType }

func (s *FunctionParameterContext) GetGcptr() antlr.Token { return s.gcptr }

func (s *FunctionParameterContext) GetRef() antlr.Token { return s.ref }

func (s *FunctionParameterContext) SetParameterName(v antlr.Token) { s.parameterName = v }

func (s *FunctionParameterContext) SetParameterType(v antlr.Token) { s.parameterType = v }

func (s *FunctionParameterContext) SetGcptr(v antlr.Token) { s.gcptr = v }

func (s *FunctionParameterContext) SetRef(v antlr.Token) { s.ref = v }

func (s *FunctionParameterContext) GetPtr() IPointerContext { return s.ptr }

func (s *FunctionParameterContext) SetPtr(v IPointerContext) { s.ptr = v }

func (s *FunctionParameterContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(AuroraParserIdentifier)
}

func (s *FunctionParameterContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(AuroraParserIdentifier, i)
}

func (s *FunctionParameterContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (p *AuroraParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AuroraParserRULE_functionParameter)
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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)

			var _m = p.Match(AuroraParserIdentifier)

			localctx.(*FunctionParameterContext).parameterName = _m
		}
		{
			p.SetState(138)
			p.Match(AuroraParserT__5)
		}
		{
			p.SetState(139)

			var _m = p.Match(AuroraParserIdentifier)

			localctx.(*FunctionParameterContext).parameterType = _m
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__14 {
			{
				p.SetState(140)

				var _x = p.Pointer()

				localctx.(*FunctionParameterContext).ptr = _x
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__6 {
			{
				p.SetState(143)

				var _m = p.Match(AuroraParserT__6)

				localctx.(*FunctionParameterContext).gcptr = _m
			}

		}

	case AuroraParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)

			var _m = p.Match(AuroraParserT__12)

			localctx.(*FunctionParameterContext).ref = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionParameterListContext is an interface to support dynamic dispatch.
type IFunctionParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterListContext differentiates from other interfaces.
	IsFunctionParameterListContext()
}

type FunctionParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterListContext() *FunctionParameterListContext {
	var p = new(FunctionParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_functionParameterList
	return p
}

func (*FunctionParameterListContext) IsFunctionParameterListContext() {}

func NewFunctionParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterListContext {
	var p = new(FunctionParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_functionParameterList

	return p
}

func (s *FunctionParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterListContext) AllFunctionParameter() []IFunctionParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem())
	var tst = make([]IFunctionParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionParameterContext)
		}
	}

	return tst
}

func (s *FunctionParameterListContext) FunctionParameter(i int) IFunctionParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterFunctionParameterList(s)
	}
}

func (s *FunctionParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitFunctionParameterList(s)
	}
}

func (p *AuroraParser) FunctionParameterList() (localctx IFunctionParameterListContext) {
	localctx = NewFunctionParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AuroraParserRULE_functionParameterList)
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
		p.SetState(149)
		p.FunctionParameter()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__13 {
		{
			p.SetState(150)
			p.Match(AuroraParserT__13)
		}
		{
			p.SetState(151)
			p.FunctionParameter()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_pointer
	return p
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }
func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterPointer(s)
	}
}

func (s *PointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitPointer(s)
	}
}

func (p *AuroraParser) Pointer() (localctx IPointerContext) {
	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AuroraParserRULE_pointer)
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
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AuroraParserT__14 {
		{
			p.SetState(157)
			p.Match(AuroraParserT__14)
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *StmtContext) Newline() antlr.TerminalNode {
	return s.GetToken(AuroraParserNewline, 0)
}

func (s *StmtContext) EOF() antlr.TerminalNode {
	return s.GetToken(AuroraParserEOF, 0)
}

func (s *StmtContext) VariableStmt() IVariableStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStmtContext)
}

func (s *StmtContext) ExpressionStmt() IExpressionStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *AuroraParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AuroraParserRULE_stmt)
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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__2, AuroraParserT__10, AuroraParserT__12, AuroraParserT__14, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AuroraParserT__2:
			{
				p.SetState(162)
				p.VariableStmt()
			}

		case AuroraParserT__10, AuroraParserT__12, AuroraParserT__14, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
			{
				p.SetState(163)
				p.ExpressionStmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(166)
		_la = p.GetTokenStream().LA(1)

		if !(((_la - -1)&-(0x1f+1)) == 0 && ((1<<uint((_la - -1)))&((1<<(AuroraParserEOF - -1))|(1<<(AuroraParserT__15 - -1))|(1<<(AuroraParserNewline - -1)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case AuroraParserT__0:
		{
			p.SetState(168)
			p.CodeBlock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expressionStmt
	return p
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpressionStmt(s)
	}
}

func (p *AuroraParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AuroraParserRULE_expressionStmt)

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
		p.SetState(171)
		p.Expr()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Expr7() IExpr7Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr7Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr7Context)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *AuroraParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AuroraParserRULE_expr)

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
		p.SetState(173)
		p.Expr7()
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
	p.EnterRule(localctx, 28, AuroraParserRULE_expr0)

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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__10:
		{
			p.SetState(175)
			p.Brackets()
		}

	case AuroraParserStringLiteral:
		{
			p.SetState(176)
			p.StringImmidiate()
		}

	case AuroraParserIntegerLiteral:
		{
			p.SetState(177)
			p.IntegerImmidiate()
		}

	case AuroraParserIdentifier:
		{
			p.SetState(178)
			p.IdentifierImmidiate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBracketsContext is an interface to support dynamic dispatch.
type IBracketsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParamExpr returns the paramExpr rule contexts.
	GetParamExpr() IExprContext

	// SetParamExpr sets the paramExpr rule contexts.
	SetParamExpr(IExprContext)

	// IsBracketsContext differentiates from other interfaces.
	IsBracketsContext()
}

type BracketsContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	paramExpr IExprContext
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

func (s *BracketsContext) GetParamExpr() IExprContext { return s.paramExpr }

func (s *BracketsContext) SetParamExpr(v IExprContext) { s.paramExpr = v }

func (s *BracketsContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 30, AuroraParserRULE_brackets)

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
		p.SetState(181)
		p.Match(AuroraParserT__10)
	}
	{
		p.SetState(182)

		var _x = p.Expr()

		localctx.(*BracketsContext).paramExpr = _x
	}
	{
		p.SetState(183)
		p.Match(AuroraParserT__11)
	}

	return localctx
}

// IStringImmidiateContext is an interface to support dynamic dispatch.
type IStringImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStringVal returns the stringVal token.
	GetStringVal() antlr.Token

	// SetStringVal sets the stringVal token.
	SetStringVal(antlr.Token)

	// IsStringImmidiateContext differentiates from other interfaces.
	IsStringImmidiateContext()
}

type StringImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	stringVal antlr.Token
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

func (s *StringImmidiateContext) GetStringVal() antlr.Token { return s.stringVal }

func (s *StringImmidiateContext) SetStringVal(v antlr.Token) { s.stringVal = v }

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
	p.EnterRule(localctx, 32, AuroraParserRULE_stringImmidiate)

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
		p.SetState(185)

		var _m = p.Match(AuroraParserStringLiteral)

		localctx.(*StringImmidiateContext).stringVal = _m
	}

	return localctx
}

// IIdentifierImmidiateContext is an interface to support dynamic dispatch.
type IIdentifierImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsIdentifierImmidiateContext differentiates from other interfaces.
	IsIdentifierImmidiateContext()
}

type IdentifierImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
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

func (s *IdentifierImmidiateContext) GetName() antlr.Token { return s.name }

func (s *IdentifierImmidiateContext) SetName(v antlr.Token) { s.name = v }

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
	p.EnterRule(localctx, 34, AuroraParserRULE_identifierImmidiate)

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
		p.SetState(187)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*IdentifierImmidiateContext).name = _m
	}

	return localctx
}

// IIntegerImmidiateContext is an interface to support dynamic dispatch.
type IIntegerImmidiateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNumberVal returns the numberVal token.
	GetNumberVal() antlr.Token

	// SetNumberVal sets the numberVal token.
	SetNumberVal(antlr.Token)

	// IsIntegerImmidiateContext differentiates from other interfaces.
	IsIntegerImmidiateContext()
}

type IntegerImmidiateContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	numberVal antlr.Token
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

func (s *IntegerImmidiateContext) GetNumberVal() antlr.Token { return s.numberVal }

func (s *IntegerImmidiateContext) SetNumberVal(v antlr.Token) { s.numberVal = v }

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
	p.EnterRule(localctx, 36, AuroraParserRULE_integerImmidiate)

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
		p.SetState(189)

		var _m = p.Match(AuroraParserIntegerLiteral)

		localctx.(*IntegerImmidiateContext).numberVal = _m
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

func (s *Expr1Context) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *Expr1Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *Expr1Context) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *Expr1Context) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
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
	return p.expr1(0)
}

func (p *AuroraParser) expr1(_p int) (localctx IExpr1Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr1Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr1Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, AuroraParserRULE_expr1, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Expr0()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr1Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr1)
			p.SetState(194)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(197)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AuroraParserT__16:
				{
					p.SetState(195)
					p.MemberAccess()
				}

			case AuroraParserT__10:
				{
					p.SetState(196)
					p.FunctionCall()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMember returns the member token.
	GetMember() antlr.Token

	// SetMember sets the member token.
	SetMember(antlr.Token)

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	member antlr.Token
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

func (s *MemberAccessContext) GetMember() antlr.Token { return s.member }

func (s *MemberAccessContext) SetMember(v antlr.Token) { s.member = v }

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
	p.EnterRule(localctx, 40, AuroraParserRULE_memberAccess)

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
		p.SetState(204)
		p.Match(AuroraParserT__16)
	}
	{
		p.SetState(205)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*MemberAccessContext).member = _m
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
	p.EnterRule(localctx, 42, AuroraParserRULE_functionCall)
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
		p.SetState(207)
		p.Match(AuroraParserT__10)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__10)|(1<<AuroraParserT__12)|(1<<AuroraParserT__14)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(208)
			p.FunctionCallParam()
		}

	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__13 {
		{
			p.SetState(211)
			p.Match(AuroraParserT__13)
		}
		{
			p.SetState(212)
			p.FunctionCallParam()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(AuroraParserT__11)
	}

	return localctx
}

// IFunctionCallParamContext is an interface to support dynamic dispatch.
type IFunctionCallParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParamExpr returns the paramExpr rule contexts.
	GetParamExpr() IExprContext

	// SetParamExpr sets the paramExpr rule contexts.
	SetParamExpr(IExprContext)

	// IsFunctionCallParamContext differentiates from other interfaces.
	IsFunctionCallParamContext()
}

type FunctionCallParamContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	paramExpr IExprContext
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

func (s *FunctionCallParamContext) GetParamExpr() IExprContext { return s.paramExpr }

func (s *FunctionCallParamContext) SetParamExpr(v IExprContext) { s.paramExpr = v }

func (s *FunctionCallParamContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 44, AuroraParserRULE_functionCallParam)

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
		p.SetState(220)

		var _x = p.Expr()

		localctx.(*FunctionCallParamContext).paramExpr = _x
	}

	return localctx
}

// IExpr2Context is an interface to support dynamic dispatch.
type IExpr2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr2Context differentiates from other interfaces.
	IsExpr2Context()
}

type Expr2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Expr2Context) LhsOperator() ILhsOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILhsOperatorContext)
}

func (s *Expr2Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
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
	p.EnterRule(localctx, 46, AuroraParserRULE_expr2)

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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__12, AuroraParserT__14, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.LhsOperator()
		}

	case AuroraParserT__10, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.expr1(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILhsOperatorContext is an interface to support dynamic dispatch.
type ILhsOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOn returns the on rule contexts.
	GetOn() IExprContext

	// SetOn sets the on rule contexts.
	SetOn(IExprContext)

	// IsLhsOperatorContext differentiates from other interfaces.
	IsLhsOperatorContext()
}

type LhsOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	on     IExprContext
}

func NewEmptyLhsOperatorContext() *LhsOperatorContext {
	var p = new(LhsOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_lhsOperator
	return p
}

func (*LhsOperatorContext) IsLhsOperatorContext() {}

func NewLhsOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsOperatorContext {
	var p = new(LhsOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_lhsOperator

	return p
}

func (s *LhsOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsOperatorContext) GetOp() antlr.Token { return s.op }

func (s *LhsOperatorContext) SetOp(v antlr.Token) { s.op = v }

func (s *LhsOperatorContext) GetOn() IExprContext { return s.on }

func (s *LhsOperatorContext) SetOn(v IExprContext) { s.on = v }

func (s *LhsOperatorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LhsOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterLhsOperator(s)
	}
}

func (s *LhsOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitLhsOperator(s)
	}
}

func (p *AuroraParser) LhsOperator() (localctx ILhsOperatorContext) {
	localctx = NewLhsOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AuroraParserRULE_lhsOperator)
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
	p.SetState(226)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*LhsOperatorContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__12)|(1<<AuroraParserT__14)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*LhsOperatorContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(227)

		var _x = p.Expr()

		localctx.(*LhsOperatorContext).on = _x
	}

	return localctx
}

// IExpr3Context is an interface to support dynamic dispatch.
type IExpr3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpr3Context

	// SetLeft sets the left rule contexts.
	SetLeft(IExpr3Context)

	// IsExpr3Context differentiates from other interfaces.
	IsExpr3Context()
}

type Expr3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr3Context
}

func NewEmptyExpr3Context() *Expr3Context {
	var p = new(Expr3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr3
	return p
}

func (*Expr3Context) IsExpr3Context() {}

func NewExpr3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr3Context {
	var p = new(Expr3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr3

	return p
}

func (s *Expr3Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr3Context) GetLeft() IExpr3Context { return s.left }

func (s *Expr3Context) SetLeft(v IExpr3Context) { s.left = v }

func (s *Expr3Context) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *Expr3Context) MulDivMod() IMulDivModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulDivModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulDivModContext)
}

func (s *Expr3Context) Expr3() IExpr3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *Expr3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr3(s)
	}
}

func (s *Expr3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr3(s)
	}
}

func (p *AuroraParser) Expr3() (localctx IExpr3Context) {
	return p.expr3(0)
}

func (p *AuroraParser) expr3(_p int) (localctx IExpr3Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr3Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr3Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, AuroraParserRULE_expr3, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Expr2()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr3Context(p, _parentctx, _parentState)
			localctx.(*Expr3Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr3)
			p.SetState(232)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(233)
				p.MulDivMod()
			}

		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IMulDivModContext is an interface to support dynamic dispatch.
type IMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IExpr3Context

	// SetRight sets the right rule contexts.
	SetRight(IExpr3Context)

	// IsMulDivModContext differentiates from other interfaces.
	IsMulDivModContext()
}

type MulDivModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  IExpr3Context
}

func NewEmptyMulDivModContext() *MulDivModContext {
	var p = new(MulDivModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_mulDivMod
	return p
}

func (*MulDivModContext) IsMulDivModContext() {}

func NewMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivModContext {
	var p = new(MulDivModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_mulDivMod

	return p
}

func (s *MulDivModContext) GetParser() antlr.Parser { return s.parser }

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRight() IExpr3Context { return s.right }

func (s *MulDivModContext) SetRight(v IExpr3Context) { s.right = v }

func (s *MulDivModContext) Expr3() IExpr3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (p *AuroraParser) MulDivMod() (localctx IMulDivModContext) {
	localctx = NewMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AuroraParserRULE_mulDivMod)
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
	p.SetState(239)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*MulDivModContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__14)|(1<<AuroraParserT__20)|(1<<AuroraParserT__21))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*MulDivModContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(240)

		var _x = p.expr3(0)

		localctx.(*MulDivModContext).right = _x
	}

	return localctx
}

// IExpr4Context is an interface to support dynamic dispatch.
type IExpr4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpr4Context

	// SetLeft sets the left rule contexts.
	SetLeft(IExpr4Context)

	// IsExpr4Context differentiates from other interfaces.
	IsExpr4Context()
}

type Expr4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr4Context
}

func NewEmptyExpr4Context() *Expr4Context {
	var p = new(Expr4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr4
	return p
}

func (*Expr4Context) IsExpr4Context() {}

func NewExpr4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr4Context {
	var p = new(Expr4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr4

	return p
}

func (s *Expr4Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr4Context) GetLeft() IExpr4Context { return s.left }

func (s *Expr4Context) SetLeft(v IExpr4Context) { s.left = v }

func (s *Expr4Context) Expr3() IExpr3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *Expr4Context) AddSub() IAddSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddSubContext)
}

func (s *Expr4Context) Expr4() IExpr4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr4Context)
}

func (s *Expr4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr4(s)
	}
}

func (s *Expr4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr4(s)
	}
}

func (p *AuroraParser) Expr4() (localctx IExpr4Context) {
	return p.expr4(0)
}

func (p *AuroraParser) expr4(_p int) (localctx IExpr4Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr4Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr4Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, AuroraParserRULE_expr4, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.expr3(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr4Context(p, _parentctx, _parentState)
			localctx.(*Expr4Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr4)
			p.SetState(245)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(246)
				p.AddSub()
			}

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IAddSubContext is an interface to support dynamic dispatch.
type IAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IExpr4Context

	// SetRight sets the right rule contexts.
	SetRight(IExpr4Context)

	// IsAddSubContext differentiates from other interfaces.
	IsAddSubContext()
}

type AddSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  IExpr4Context
}

func NewEmptyAddSubContext() *AddSubContext {
	var p = new(AddSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_addSub
	return p
}

func (*AddSubContext) IsAddSubContext() {}

func NewAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddSubContext {
	var p = new(AddSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_addSub

	return p
}

func (s *AddSubContext) GetParser() antlr.Parser { return s.parser }

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRight() IExpr4Context { return s.right }

func (s *AddSubContext) SetRight(v IExpr4Context) { s.right = v }

func (s *AddSubContext) Expr4() IExpr4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr4Context)
}

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (p *AuroraParser) AddSub() (localctx IAddSubContext) {
	localctx = NewAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AuroraParserRULE_addSub)
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
	p.SetState(252)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*AddSubContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(_la == AuroraParserT__17 || _la == AuroraParserT__18) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*AddSubContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(253)

		var _x = p.expr4(0)

		localctx.(*AddSubContext).right = _x
	}

	return localctx
}

// IExpr5Context is an interface to support dynamic dispatch.
type IExpr5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpr5Context

	// SetLeft sets the left rule contexts.
	SetLeft(IExpr5Context)

	// IsExpr5Context differentiates from other interfaces.
	IsExpr5Context()
}

type Expr5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr5Context
}

func NewEmptyExpr5Context() *Expr5Context {
	var p = new(Expr5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr5
	return p
}

func (*Expr5Context) IsExpr5Context() {}

func NewExpr5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr5Context {
	var p = new(Expr5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr5

	return p
}

func (s *Expr5Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr5Context) GetLeft() IExpr5Context { return s.left }

func (s *Expr5Context) SetLeft(v IExpr5Context) { s.left = v }

func (s *Expr5Context) Expr4() IExpr4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr4Context)
}

func (s *Expr5Context) LogicalAnd() ILogicalAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalAndContext)
}

func (s *Expr5Context) Expr5() IExpr5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr5Context)
}

func (s *Expr5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr5(s)
	}
}

func (s *Expr5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr5(s)
	}
}

func (p *AuroraParser) Expr5() (localctx IExpr5Context) {
	return p.expr5(0)
}

func (p *AuroraParser) expr5(_p int) (localctx IExpr5Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr5Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr5Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, AuroraParserRULE_expr5, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.expr4(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr5Context(p, _parentctx, _parentState)
			localctx.(*Expr5Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr5)
			p.SetState(258)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(259)
				p.LogicalAnd()
			}

		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicalAndContext is an interface to support dynamic dispatch.
type ILogicalAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRight returns the right rule contexts.
	GetRight() IExpr5Context

	// SetRight sets the right rule contexts.
	SetRight(IExpr5Context)

	// IsLogicalAndContext differentiates from other interfaces.
	IsLogicalAndContext()
}

type LogicalAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	right  IExpr5Context
}

func NewEmptyLogicalAndContext() *LogicalAndContext {
	var p = new(LogicalAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_logicalAnd
	return p
}

func (*LogicalAndContext) IsLogicalAndContext() {}

func NewLogicalAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndContext {
	var p = new(LogicalAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_logicalAnd

	return p
}

func (s *LogicalAndContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndContext) GetRight() IExpr5Context { return s.right }

func (s *LogicalAndContext) SetRight(v IExpr5Context) { s.right = v }

func (s *LogicalAndContext) Expr5() IExpr5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr5Context)
}

func (s *LogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterLogicalAnd(s)
	}
}

func (s *LogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitLogicalAnd(s)
	}
}

func (p *AuroraParser) LogicalAnd() (localctx ILogicalAndContext) {
	localctx = NewLogicalAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AuroraParserRULE_logicalAnd)

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
		p.SetState(265)
		p.Match(AuroraParserT__22)
	}
	{
		p.SetState(266)

		var _x = p.expr5(0)

		localctx.(*LogicalAndContext).right = _x
	}

	return localctx
}

// IExpr6Context is an interface to support dynamic dispatch.
type IExpr6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpr6Context

	// SetLeft sets the left rule contexts.
	SetLeft(IExpr6Context)

	// IsExpr6Context differentiates from other interfaces.
	IsExpr6Context()
}

type Expr6Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr6Context
}

func NewEmptyExpr6Context() *Expr6Context {
	var p = new(Expr6Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr6
	return p
}

func (*Expr6Context) IsExpr6Context() {}

func NewExpr6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr6Context {
	var p = new(Expr6Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr6

	return p
}

func (s *Expr6Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr6Context) GetLeft() IExpr6Context { return s.left }

func (s *Expr6Context) SetLeft(v IExpr6Context) { s.left = v }

func (s *Expr6Context) Expr5() IExpr5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr5Context)
}

func (s *Expr6Context) LogicalOr() ILogicalOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOrContext)
}

func (s *Expr6Context) Expr6() IExpr6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr6Context)
}

func (s *Expr6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr6(s)
	}
}

func (s *Expr6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr6(s)
	}
}

func (p *AuroraParser) Expr6() (localctx IExpr6Context) {
	return p.expr6(0)
}

func (p *AuroraParser) expr6(_p int) (localctx IExpr6Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr6Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr6Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, AuroraParserRULE_expr6, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.expr5(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr6Context(p, _parentctx, _parentState)
			localctx.(*Expr6Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr6)
			p.SetState(271)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(272)
				p.LogicalOr()
			}

		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicalOrContext is an interface to support dynamic dispatch.
type ILogicalOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRight returns the right rule contexts.
	GetRight() IExpr6Context

	// SetRight sets the right rule contexts.
	SetRight(IExpr6Context)

	// IsLogicalOrContext differentiates from other interfaces.
	IsLogicalOrContext()
}

type LogicalOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	right  IExpr6Context
}

func NewEmptyLogicalOrContext() *LogicalOrContext {
	var p = new(LogicalOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_logicalOr
	return p
}

func (*LogicalOrContext) IsLogicalOrContext() {}

func NewLogicalOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrContext {
	var p = new(LogicalOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_logicalOr

	return p
}

func (s *LogicalOrContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrContext) GetRight() IExpr6Context { return s.right }

func (s *LogicalOrContext) SetRight(v IExpr6Context) { s.right = v }

func (s *LogicalOrContext) Expr6() IExpr6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr6Context)
}

func (s *LogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterLogicalOr(s)
	}
}

func (s *LogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitLogicalOr(s)
	}
}

func (p *AuroraParser) LogicalOr() (localctx ILogicalOrContext) {
	localctx = NewLogicalOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AuroraParserRULE_logicalOr)

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
		p.SetState(278)
		p.Match(AuroraParserT__23)
	}
	{
		p.SetState(279)

		var _x = p.expr6(0)

		localctx.(*LogicalOrContext).right = _x
	}

	return localctx
}

// IExpr7Context is an interface to support dynamic dispatch.
type IExpr7Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr7Context differentiates from other interfaces.
	IsExpr7Context()
}

type Expr7Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr7Context() *Expr7Context {
	var p = new(Expr7Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_expr7
	return p
}

func (*Expr7Context) IsExpr7Context() {}

func NewExpr7Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr7Context {
	var p = new(Expr7Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_expr7

	return p
}

func (s *Expr7Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr7Context) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Expr7Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *Expr7Context) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *Expr7Context) IdentifierImmidiate() IIdentifierImmidiateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierImmidiateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierImmidiateContext)
}

func (s *Expr7Context) Expr6() IExpr6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr6Context)
}

func (s *Expr7Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr7Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr7Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterExpr7(s)
	}
}

func (s *Expr7Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitExpr7(s)
	}
}

func (p *AuroraParser) Expr7() (localctx IExpr7Context) {
	localctx = NewExpr7Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AuroraParserRULE_expr7)

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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(281)
				p.expr1(0)
			}
			{
				p.SetState(282)
				p.MemberAccess()
			}

		case 2:
			{
				p.SetState(284)
				p.IdentifierImmidiate()
			}

		}
		{
			p.SetState(287)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.expr6(0)
		}

	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	right  IExprContext
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuroraParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuroraParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) GetRight() IExprContext { return s.right }

func (s *AssignContext) SetRight(v IExprContext) { s.right = v }

func (s *AssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuroraListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *AuroraParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AuroraParserRULE_assign)

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
		p.SetState(292)
		p.Match(AuroraParserT__7)
	}
	{
		p.SetState(293)

		var _x = p.Expr()

		localctx.(*AssignContext).right = _x
	}

	return localctx
}

func (p *AuroraParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *Expr1Context = nil
		if localctx != nil {
			t = localctx.(*Expr1Context)
		}
		return p.Expr1_Sempred(t, predIndex)

	case 25:
		var t *Expr3Context = nil
		if localctx != nil {
			t = localctx.(*Expr3Context)
		}
		return p.Expr3_Sempred(t, predIndex)

	case 27:
		var t *Expr4Context = nil
		if localctx != nil {
			t = localctx.(*Expr4Context)
		}
		return p.Expr4_Sempred(t, predIndex)

	case 29:
		var t *Expr5Context = nil
		if localctx != nil {
			t = localctx.(*Expr5Context)
		}
		return p.Expr5_Sempred(t, predIndex)

	case 31:
		var t *Expr6Context = nil
		if localctx != nil {
			t = localctx.(*Expr6Context)
		}
		return p.Expr6_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AuroraParser) Expr1_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AuroraParser) Expr3_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AuroraParser) Expr4_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AuroraParser) Expr5_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AuroraParser) Expr6_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
