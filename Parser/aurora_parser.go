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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 305,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 5, 2, 74, 10, 2, 3, 2, 5, 2, 77,
	10, 2, 3, 3, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 90, 10, 4, 3, 5, 3, 5, 7, 5, 94, 10, 5, 12, 5, 14, 5,
	97, 11, 5, 3, 6, 3, 6, 5, 6, 101, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 5, 8, 109, 10, 8, 3, 8, 5, 8, 112, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 118, 10, 8, 3, 8, 5, 8, 121, 10, 8, 5, 8, 123, 10, 8, 3, 8, 3, 8,
	5, 8, 127, 10, 8, 3, 9, 3, 9, 5, 9, 131, 10, 9, 3, 9, 5, 9, 134, 10, 9,
	3, 9, 3, 9, 3, 9, 5, 9, 139, 10, 9, 3, 9, 3, 9, 3, 9, 3, 10, 5, 10, 145,
	10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 151, 10, 10, 3, 10, 5, 10, 154,
	10, 10, 3, 10, 5, 10, 157, 10, 10, 3, 11, 3, 11, 3, 11, 7, 11, 162, 10,
	11, 12, 11, 14, 11, 165, 11, 11, 3, 12, 6, 12, 168, 10, 12, 13, 12, 14,
	12, 169, 3, 13, 3, 13, 5, 13, 174, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13,
	179, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 189, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 207, 10,
	21, 7, 21, 209, 10, 21, 12, 21, 14, 21, 212, 11, 21, 3, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 5, 23, 219, 10, 23, 3, 23, 3, 23, 7, 23, 223, 10, 23, 12,
	23, 14, 23, 226, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25,
	234, 10, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 7,
	27, 244, 10, 27, 12, 27, 14, 27, 247, 11, 27, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 257, 10, 29, 12, 29, 14, 29, 260, 11,
	29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 270,
	10, 31, 12, 31, 14, 31, 273, 11, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 7, 33, 283, 10, 33, 12, 33, 14, 33, 286, 11, 33, 3,
	34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 295, 10, 35, 3, 35,
	3, 35, 3, 35, 5, 35, 300, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 2, 7, 40,
	52, 56, 60, 64, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 2, 6, 4, 3, 3, 3, 31, 31, 5, 2, 16, 16, 18, 18, 20, 22, 4,
	2, 18, 18, 23, 24, 3, 2, 20, 21, 2, 306, 2, 73, 3, 2, 2, 2, 4, 78, 3, 2,
	2, 2, 6, 89, 3, 2, 2, 2, 8, 91, 3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12, 104,
	3, 2, 2, 2, 14, 106, 3, 2, 2, 2, 16, 128, 3, 2, 2, 2, 18, 144, 3, 2, 2,
	2, 20, 158, 3, 2, 2, 2, 22, 167, 3, 2, 2, 2, 24, 178, 3, 2, 2, 2, 26, 180,
	3, 2, 2, 2, 28, 182, 3, 2, 2, 2, 30, 188, 3, 2, 2, 2, 32, 190, 3, 2, 2,
	2, 34, 194, 3, 2, 2, 2, 36, 196, 3, 2, 2, 2, 38, 198, 3, 2, 2, 2, 40, 200,
	3, 2, 2, 2, 42, 213, 3, 2, 2, 2, 44, 216, 3, 2, 2, 2, 46, 229, 3, 2, 2,
	2, 48, 233, 3, 2, 2, 2, 50, 235, 3, 2, 2, 2, 52, 238, 3, 2, 2, 2, 54, 248,
	3, 2, 2, 2, 56, 251, 3, 2, 2, 2, 58, 261, 3, 2, 2, 2, 60, 264, 3, 2, 2,
	2, 62, 274, 3, 2, 2, 2, 64, 277, 3, 2, 2, 2, 66, 287, 3, 2, 2, 2, 68, 299,
	3, 2, 2, 2, 70, 301, 3, 2, 2, 2, 72, 74, 5, 4, 3, 2, 73, 72, 3, 2, 2, 2,
	73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77, 7, 2, 2, 3, 76, 75, 3,
	2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 3, 3, 2, 2, 2, 78, 82, 5, 6, 4, 2, 79,
	81, 5, 6, 4, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 86, 5,
	12, 7, 2, 86, 87, 7, 3, 2, 2, 87, 90, 3, 2, 2, 2, 88, 90, 5, 16, 9, 2,
	89, 85, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 7, 3, 2, 2, 2, 91, 95, 5, 24,
	13, 2, 92, 94, 5, 24, 13, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95,
	93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 95, 3, 2, 2,
	2, 98, 100, 7, 4, 2, 2, 99, 101, 5, 8, 5, 2, 100, 99, 3, 2, 2, 2, 100,
	101, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103, 7, 5, 2, 2, 103, 11, 3,
	2, 2, 2, 104, 105, 5, 14, 8, 2, 105, 13, 3, 2, 2, 2, 106, 108, 7, 6, 2,
	2, 107, 109, 7, 7, 2, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	111, 3, 2, 2, 2, 110, 112, 7, 8, 2, 2, 111, 110, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 7, 29, 2, 2, 114, 122, 7, 9,
	2, 2, 115, 117, 7, 29, 2, 2, 116, 118, 5, 22, 12, 2, 117, 116, 3, 2, 2,
	2, 117, 118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 121, 7, 10, 2, 2, 120,
	119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 115,
	3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 125, 7, 11,
	2, 2, 125, 127, 5, 28, 15, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2,
	2, 127, 15, 3, 2, 2, 2, 128, 130, 7, 12, 2, 2, 129, 131, 7, 7, 2, 2, 130,
	129, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 133, 3, 2, 2, 2, 132, 134,
	7, 13, 2, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2,
	2, 2, 135, 136, 7, 29, 2, 2, 136, 138, 7, 14, 2, 2, 137, 139, 5, 20, 11,
	2, 138, 137, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140,
	141, 7, 15, 2, 2, 141, 142, 5, 10, 6, 2, 142, 17, 3, 2, 2, 2, 143, 145,
	7, 8, 2, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2,
	2, 2, 146, 147, 7, 29, 2, 2, 147, 148, 7, 9, 2, 2, 148, 156, 7, 29, 2,
	2, 149, 151, 5, 22, 12, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2,
	151, 153, 3, 2, 2, 2, 152, 154, 7, 10, 2, 2, 153, 152, 3, 2, 2, 2, 153,
	154, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 157, 7, 16, 2, 2, 156, 150,
	3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 19, 3, 2, 2, 2, 158, 163, 5, 18,
	10, 2, 159, 160, 7, 17, 2, 2, 160, 162, 5, 18, 10, 2, 161, 159, 3, 2, 2,
	2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164,
	21, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 168, 7, 18, 2, 2, 167, 166,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2,
	2, 2, 170, 23, 3, 2, 2, 2, 171, 174, 5, 14, 8, 2, 172, 174, 5, 26, 14,
	2, 173, 171, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175,
	176, 9, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 179, 5, 10, 6, 2, 178, 173,
	3, 2, 2, 2, 178, 177, 3, 2, 2, 2, 179, 25, 3, 2, 2, 2, 180, 181, 5, 28,
	15, 2, 181, 27, 3, 2, 2, 2, 182, 183, 5, 68, 35, 2, 183, 29, 3, 2, 2, 2,
	184, 189, 5, 32, 17, 2, 185, 189, 5, 34, 18, 2, 186, 189, 5, 38, 20, 2,
	187, 189, 5, 36, 19, 2, 188, 184, 3, 2, 2, 2, 188, 185, 3, 2, 2, 2, 188,
	186, 3, 2, 2, 2, 188, 187, 3, 2, 2, 2, 189, 31, 3, 2, 2, 2, 190, 191, 7,
	14, 2, 2, 191, 192, 5, 28, 15, 2, 192, 193, 7, 15, 2, 2, 193, 33, 3, 2,
	2, 2, 194, 195, 7, 28, 2, 2, 195, 35, 3, 2, 2, 2, 196, 197, 7, 29, 2, 2,
	197, 37, 3, 2, 2, 2, 198, 199, 7, 27, 2, 2, 199, 39, 3, 2, 2, 2, 200, 201,
	8, 21, 1, 2, 201, 202, 5, 30, 16, 2, 202, 210, 3, 2, 2, 2, 203, 206, 12,
	4, 2, 2, 204, 207, 5, 42, 22, 2, 205, 207, 5, 44, 23, 2, 206, 204, 3, 2,
	2, 2, 206, 205, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 203, 3, 2, 2, 2,
	209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211,
	41, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 7, 19, 2, 2, 214, 215,
	7, 29, 2, 2, 215, 43, 3, 2, 2, 2, 216, 218, 7, 14, 2, 2, 217, 219, 5, 46,
	24, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 224, 3, 2, 2, 2,
	220, 221, 7, 17, 2, 2, 221, 223, 5, 46, 24, 2, 222, 220, 3, 2, 2, 2, 223,
	226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 227,
	3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 227, 228, 7, 15, 2, 2, 228, 45, 3, 2,
	2, 2, 229, 230, 5, 28, 15, 2, 230, 47, 3, 2, 2, 2, 231, 234, 5, 50, 26,
	2, 232, 234, 5, 40, 21, 2, 233, 231, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2,
	234, 49, 3, 2, 2, 2, 235, 236, 9, 3, 2, 2, 236, 237, 5, 28, 15, 2, 237,
	51, 3, 2, 2, 2, 238, 239, 8, 27, 1, 2, 239, 240, 5, 48, 25, 2, 240, 245,
	3, 2, 2, 2, 241, 242, 12, 4, 2, 2, 242, 244, 5, 54, 28, 2, 243, 241, 3,
	2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2,
	2, 246, 53, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 249, 9, 4, 2, 2, 249,
	250, 5, 52, 27, 2, 250, 55, 3, 2, 2, 2, 251, 252, 8, 29, 1, 2, 252, 253,
	5, 52, 27, 2, 253, 258, 3, 2, 2, 2, 254, 255, 12, 4, 2, 2, 255, 257, 5,
	58, 30, 2, 256, 254, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2,
	2, 2, 258, 259, 3, 2, 2, 2, 259, 57, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2,
	261, 262, 9, 5, 2, 2, 262, 263, 5, 56, 29, 2, 263, 59, 3, 2, 2, 2, 264,
	265, 8, 31, 1, 2, 265, 266, 5, 56, 29, 2, 266, 271, 3, 2, 2, 2, 267, 268,
	12, 4, 2, 2, 268, 270, 5, 62, 32, 2, 269, 267, 3, 2, 2, 2, 270, 273, 3,
	2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 61, 3, 2, 2,
	2, 273, 271, 3, 2, 2, 2, 274, 275, 7, 25, 2, 2, 275, 276, 5, 60, 31, 2,
	276, 63, 3, 2, 2, 2, 277, 278, 8, 33, 1, 2, 278, 279, 5, 60, 31, 2, 279,
	284, 3, 2, 2, 2, 280, 281, 12, 4, 2, 2, 281, 283, 5, 66, 34, 2, 282, 280,
	3, 2, 2, 2, 283, 286, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2,
	2, 2, 285, 65, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 288, 7, 26, 2, 2,
	288, 289, 5, 64, 33, 2, 289, 67, 3, 2, 2, 2, 290, 291, 5, 40, 21, 2, 291,
	292, 5, 42, 22, 2, 292, 295, 3, 2, 2, 2, 293, 295, 5, 36, 19, 2, 294, 290,
	3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 5, 70,
	36, 2, 297, 300, 3, 2, 2, 2, 298, 300, 5, 64, 33, 2, 299, 294, 3, 2, 2,
	2, 299, 298, 3, 2, 2, 2, 300, 69, 3, 2, 2, 2, 301, 302, 7, 11, 2, 2, 302,
	303, 5, 28, 15, 2, 303, 71, 3, 2, 2, 2, 37, 73, 76, 82, 89, 95, 100, 108,
	111, 117, 120, 122, 126, 130, 133, 138, 144, 150, 153, 156, 163, 169, 173,
	178, 188, 206, 210, 218, 224, 233, 245, 258, 271, 284, 294, 299,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'{'", "'}'", "'var'", "'export'", "'optional'", "':'", "'^'",
	"'='", "'func'", "'native'", "'('", "')'", "'&'", "','", "'*'", "'.'",
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

	if _la == AuroraParserT__3 || _la == AuroraParserT__9 {
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

	for _la == AuroraParserT__3 || _la == AuroraParserT__9 {
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.VariableDeclaration()
		}
		{
			p.SetState(84)
			p.Match(AuroraParserT__0)
		}

	case AuroraParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
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
		p.SetState(89)
		p.Stmt()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__1)|(1<<AuroraParserT__3)|(1<<AuroraParserT__11)|(1<<AuroraParserT__13)|(1<<AuroraParserT__15)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(90)
			p.Stmt()
		}

		p.SetState(95)
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
		p.SetState(96)
		p.Match(AuroraParserT__1)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__1)|(1<<AuroraParserT__3)|(1<<AuroraParserT__11)|(1<<AuroraParserT__13)|(1<<AuroraParserT__15)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(97)
			p.StatementList()
		}

	}
	{
		p.SetState(100)
		p.Match(AuroraParserT__2)
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
		p.SetState(102)
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
		p.SetState(104)
		p.Match(AuroraParserT__3)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__4 {
		{
			p.SetState(105)

			var _m = p.Match(AuroraParserT__4)

			localctx.(*VariableStmtContext).export = _m
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__5 {
		{
			p.SetState(108)

			var _m = p.Match(AuroraParserT__5)

			localctx.(*VariableStmtContext).optional = _m
		}

	}
	{
		p.SetState(111)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*VariableStmtContext).name = _m
	}
	{
		p.SetState(112)
		p.Match(AuroraParserT__6)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserIdentifier {
		{
			p.SetState(113)

			var _m = p.Match(AuroraParserIdentifier)

			localctx.(*VariableStmtContext).variableType = _m
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__15 {
			{
				p.SetState(114)

				var _x = p.Pointer()

				localctx.(*VariableStmtContext).ptr = _x
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__7 {
			{
				p.SetState(117)

				var _m = p.Match(AuroraParserT__7)

				localctx.(*VariableStmtContext).gcptr = _m
			}

		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__8 {
		{
			p.SetState(122)
			p.Match(AuroraParserT__8)
		}
		{
			p.SetState(123)
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

func (s *FunctionDeclarationContext) FunctionParameterList() IFunctionParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterListContext)
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
		p.SetState(126)
		p.Match(AuroraParserT__9)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__4 {
		{
			p.SetState(127)

			var _m = p.Match(AuroraParserT__4)

			localctx.(*FunctionDeclarationContext).export = _m
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__10 {
		{
			p.SetState(130)

			var _m = p.Match(AuroraParserT__10)

			localctx.(*FunctionDeclarationContext).native = _m
		}

	}
	{
		p.SetState(133)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*FunctionDeclarationContext).name = _m
	}
	{
		p.SetState(134)
		p.Match(AuroraParserT__11)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__5 || _la == AuroraParserIdentifier {
		{
			p.SetState(135)
			p.FunctionParameterList()
		}

	}
	{
		p.SetState(138)
		p.Match(AuroraParserT__12)
	}
	{
		p.SetState(139)

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

	// GetOptional returns the optional token.
	GetOptional() antlr.Token

	// GetParameterName returns the parameterName token.
	GetParameterName() antlr.Token

	// GetParameterType returns the parameterType token.
	GetParameterType() antlr.Token

	// GetGcptr returns the gcptr token.
	GetGcptr() antlr.Token

	// GetRef returns the ref token.
	GetRef() antlr.Token

	// SetOptional sets the optional token.
	SetOptional(antlr.Token)

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
	optional      antlr.Token
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

func (s *FunctionParameterContext) GetOptional() antlr.Token { return s.optional }

func (s *FunctionParameterContext) GetParameterName() antlr.Token { return s.parameterName }

func (s *FunctionParameterContext) GetParameterType() antlr.Token { return s.parameterType }

func (s *FunctionParameterContext) GetGcptr() antlr.Token { return s.gcptr }

func (s *FunctionParameterContext) GetRef() antlr.Token { return s.ref }

func (s *FunctionParameterContext) SetOptional(v antlr.Token) { s.optional = v }

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AuroraParserT__5 {
		{
			p.SetState(141)

			var _m = p.Match(AuroraParserT__5)

			localctx.(*FunctionParameterContext).optional = _m
		}

	}
	{
		p.SetState(144)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*FunctionParameterContext).parameterName = _m
	}
	{
		p.SetState(145)
		p.Match(AuroraParserT__6)
	}
	{
		p.SetState(146)

		var _m = p.Match(AuroraParserIdentifier)

		localctx.(*FunctionParameterContext).parameterType = _m
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__7, AuroraParserT__12, AuroraParserT__14, AuroraParserT__15:
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__15 {
			{
				p.SetState(147)

				var _x = p.Pointer()

				localctx.(*FunctionParameterContext).ptr = _x
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AuroraParserT__7 {
			{
				p.SetState(150)

				var _m = p.Match(AuroraParserT__7)

				localctx.(*FunctionParameterContext).gcptr = _m
			}

		}

	case AuroraParserT__13:
		{
			p.SetState(153)

			var _m = p.Match(AuroraParserT__13)

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
		p.SetState(156)
		p.FunctionParameter()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__14 {
		{
			p.SetState(157)
			p.Match(AuroraParserT__14)
		}
		{
			p.SetState(158)
			p.FunctionParameter()
		}

		p.SetState(163)
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
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AuroraParserT__15 {
		{
			p.SetState(164)
			p.Match(AuroraParserT__15)
		}

		p.SetState(167)
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
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__3, AuroraParserT__11, AuroraParserT__13, AuroraParserT__15, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
		p.SetState(171)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AuroraParserT__3:
			{
				p.SetState(169)
				p.VariableStmt()
			}

		case AuroraParserT__11, AuroraParserT__13, AuroraParserT__15, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
			{
				p.SetState(170)
				p.ExpressionStmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(173)
		_la = p.GetTokenStream().LA(1)

		if !(((_la - -1)&-(0x1f+1)) == 0 && ((1<<uint((_la - -1)))&((1<<(AuroraParserEOF - -1))|(1<<(AuroraParserT__0 - -1))|(1<<(AuroraParserNewline - -1)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case AuroraParserT__1:
		{
			p.SetState(175)
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
		p.SetState(178)
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
		p.SetState(180)
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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__11:
		{
			p.SetState(182)
			p.Brackets()
		}

	case AuroraParserStringLiteral:
		{
			p.SetState(183)
			p.StringImmidiate()
		}

	case AuroraParserIntegerLiteral:
		{
			p.SetState(184)
			p.IntegerImmidiate()
		}

	case AuroraParserIdentifier:
		{
			p.SetState(185)
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
		p.SetState(188)
		p.Match(AuroraParserT__11)
	}
	{
		p.SetState(189)

		var _x = p.Expr()

		localctx.(*BracketsContext).paramExpr = _x
	}
	{
		p.SetState(190)
		p.Match(AuroraParserT__12)
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
		p.SetState(192)

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
		p.SetState(194)

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
		p.SetState(196)

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
		p.SetState(199)
		p.Expr0()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr1Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr1)
			p.SetState(201)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(204)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AuroraParserT__16:
				{
					p.SetState(202)
					p.MemberAccess()
				}

			case AuroraParserT__11:
				{
					p.SetState(203)
					p.FunctionCall()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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
		p.SetState(211)
		p.Match(AuroraParserT__16)
	}
	{
		p.SetState(212)

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
		p.SetState(214)
		p.Match(AuroraParserT__11)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__11)|(1<<AuroraParserT__13)|(1<<AuroraParserT__15)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19)|(1<<AuroraParserIntegerLiteral)|(1<<AuroraParserStringLiteral)|(1<<AuroraParserIdentifier))) != 0 {
		{
			p.SetState(215)
			p.FunctionCallParam()
		}

	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AuroraParserT__14 {
		{
			p.SetState(218)
			p.Match(AuroraParserT__14)
		}
		{
			p.SetState(219)
			p.FunctionCallParam()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(AuroraParserT__12)
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
		p.SetState(227)

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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AuroraParserT__13, AuroraParserT__15, AuroraParserT__17, AuroraParserT__18, AuroraParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.LhsOperator()
		}

	case AuroraParserT__11, AuroraParserIntegerLiteral, AuroraParserStringLiteral, AuroraParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
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
	p.SetState(233)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*LhsOperatorContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__13)|(1<<AuroraParserT__15)|(1<<AuroraParserT__17)|(1<<AuroraParserT__18)|(1<<AuroraParserT__19))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*LhsOperatorContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(234)

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
		p.SetState(237)
		p.Expr2()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr3Context(p, _parentctx, _parentState)
			localctx.(*Expr3Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr3)
			p.SetState(239)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(240)
				p.MulDivMod()
			}

		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	p.SetState(246)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*MulDivModContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AuroraParserT__15)|(1<<AuroraParserT__20)|(1<<AuroraParserT__21))) != 0) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*MulDivModContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(247)

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
		p.SetState(250)
		p.expr3(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr4Context(p, _parentctx, _parentState)
			localctx.(*Expr4Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr4)
			p.SetState(252)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(253)
				p.AddSub()
			}

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.SetState(259)

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
		p.SetState(260)

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
		p.SetState(263)
		p.expr4(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr5Context(p, _parentctx, _parentState)
			localctx.(*Expr5Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr5)
			p.SetState(265)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(266)
				p.LogicalAnd()
			}

		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
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
		p.SetState(272)
		p.Match(AuroraParserT__22)
	}
	{
		p.SetState(273)

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
		p.SetState(276)
		p.expr5(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr6Context(p, _parentctx, _parentState)
			localctx.(*Expr6Context).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, AuroraParserRULE_expr6)
			p.SetState(278)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(279)
				p.LogicalOr()
			}

		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
		p.SetState(285)
		p.Match(AuroraParserT__23)
	}
	{
		p.SetState(286)

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

	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(288)
				p.expr1(0)
			}
			{
				p.SetState(289)
				p.MemberAccess()
			}

		case 2:
			{
				p.SetState(291)
				p.IdentifierImmidiate()
			}

		}
		{
			p.SetState(294)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
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
		p.SetState(299)
		p.Match(AuroraParserT__8)
	}
	{
		p.SetState(300)

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
