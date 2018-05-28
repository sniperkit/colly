// Generated from AntlrThrift.g4 by ANTLR 4.7.

package thrift // AntlrThrift

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 58, 411,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 7, 2, 80, 10, 2, 12, 2, 14, 2, 83, 11, 2, 3, 2, 7, 2, 86, 10,
	2, 12, 2, 14, 2, 89, 11, 2, 3, 3, 3, 3, 3, 3, 5, 3, 94, 10, 3, 3, 4, 3,
	4, 3, 4, 5, 4, 99, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 104, 10, 5, 3, 6, 3,
	6, 3, 6, 5, 6, 109, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 115, 10, 7, 3,
	7, 5, 7, 118, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 125, 10, 9, 3,
	9, 5, 9, 128, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 133, 10, 10, 3, 10, 5,
	10, 136, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 146, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 154,
	10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 162, 10, 13, 3,
	14, 3, 14, 3, 14, 5, 14, 167, 10, 14, 7, 14, 169, 10, 14, 12, 14, 14, 14,
	172, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 15, 179, 10, 15, 7,
	15, 181, 10, 15, 12, 15, 14, 15, 184, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 196, 10, 17, 3, 17, 5,
	17, 199, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 205, 10, 18, 12, 18,
	14, 18, 208, 11, 18, 3, 18, 3, 18, 5, 18, 212, 10, 18, 3, 19, 3, 19, 3,
	19, 5, 19, 217, 10, 19, 3, 19, 5, 19, 220, 10, 19, 3, 19, 5, 19, 223, 10,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 229, 10, 20, 12, 20, 14, 20, 232,
	11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 240, 10, 21, 12,
	21, 14, 21, 243, 11, 21, 3, 21, 3, 21, 5, 21, 247, 10, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 7, 22, 253, 10, 22, 12, 22, 14, 22, 256, 11, 22, 3, 22, 3,
	22, 5, 22, 260, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 266, 10, 23,
	12, 23, 14, 23, 269, 11, 23, 3, 23, 3, 23, 5, 23, 273, 10, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 279, 10, 24, 3, 24, 3, 24, 7, 24, 283, 10, 24,
	12, 24, 14, 24, 286, 11, 24, 3, 24, 3, 24, 5, 24, 290, 10, 24, 3, 25, 5,
	25, 293, 10, 25, 3, 25, 3, 25, 5, 25, 297, 10, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 302, 10, 25, 3, 25, 5, 25, 305, 10, 25, 3, 25, 5, 25, 308, 10, 25,
	3, 26, 3, 26, 7, 26, 312, 10, 26, 12, 26, 14, 26, 315, 11, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 5, 27, 321, 10, 27, 3, 27, 5, 27, 324, 10, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 5, 27, 330, 10, 27, 3, 27, 5, 27, 333, 10, 27, 3,
	27, 5, 27, 336, 10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	5, 30, 345, 10, 30, 3, 30, 3, 30, 5, 30, 349, 10, 30, 3, 30, 3, 30, 5,
	30, 353, 10, 30, 5, 30, 355, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	5, 32, 362, 10, 32, 3, 33, 3, 33, 5, 33, 366, 10, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 379,
	10, 34, 3, 35, 3, 35, 5, 35, 383, 10, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 7, 37, 394, 10, 37, 12, 37, 14, 37, 397,
	11, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 5, 38, 404, 10, 38, 3, 38, 5,
	38, 407, 10, 38, 3, 39, 3, 39, 3, 39, 2, 2, 40, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 2, 6, 4, 2, 6, 6, 50,
	50, 3, 2, 26, 27, 3, 2, 29, 38, 3, 2, 48, 49, 2, 444, 2, 81, 3, 2, 2, 2,
	4, 93, 3, 2, 2, 2, 6, 95, 3, 2, 2, 2, 8, 100, 3, 2, 2, 2, 10, 108, 3, 2,
	2, 2, 12, 110, 3, 2, 2, 2, 14, 119, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18,
	129, 3, 2, 2, 2, 20, 145, 3, 2, 2, 2, 22, 147, 3, 2, 2, 2, 24, 161, 3,
	2, 2, 2, 26, 163, 3, 2, 2, 2, 28, 175, 3, 2, 2, 2, 30, 187, 3, 2, 2, 2,
	32, 191, 3, 2, 2, 2, 34, 200, 3, 2, 2, 2, 36, 213, 3, 2, 2, 2, 38, 224,
	3, 2, 2, 2, 40, 235, 3, 2, 2, 2, 42, 248, 3, 2, 2, 2, 44, 261, 3, 2, 2,
	2, 46, 274, 3, 2, 2, 2, 48, 292, 3, 2, 2, 2, 50, 309, 3, 2, 2, 2, 52, 320,
	3, 2, 2, 2, 54, 337, 3, 2, 2, 2, 56, 339, 3, 2, 2, 2, 58, 354, 3, 2, 2,
	2, 60, 356, 3, 2, 2, 2, 62, 361, 3, 2, 2, 2, 64, 363, 3, 2, 2, 2, 66, 373,
	3, 2, 2, 2, 68, 380, 3, 2, 2, 2, 70, 388, 3, 2, 2, 2, 72, 391, 3, 2, 2,
	2, 74, 400, 3, 2, 2, 2, 76, 408, 3, 2, 2, 2, 78, 80, 5, 4, 3, 2, 79, 78,
	3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 87, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 86, 5, 20, 11, 2, 85, 84, 3,
	2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88,
	3, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 94, 5, 6, 4, 2, 91, 94, 5, 8, 5,
	2, 92, 94, 5, 10, 6, 2, 93, 90, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 92,
	3, 2, 2, 2, 94, 5, 3, 2, 2, 2, 95, 96, 7, 3, 2, 2, 96, 98, 7, 45, 2, 2,
	97, 99, 5, 76, 39, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 7, 3,
	2, 2, 2, 100, 101, 7, 4, 2, 2, 101, 103, 7, 45, 2, 2, 102, 104, 5, 76,
	39, 2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 9, 3, 2, 2, 2,
	105, 109, 5, 12, 7, 2, 106, 109, 5, 16, 9, 2, 107, 109, 5, 18, 10, 2, 108,
	105, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 11, 3,
	2, 2, 2, 110, 111, 7, 5, 2, 2, 111, 112, 5, 14, 8, 2, 112, 114, 7, 50,
	2, 2, 113, 115, 5, 72, 37, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2, 2,
	2, 115, 117, 3, 2, 2, 2, 116, 118, 5, 76, 39, 2, 117, 116, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120, 9, 2, 2, 2, 120, 15,
	3, 2, 2, 2, 121, 122, 7, 7, 2, 2, 122, 124, 7, 45, 2, 2, 123, 125, 5, 72,
	37, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2,
	126, 128, 5, 76, 39, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	17, 3, 2, 2, 2, 129, 130, 7, 8, 2, 2, 130, 132, 7, 45, 2, 2, 131, 133,
	5, 72, 37, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3,
	2, 2, 2, 134, 136, 5, 76, 39, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2,
	2, 2, 136, 19, 3, 2, 2, 2, 137, 146, 5, 22, 12, 2, 138, 146, 5, 32, 17,
	2, 139, 146, 5, 34, 18, 2, 140, 146, 5, 38, 20, 2, 141, 146, 5, 40, 21,
	2, 142, 146, 5, 42, 22, 2, 143, 146, 5, 44, 23, 2, 144, 146, 5, 46, 24,
	2, 145, 137, 3, 2, 2, 2, 145, 138, 3, 2, 2, 2, 145, 139, 3, 2, 2, 2, 145,
	140, 3, 2, 2, 2, 145, 141, 3, 2, 2, 2, 145, 142, 3, 2, 2, 2, 145, 143,
	3, 2, 2, 2, 145, 144, 3, 2, 2, 2, 146, 21, 3, 2, 2, 2, 147, 148, 7, 9,
	2, 2, 148, 149, 5, 58, 30, 2, 149, 150, 7, 50, 2, 2, 150, 151, 7, 10, 2,
	2, 151, 153, 5, 24, 13, 2, 152, 154, 5, 76, 39, 2, 153, 152, 3, 2, 2, 2,
	153, 154, 3, 2, 2, 2, 154, 23, 3, 2, 2, 2, 155, 162, 7, 52, 2, 2, 156,
	162, 7, 53, 2, 2, 157, 162, 7, 45, 2, 2, 158, 162, 7, 50, 2, 2, 159, 162,
	5, 26, 14, 2, 160, 162, 5, 28, 15, 2, 161, 155, 3, 2, 2, 2, 161, 156, 3,
	2, 2, 2, 161, 157, 3, 2, 2, 2, 161, 158, 3, 2, 2, 2, 161, 159, 3, 2, 2,
	2, 161, 160, 3, 2, 2, 2, 162, 25, 3, 2, 2, 2, 163, 170, 7, 11, 2, 2, 164,
	166, 5, 24, 13, 2, 165, 167, 5, 76, 39, 2, 166, 165, 3, 2, 2, 2, 166, 167,
	3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 164, 3, 2, 2, 2, 169, 172, 3, 2,
	2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 173, 174, 7, 12, 2, 2, 174, 27, 3, 2, 2, 2, 175,
	182, 7, 13, 2, 2, 176, 178, 5, 30, 16, 2, 177, 179, 5, 76, 39, 2, 178,
	177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 176,
	3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2,
	2, 2, 183, 185, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186, 7, 14, 2, 2,
	186, 29, 3, 2, 2, 2, 187, 188, 5, 24, 13, 2, 188, 189, 7, 15, 2, 2, 189,
	190, 5, 24, 13, 2, 190, 31, 3, 2, 2, 2, 191, 192, 7, 16, 2, 2, 192, 193,
	5, 58, 30, 2, 193, 195, 7, 50, 2, 2, 194, 196, 5, 72, 37, 2, 195, 194,
	3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 199, 5, 76,
	39, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 33, 3, 2, 2, 2,
	200, 201, 7, 17, 2, 2, 201, 202, 7, 50, 2, 2, 202, 206, 7, 13, 2, 2, 203,
	205, 5, 36, 19, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 206, 3, 2,
	2, 2, 209, 211, 7, 14, 2, 2, 210, 212, 5, 72, 37, 2, 211, 210, 3, 2, 2,
	2, 211, 212, 3, 2, 2, 2, 212, 35, 3, 2, 2, 2, 213, 216, 7, 50, 2, 2, 214,
	215, 7, 10, 2, 2, 215, 217, 7, 52, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217,
	3, 2, 2, 2, 217, 219, 3, 2, 2, 2, 218, 220, 5, 72, 37, 2, 219, 218, 3,
	2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2, 221, 223, 5, 76, 39,
	2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 37, 3, 2, 2, 2, 224,
	225, 7, 18, 2, 2, 225, 226, 7, 50, 2, 2, 226, 230, 7, 13, 2, 2, 227, 229,
	5, 36, 19, 2, 228, 227, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3,
	2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 230, 3, 2, 2,
	2, 233, 234, 7, 14, 2, 2, 234, 39, 3, 2, 2, 2, 235, 236, 7, 19, 2, 2, 236,
	237, 7, 50, 2, 2, 237, 241, 7, 13, 2, 2, 238, 240, 5, 52, 27, 2, 239, 238,
	3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2,
	2, 2, 242, 244, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 246, 7, 14, 2, 2,
	245, 247, 5, 72, 37, 2, 246, 245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	41, 3, 2, 2, 2, 248, 249, 7, 20, 2, 2, 249, 250, 7, 50, 2, 2, 250, 254,
	7, 13, 2, 2, 251, 253, 5, 52, 27, 2, 252, 251, 3, 2, 2, 2, 253, 256, 3,
	2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 257, 3, 2, 2,
	2, 256, 254, 3, 2, 2, 2, 257, 259, 7, 14, 2, 2, 258, 260, 5, 72, 37, 2,
	259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 43, 3, 2, 2, 2, 261, 262,
	7, 21, 2, 2, 262, 263, 7, 50, 2, 2, 263, 267, 7, 13, 2, 2, 264, 266, 5,
	52, 27, 2, 265, 264, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3, 2,
	2, 2, 267, 268, 3, 2, 2, 2, 268, 270, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2,
	270, 272, 7, 14, 2, 2, 271, 273, 5, 72, 37, 2, 272, 271, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 45, 3, 2, 2, 2, 274, 275, 7, 22, 2, 2, 275, 278,
	7, 50, 2, 2, 276, 277, 7, 23, 2, 2, 277, 279, 5, 58, 30, 2, 278, 276, 3,
	2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 284, 7, 13, 2,
	2, 281, 283, 5, 48, 25, 2, 282, 281, 3, 2, 2, 2, 283, 286, 3, 2, 2, 2,
	284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 287, 3, 2, 2, 2, 286,
	284, 3, 2, 2, 2, 287, 289, 7, 14, 2, 2, 288, 290, 5, 72, 37, 2, 289, 288,
	3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 47, 3, 2, 2, 2, 291, 293, 7, 47,
	2, 2, 292, 291, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2,
	294, 297, 7, 46, 2, 2, 295, 297, 5, 58, 30, 2, 296, 294, 3, 2, 2, 2, 296,
	295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299, 7, 50, 2, 2, 299, 301,
	5, 50, 26, 2, 300, 302, 5, 56, 29, 2, 301, 300, 3, 2, 2, 2, 301, 302, 3,
	2, 2, 2, 302, 304, 3, 2, 2, 2, 303, 305, 5, 72, 37, 2, 304, 303, 3, 2,
	2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 308, 5, 76, 39,
	2, 307, 306, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 49, 3, 2, 2, 2, 309,
	313, 7, 24, 2, 2, 310, 312, 5, 52, 27, 2, 311, 310, 3, 2, 2, 2, 312, 315,
	3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 316, 3, 2,
	2, 2, 315, 313, 3, 2, 2, 2, 316, 317, 7, 25, 2, 2, 317, 51, 3, 2, 2, 2,
	318, 319, 7, 52, 2, 2, 319, 321, 7, 15, 2, 2, 320, 318, 3, 2, 2, 2, 320,
	321, 3, 2, 2, 2, 321, 323, 3, 2, 2, 2, 322, 324, 5, 54, 28, 2, 323, 322,
	3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 326, 5, 58,
	30, 2, 326, 329, 7, 50, 2, 2, 327, 328, 7, 10, 2, 2, 328, 330, 5, 24, 13,
	2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 332, 3, 2, 2, 2, 331,
	333, 5, 72, 37, 2, 332, 331, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335,
	3, 2, 2, 2, 334, 336, 5, 76, 39, 2, 335, 334, 3, 2, 2, 2, 335, 336, 3,
	2, 2, 2, 336, 53, 3, 2, 2, 2, 337, 338, 9, 3, 2, 2, 338, 55, 3, 2, 2, 2,
	339, 340, 7, 28, 2, 2, 340, 341, 5, 50, 26, 2, 341, 57, 3, 2, 2, 2, 342,
	344, 5, 60, 31, 2, 343, 345, 5, 72, 37, 2, 344, 343, 3, 2, 2, 2, 344, 345,
	3, 2, 2, 2, 345, 355, 3, 2, 2, 2, 346, 348, 5, 62, 32, 2, 347, 349, 5,
	72, 37, 2, 348, 347, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 355, 3, 2,
	2, 2, 350, 352, 7, 50, 2, 2, 351, 353, 5, 72, 37, 2, 352, 351, 3, 2, 2,
	2, 352, 353, 3, 2, 2, 2, 353, 355, 3, 2, 2, 2, 354, 342, 3, 2, 2, 2, 354,
	346, 3, 2, 2, 2, 354, 350, 3, 2, 2, 2, 355, 59, 3, 2, 2, 2, 356, 357, 9,
	4, 2, 2, 357, 61, 3, 2, 2, 2, 358, 362, 5, 64, 33, 2, 359, 362, 5, 68,
	35, 2, 360, 362, 5, 66, 34, 2, 361, 358, 3, 2, 2, 2, 361, 359, 3, 2, 2,
	2, 361, 360, 3, 2, 2, 2, 362, 63, 3, 2, 2, 2, 363, 365, 7, 39, 2, 2, 364,
	366, 5, 70, 36, 2, 365, 364, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367,
	3, 2, 2, 2, 367, 368, 7, 40, 2, 2, 368, 369, 5, 58, 30, 2, 369, 370, 7,
	48, 2, 2, 370, 371, 5, 58, 30, 2, 371, 372, 7, 41, 2, 2, 372, 65, 3, 2,
	2, 2, 373, 374, 7, 42, 2, 2, 374, 375, 7, 40, 2, 2, 375, 376, 5, 58, 30,
	2, 376, 378, 7, 41, 2, 2, 377, 379, 5, 70, 36, 2, 378, 377, 3, 2, 2, 2,
	378, 379, 3, 2, 2, 2, 379, 67, 3, 2, 2, 2, 380, 382, 7, 43, 2, 2, 381,
	383, 5, 70, 36, 2, 382, 381, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 384,
	3, 2, 2, 2, 384, 385, 7, 40, 2, 2, 385, 386, 5, 58, 30, 2, 386, 387, 7,
	41, 2, 2, 387, 69, 3, 2, 2, 2, 388, 389, 7, 44, 2, 2, 389, 390, 7, 45,
	2, 2, 390, 71, 3, 2, 2, 2, 391, 395, 7, 24, 2, 2, 392, 394, 5, 74, 38,
	2, 393, 392, 3, 2, 2, 2, 394, 397, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 398, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 398, 399,
	7, 25, 2, 2, 399, 73, 3, 2, 2, 2, 400, 403, 7, 50, 2, 2, 401, 402, 7, 10,
	2, 2, 402, 404, 7, 45, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2,
	404, 406, 3, 2, 2, 2, 405, 407, 5, 76, 39, 2, 406, 405, 3, 2, 2, 2, 406,
	407, 3, 2, 2, 2, 407, 75, 3, 2, 2, 2, 408, 409, 9, 5, 2, 2, 409, 77, 3,
	2, 2, 2, 60, 81, 87, 93, 98, 103, 108, 114, 117, 124, 127, 132, 135, 145,
	153, 161, 166, 170, 178, 182, 195, 198, 206, 211, 216, 219, 222, 230, 241,
	246, 254, 259, 267, 272, 278, 284, 289, 292, 296, 301, 304, 307, 313, 320,
	323, 329, 332, 335, 344, 348, 352, 354, 361, 365, 378, 382, 395, 403, 406,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'include'", "'cpp_include'", "'namespace'", "'*'", "'php_namespace'",
	"'xsd_namespace'", "'const'", "'='", "'['", "']'", "'{'", "'}'", "':'",
	"'typedef'", "'enum'", "'senum'", "'struct'", "'union'", "'exception'",
	"'service'", "'extends'", "'('", "')'", "'optional'", "'required'", "'throws'",
	"'bool'", "'byte'", "'i8'", "'i16'", "'i32'", "'i64'", "'double'", "'string'",
	"'binary'", "'slist'", "'map'", "'<'", "'>'", "'list'", "'set'", "'cpp_type'",
	"", "'void'", "'oneway'", "','", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "LITERAL", "VOID", "ONEWAY", "COMMA", "SEMICOLON",
	"IDENTIFIER", "NS_SCOPE", "INTEGER", "DOUBLE", "SLASH_SLASH_COMMENT", "HASH_COMMENT",
	"MULTILINE_COMMENT", "NEWLINE", "WHITESPACE",
}

var ruleNames = []string{
	"document", "header", "include", "cppInclude", "namespace", "standardNamespace",
	"namespaceScope", "phpNamespace", "xsdNamespace", "definition", "constDef",
	"constValue", "constList", "constMap", "constMapEntry", "typedef", "enumDef",
	"enumMember", "senum", "structDef", "unionDef", "exceptionDef", "serviceDef",
	"function", "fieldList", "field", "requiredness", "throwsList", "fieldType",
	"baseType", "containerType", "mapType", "listType", "setType", "cppType",
	"annotationList", "annotation", "separator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AntlrThriftParser struct {
	*antlr.BaseParser
}

func NewAntlrThriftParser(input antlr.TokenStream) *AntlrThriftParser {
	this := new(AntlrThriftParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AntlrThrift.g4"

	return this
}

// AntlrThriftParser tokens.
const (
	AntlrThriftParserEOF                 = antlr.TokenEOF
	AntlrThriftParserT__0                = 1
	AntlrThriftParserT__1                = 2
	AntlrThriftParserT__2                = 3
	AntlrThriftParserT__3                = 4
	AntlrThriftParserT__4                = 5
	AntlrThriftParserT__5                = 6
	AntlrThriftParserT__6                = 7
	AntlrThriftParserT__7                = 8
	AntlrThriftParserT__8                = 9
	AntlrThriftParserT__9                = 10
	AntlrThriftParserT__10               = 11
	AntlrThriftParserT__11               = 12
	AntlrThriftParserT__12               = 13
	AntlrThriftParserT__13               = 14
	AntlrThriftParserT__14               = 15
	AntlrThriftParserT__15               = 16
	AntlrThriftParserT__16               = 17
	AntlrThriftParserT__17               = 18
	AntlrThriftParserT__18               = 19
	AntlrThriftParserT__19               = 20
	AntlrThriftParserT__20               = 21
	AntlrThriftParserT__21               = 22
	AntlrThriftParserT__22               = 23
	AntlrThriftParserT__23               = 24
	AntlrThriftParserT__24               = 25
	AntlrThriftParserT__25               = 26
	AntlrThriftParserT__26               = 27
	AntlrThriftParserT__27               = 28
	AntlrThriftParserT__28               = 29
	AntlrThriftParserT__29               = 30
	AntlrThriftParserT__30               = 31
	AntlrThriftParserT__31               = 32
	AntlrThriftParserT__32               = 33
	AntlrThriftParserT__33               = 34
	AntlrThriftParserT__34               = 35
	AntlrThriftParserT__35               = 36
	AntlrThriftParserT__36               = 37
	AntlrThriftParserT__37               = 38
	AntlrThriftParserT__38               = 39
	AntlrThriftParserT__39               = 40
	AntlrThriftParserT__40               = 41
	AntlrThriftParserT__41               = 42
	AntlrThriftParserLITERAL             = 43
	AntlrThriftParserVOID                = 44
	AntlrThriftParserONEWAY              = 45
	AntlrThriftParserCOMMA               = 46
	AntlrThriftParserSEMICOLON           = 47
	AntlrThriftParserIDENTIFIER          = 48
	AntlrThriftParserNS_SCOPE            = 49
	AntlrThriftParserINTEGER             = 50
	AntlrThriftParserDOUBLE              = 51
	AntlrThriftParserSLASH_SLASH_COMMENT = 52
	AntlrThriftParserHASH_COMMENT        = 53
	AntlrThriftParserMULTILINE_COMMENT   = 54
	AntlrThriftParserNEWLINE             = 55
	AntlrThriftParserWHITESPACE          = 56
)

// AntlrThriftParser rules.
const (
	AntlrThriftParserRULE_document          = 0
	AntlrThriftParserRULE_header            = 1
	AntlrThriftParserRULE_include           = 2
	AntlrThriftParserRULE_cppInclude        = 3
	AntlrThriftParserRULE_namespace         = 4
	AntlrThriftParserRULE_standardNamespace = 5
	AntlrThriftParserRULE_namespaceScope    = 6
	AntlrThriftParserRULE_phpNamespace      = 7
	AntlrThriftParserRULE_xsdNamespace      = 8
	AntlrThriftParserRULE_definition        = 9
	AntlrThriftParserRULE_constDef          = 10
	AntlrThriftParserRULE_constValue        = 11
	AntlrThriftParserRULE_constList         = 12
	AntlrThriftParserRULE_constMap          = 13
	AntlrThriftParserRULE_constMapEntry     = 14
	AntlrThriftParserRULE_typedef           = 15
	AntlrThriftParserRULE_enumDef           = 16
	AntlrThriftParserRULE_enumMember        = 17
	AntlrThriftParserRULE_senum             = 18
	AntlrThriftParserRULE_structDef         = 19
	AntlrThriftParserRULE_unionDef          = 20
	AntlrThriftParserRULE_exceptionDef      = 21
	AntlrThriftParserRULE_serviceDef        = 22
	AntlrThriftParserRULE_function          = 23
	AntlrThriftParserRULE_fieldList         = 24
	AntlrThriftParserRULE_field             = 25
	AntlrThriftParserRULE_requiredness      = 26
	AntlrThriftParserRULE_throwsList        = 27
	AntlrThriftParserRULE_fieldType         = 28
	AntlrThriftParserRULE_baseType          = 29
	AntlrThriftParserRULE_containerType     = 30
	AntlrThriftParserRULE_mapType           = 31
	AntlrThriftParserRULE_listType          = 32
	AntlrThriftParserRULE_setType           = 33
	AntlrThriftParserRULE_cppType           = 34
	AntlrThriftParserRULE_annotationList    = 35
	AntlrThriftParserRULE_annotation        = 36
	AntlrThriftParserRULE_separator         = 37
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllHeader() []IHeaderContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderContext)(nil)).Elem())
	var tst = make([]IHeaderContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderContext)
		}
	}

	return tst
}

func (s *DocumentContext) Header(i int) IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *AntlrThriftParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AntlrThriftParserRULE_document)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AntlrThriftParserT__0)|(1<<AntlrThriftParserT__1)|(1<<AntlrThriftParserT__2)|(1<<AntlrThriftParserT__4)|(1<<AntlrThriftParserT__5))) != 0 {
		{
			p.SetState(76)
			p.Header()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AntlrThriftParserT__6)|(1<<AntlrThriftParserT__13)|(1<<AntlrThriftParserT__14)|(1<<AntlrThriftParserT__15)|(1<<AntlrThriftParserT__16)|(1<<AntlrThriftParserT__17)|(1<<AntlrThriftParserT__18)|(1<<AntlrThriftParserT__19))) != 0 {
		{
			p.SetState(82)
			p.Definition()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Include() IIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *HeaderContext) CppInclude() ICppIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICppIncludeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICppIncludeContext)
}

func (s *HeaderContext) Namespace() INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *AntlrThriftParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AntlrThriftParserRULE_header)

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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Include()
		}

	case AntlrThriftParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.CppInclude()
		}

	case AntlrThriftParserT__2, AntlrThriftParserT__4, AntlrThriftParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Namespace()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *IncludeContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (p *AntlrThriftParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AntlrThriftParserRULE_include)
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
		p.SetState(93)
		p.Match(AntlrThriftParserT__0)
	}
	{
		p.SetState(94)
		p.Match(AntlrThriftParserLITERAL)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(95)
			p.Separator()
		}

	}

	return localctx
}

// ICppIncludeContext is an interface to support dynamic dispatch.
type ICppIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCppIncludeContext differentiates from other interfaces.
	IsCppIncludeContext()
}

type CppIncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCppIncludeContext() *CppIncludeContext {
	var p = new(CppIncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_cppInclude
	return p
}

func (*CppIncludeContext) IsCppIncludeContext() {}

func NewCppIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CppIncludeContext {
	var p = new(CppIncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_cppInclude

	return p
}

func (s *CppIncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *CppIncludeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *CppIncludeContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *CppIncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CppIncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CppIncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterCppInclude(s)
	}
}

func (s *CppIncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitCppInclude(s)
	}
}

func (p *AntlrThriftParser) CppInclude() (localctx ICppIncludeContext) {
	localctx = NewCppIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AntlrThriftParserRULE_cppInclude)
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
		p.SetState(98)
		p.Match(AntlrThriftParserT__1)
	}
	{
		p.SetState(99)
		p.Match(AntlrThriftParserLITERAL)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(100)
			p.Separator()
		}

	}

	return localctx
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_namespace
	return p
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) StandardNamespace() IStandardNamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStandardNamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStandardNamespaceContext)
}

func (s *NamespaceContext) PhpNamespace() IPhpNamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhpNamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPhpNamespaceContext)
}

func (s *NamespaceContext) XsdNamespace() IXsdNamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXsdNamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXsdNamespaceContext)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *AntlrThriftParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AntlrThriftParserRULE_namespace)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.StandardNamespace()
		}

	case AntlrThriftParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.PhpNamespace()
		}

	case AntlrThriftParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.XsdNamespace()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStandardNamespaceContext is an interface to support dynamic dispatch.
type IStandardNamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNs returns the ns token.
	GetNs() antlr.Token

	// SetNs sets the ns token.
	SetNs(antlr.Token)

	// IsStandardNamespaceContext differentiates from other interfaces.
	IsStandardNamespaceContext()
}

type StandardNamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ns     antlr.Token
}

func NewEmptyStandardNamespaceContext() *StandardNamespaceContext {
	var p = new(StandardNamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_standardNamespace
	return p
}

func (*StandardNamespaceContext) IsStandardNamespaceContext() {}

func NewStandardNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StandardNamespaceContext {
	var p = new(StandardNamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_standardNamespace

	return p
}

func (s *StandardNamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *StandardNamespaceContext) GetNs() antlr.Token { return s.ns }

func (s *StandardNamespaceContext) SetNs(v antlr.Token) { s.ns = v }

func (s *StandardNamespaceContext) NamespaceScope() INamespaceScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceScopeContext)
}

func (s *StandardNamespaceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *StandardNamespaceContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *StandardNamespaceContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *StandardNamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandardNamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StandardNamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterStandardNamespace(s)
	}
}

func (s *StandardNamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitStandardNamespace(s)
	}
}

func (p *AntlrThriftParser) StandardNamespace() (localctx IStandardNamespaceContext) {
	localctx = NewStandardNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AntlrThriftParserRULE_standardNamespace)
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
		p.SetState(108)
		p.Match(AntlrThriftParserT__2)
	}
	{
		p.SetState(109)
		p.NamespaceScope()
	}
	{
		p.SetState(110)

		var _m = p.Match(AntlrThriftParserIDENTIFIER)

		localctx.(*StandardNamespaceContext).ns = _m
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(111)
			p.AnnotationList()
		}

	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(114)
			p.Separator()
		}

	}

	return localctx
}

// INamespaceScopeContext is an interface to support dynamic dispatch.
type INamespaceScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceScopeContext differentiates from other interfaces.
	IsNamespaceScopeContext()
}

type NamespaceScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceScopeContext() *NamespaceScopeContext {
	var p = new(NamespaceScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_namespaceScope
	return p
}

func (*NamespaceScopeContext) IsNamespaceScopeContext() {}

func NewNamespaceScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceScopeContext {
	var p = new(NamespaceScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_namespaceScope

	return p
}

func (s *NamespaceScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceScopeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *NamespaceScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterNamespaceScope(s)
	}
}

func (s *NamespaceScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitNamespaceScope(s)
	}
}

func (p *AntlrThriftParser) NamespaceScope() (localctx INamespaceScopeContext) {
	localctx = NewNamespaceScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AntlrThriftParserRULE_namespaceScope)
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
	p.SetState(117)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AntlrThriftParserT__3 || _la == AntlrThriftParserIDENTIFIER) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPhpNamespaceContext is an interface to support dynamic dispatch.
type IPhpNamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNs returns the ns token.
	GetNs() antlr.Token

	// SetNs sets the ns token.
	SetNs(antlr.Token)

	// IsPhpNamespaceContext differentiates from other interfaces.
	IsPhpNamespaceContext()
}

type PhpNamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ns     antlr.Token
}

func NewEmptyPhpNamespaceContext() *PhpNamespaceContext {
	var p = new(PhpNamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_phpNamespace
	return p
}

func (*PhpNamespaceContext) IsPhpNamespaceContext() {}

func NewPhpNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhpNamespaceContext {
	var p = new(PhpNamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_phpNamespace

	return p
}

func (s *PhpNamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *PhpNamespaceContext) GetNs() antlr.Token { return s.ns }

func (s *PhpNamespaceContext) SetNs(v antlr.Token) { s.ns = v }

func (s *PhpNamespaceContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *PhpNamespaceContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *PhpNamespaceContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *PhpNamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhpNamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhpNamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterPhpNamespace(s)
	}
}

func (s *PhpNamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitPhpNamespace(s)
	}
}

func (p *AntlrThriftParser) PhpNamespace() (localctx IPhpNamespaceContext) {
	localctx = NewPhpNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AntlrThriftParserRULE_phpNamespace)
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
		p.SetState(119)
		p.Match(AntlrThriftParserT__4)
	}
	{
		p.SetState(120)

		var _m = p.Match(AntlrThriftParserLITERAL)

		localctx.(*PhpNamespaceContext).ns = _m
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(121)
			p.AnnotationList()
		}

	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(124)
			p.Separator()
		}

	}

	return localctx
}

// IXsdNamespaceContext is an interface to support dynamic dispatch.
type IXsdNamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNs returns the ns token.
	GetNs() antlr.Token

	// SetNs sets the ns token.
	SetNs(antlr.Token)

	// IsXsdNamespaceContext differentiates from other interfaces.
	IsXsdNamespaceContext()
}

type XsdNamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ns     antlr.Token
}

func NewEmptyXsdNamespaceContext() *XsdNamespaceContext {
	var p = new(XsdNamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_xsdNamespace
	return p
}

func (*XsdNamespaceContext) IsXsdNamespaceContext() {}

func NewXsdNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XsdNamespaceContext {
	var p = new(XsdNamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_xsdNamespace

	return p
}

func (s *XsdNamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *XsdNamespaceContext) GetNs() antlr.Token { return s.ns }

func (s *XsdNamespaceContext) SetNs(v antlr.Token) { s.ns = v }

func (s *XsdNamespaceContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *XsdNamespaceContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *XsdNamespaceContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *XsdNamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XsdNamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XsdNamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterXsdNamespace(s)
	}
}

func (s *XsdNamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitXsdNamespace(s)
	}
}

func (p *AntlrThriftParser) XsdNamespace() (localctx IXsdNamespaceContext) {
	localctx = NewXsdNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AntlrThriftParserRULE_xsdNamespace)
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
		p.SetState(127)
		p.Match(AntlrThriftParserT__5)
	}
	{
		p.SetState(128)

		var _m = p.Match(AntlrThriftParserLITERAL)

		localctx.(*XsdNamespaceContext).ns = _m
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(129)
			p.AnnotationList()
		}

	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(132)
			p.Separator()
		}

	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) ConstDef() IConstDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *DefinitionContext) Typedef() ITypedefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedefContext)
}

func (s *DefinitionContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *DefinitionContext) Senum() ISenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISenumContext)
}

func (s *DefinitionContext) StructDef() IStructDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *DefinitionContext) UnionDef() IUnionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionDefContext)
}

func (s *DefinitionContext) ExceptionDef() IExceptionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExceptionDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExceptionDefContext)
}

func (s *DefinitionContext) ServiceDef() IServiceDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceDefContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *AntlrThriftParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AntlrThriftParserRULE_definition)

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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.ConstDef()
		}

	case AntlrThriftParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Typedef()
		}

	case AntlrThriftParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.EnumDef()
		}

	case AntlrThriftParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Senum()
		}

	case AntlrThriftParserT__16:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.StructDef()
		}

	case AntlrThriftParserT__17:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.UnionDef()
		}

	case AntlrThriftParserT__18:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)
			p.ExceptionDef()
		}

	case AntlrThriftParserT__19:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)
			p.ServiceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstDefContext is an interface to support dynamic dispatch.
type IConstDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstDefContext differentiates from other interfaces.
	IsConstDefContext()
}

type ConstDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDefContext() *ConstDefContext {
	var p = new(ConstDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_constDef
	return p
}

func (*ConstDefContext) IsConstDefContext() {}

func NewConstDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDefContext {
	var p = new(ConstDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_constDef

	return p
}

func (s *ConstDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDefContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *ConstDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *ConstDefContext) ConstValue() IConstValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstValueContext)
}

func (s *ConstDefContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *ConstDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterConstDef(s)
	}
}

func (s *ConstDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitConstDef(s)
	}
}

func (p *AntlrThriftParser) ConstDef() (localctx IConstDefContext) {
	localctx = NewConstDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AntlrThriftParserRULE_constDef)
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
		p.SetState(145)
		p.Match(AntlrThriftParserT__6)
	}
	{
		p.SetState(146)
		p.FieldType()
	}
	{
		p.SetState(147)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(148)
		p.Match(AntlrThriftParserT__7)
	}
	{
		p.SetState(149)
		p.ConstValue()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(150)
			p.Separator()
		}

	}

	return localctx
}

// IConstValueContext is an interface to support dynamic dispatch.
type IConstValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstValueContext differentiates from other interfaces.
	IsConstValueContext()
}

type ConstValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueContext() *ConstValueContext {
	var p = new(ConstValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_constValue
	return p
}

func (*ConstValueContext) IsConstValueContext() {}

func NewConstValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueContext {
	var p = new(ConstValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_constValue

	return p
}

func (s *ConstValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserINTEGER, 0)
}

func (s *ConstValueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserDOUBLE, 0)
}

func (s *ConstValueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *ConstValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *ConstValueContext) ConstList() IConstListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstListContext)
}

func (s *ConstValueContext) ConstMap() IConstMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstMapContext)
}

func (s *ConstValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterConstValue(s)
	}
}

func (s *ConstValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitConstValue(s)
	}
}

func (p *AntlrThriftParser) ConstValue() (localctx IConstValueContext) {
	localctx = NewConstValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AntlrThriftParserRULE_constValue)

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

	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(AntlrThriftParserINTEGER)
		}

	case AntlrThriftParserDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(AntlrThriftParserDOUBLE)
		}

	case AntlrThriftParserLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Match(AntlrThriftParserLITERAL)
		}

	case AntlrThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Match(AntlrThriftParserIDENTIFIER)
		}

	case AntlrThriftParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.ConstList()
		}

	case AntlrThriftParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(158)
			p.ConstMap()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstListContext is an interface to support dynamic dispatch.
type IConstListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstListContext differentiates from other interfaces.
	IsConstListContext()
}

type ConstListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstListContext() *ConstListContext {
	var p = new(ConstListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_constList
	return p
}

func (*ConstListContext) IsConstListContext() {}

func NewConstListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstListContext {
	var p = new(ConstListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_constList

	return p
}

func (s *ConstListContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstListContext) AllConstValue() []IConstValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstValueContext)(nil)).Elem())
	var tst = make([]IConstValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstValueContext)
		}
	}

	return tst
}

func (s *ConstListContext) ConstValue(i int) IConstValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstValueContext)
}

func (s *ConstListContext) AllSeparator() []ISeparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeparatorContext)(nil)).Elem())
	var tst = make([]ISeparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeparatorContext)
		}
	}

	return tst
}

func (s *ConstListContext) Separator(i int) ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *ConstListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterConstList(s)
	}
}

func (s *ConstListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitConstList(s)
	}
}

func (p *AntlrThriftParser) ConstList() (localctx IConstListContext) {
	localctx = NewConstListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AntlrThriftParserRULE_constList)
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
		p.SetState(161)
		p.Match(AntlrThriftParserT__8)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AntlrThriftParserT__8 || _la == AntlrThriftParserT__10 || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(AntlrThriftParserLITERAL-43))|(1<<(AntlrThriftParserIDENTIFIER-43))|(1<<(AntlrThriftParserINTEGER-43))|(1<<(AntlrThriftParserDOUBLE-43)))) != 0) {
		{
			p.SetState(162)
			p.ConstValue()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
			{
				p.SetState(163)
				p.Separator()
			}

		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
		p.Match(AntlrThriftParserT__9)
	}

	return localctx
}

// IConstMapContext is an interface to support dynamic dispatch.
type IConstMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstMapContext differentiates from other interfaces.
	IsConstMapContext()
}

type ConstMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstMapContext() *ConstMapContext {
	var p = new(ConstMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_constMap
	return p
}

func (*ConstMapContext) IsConstMapContext() {}

func NewConstMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstMapContext {
	var p = new(ConstMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_constMap

	return p
}

func (s *ConstMapContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstMapContext) AllConstMapEntry() []IConstMapEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstMapEntryContext)(nil)).Elem())
	var tst = make([]IConstMapEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstMapEntryContext)
		}
	}

	return tst
}

func (s *ConstMapContext) ConstMapEntry(i int) IConstMapEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstMapEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstMapEntryContext)
}

func (s *ConstMapContext) AllSeparator() []ISeparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeparatorContext)(nil)).Elem())
	var tst = make([]ISeparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeparatorContext)
		}
	}

	return tst
}

func (s *ConstMapContext) Separator(i int) ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *ConstMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterConstMap(s)
	}
}

func (s *ConstMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitConstMap(s)
	}
}

func (p *AntlrThriftParser) ConstMap() (localctx IConstMapContext) {
	localctx = NewConstMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AntlrThriftParserRULE_constMap)
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
		p.SetState(173)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AntlrThriftParserT__8 || _la == AntlrThriftParserT__10 || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(AntlrThriftParserLITERAL-43))|(1<<(AntlrThriftParserIDENTIFIER-43))|(1<<(AntlrThriftParserINTEGER-43))|(1<<(AntlrThriftParserDOUBLE-43)))) != 0) {
		{
			p.SetState(174)
			p.ConstMapEntry()
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
			{
				p.SetState(175)
				p.Separator()
			}

		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(AntlrThriftParserT__11)
	}

	return localctx
}

// IConstMapEntryContext is an interface to support dynamic dispatch.
type IConstMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IConstValueContext

	// GetValue returns the value rule contexts.
	GetValue() IConstValueContext

	// SetKey sets the key rule contexts.
	SetKey(IConstValueContext)

	// SetValue sets the value rule contexts.
	SetValue(IConstValueContext)

	// IsConstMapEntryContext differentiates from other interfaces.
	IsConstMapEntryContext()
}

type ConstMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IConstValueContext
	value  IConstValueContext
}

func NewEmptyConstMapEntryContext() *ConstMapEntryContext {
	var p = new(ConstMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_constMapEntry
	return p
}

func (*ConstMapEntryContext) IsConstMapEntryContext() {}

func NewConstMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstMapEntryContext {
	var p = new(ConstMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_constMapEntry

	return p
}

func (s *ConstMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstMapEntryContext) GetKey() IConstValueContext { return s.key }

func (s *ConstMapEntryContext) GetValue() IConstValueContext { return s.value }

func (s *ConstMapEntryContext) SetKey(v IConstValueContext) { s.key = v }

func (s *ConstMapEntryContext) SetValue(v IConstValueContext) { s.value = v }

func (s *ConstMapEntryContext) AllConstValue() []IConstValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstValueContext)(nil)).Elem())
	var tst = make([]IConstValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstValueContext)
		}
	}

	return tst
}

func (s *ConstMapEntryContext) ConstValue(i int) IConstValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstValueContext)
}

func (s *ConstMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterConstMapEntry(s)
	}
}

func (s *ConstMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitConstMapEntry(s)
	}
}

func (p *AntlrThriftParser) ConstMapEntry() (localctx IConstMapEntryContext) {
	localctx = NewConstMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AntlrThriftParserRULE_constMapEntry)

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

		var _x = p.ConstValue()

		localctx.(*ConstMapEntryContext).key = _x
	}
	{
		p.SetState(186)
		p.Match(AntlrThriftParserT__12)
	}
	{
		p.SetState(187)

		var _x = p.ConstValue()

		localctx.(*ConstMapEntryContext).value = _x
	}

	return localctx
}

// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_typedef
	return p
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_typedef

	return p
}

func (s *TypedefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TypedefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *TypedefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *TypedefContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *TypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterTypedef(s)
	}
}

func (s *TypedefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitTypedef(s)
	}
}

func (p *AntlrThriftParser) Typedef() (localctx ITypedefContext) {
	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AntlrThriftParserRULE_typedef)
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
		p.SetState(189)
		p.Match(AntlrThriftParserT__13)
	}
	{
		p.SetState(190)
		p.FieldType()
	}
	{
		p.SetState(191)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(192)
			p.AnnotationList()
		}

	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(195)
			p.Separator()
		}

	}

	return localctx
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_enumDef
	return p
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *EnumDefContext) AllEnumMember() []IEnumMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumMemberContext)(nil)).Elem())
	var tst = make([]IEnumMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumMemberContext)
		}
	}

	return tst
}

func (s *EnumDefContext) EnumMember(i int) IEnumMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumMemberContext)
}

func (s *EnumDefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *AntlrThriftParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AntlrThriftParserRULE_enumDef)
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
		p.SetState(198)
		p.Match(AntlrThriftParserT__14)
	}
	{
		p.SetState(199)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(200)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AntlrThriftParserIDENTIFIER {
		{
			p.SetState(201)
			p.EnumMember()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(AntlrThriftParserT__11)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(208)
			p.AnnotationList()
		}

	}

	return localctx
}

// IEnumMemberContext is an interface to support dynamic dispatch.
type IEnumMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumMemberContext differentiates from other interfaces.
	IsEnumMemberContext()
}

type EnumMemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumMemberContext() *EnumMemberContext {
	var p = new(EnumMemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_enumMember
	return p
}

func (*EnumMemberContext) IsEnumMemberContext() {}

func NewEnumMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberContext {
	var p = new(EnumMemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_enumMember

	return p
}

func (s *EnumMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumMemberContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *EnumMemberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserINTEGER, 0)
}

func (s *EnumMemberContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *EnumMemberContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *EnumMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterEnumMember(s)
	}
}

func (s *EnumMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitEnumMember(s)
	}
}

func (p *AntlrThriftParser) EnumMember() (localctx IEnumMemberContext) {
	localctx = NewEnumMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AntlrThriftParserRULE_enumMember)
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
		p.SetState(211)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__7 {
		{
			p.SetState(212)
			p.Match(AntlrThriftParserT__7)
		}
		{
			p.SetState(213)
			p.Match(AntlrThriftParserINTEGER)
		}

	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(216)
			p.AnnotationList()
		}

	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(219)
			p.Separator()
		}

	}

	return localctx
}

// ISenumContext is an interface to support dynamic dispatch.
type ISenumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSenumContext differentiates from other interfaces.
	IsSenumContext()
}

type SenumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySenumContext() *SenumContext {
	var p = new(SenumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_senum
	return p
}

func (*SenumContext) IsSenumContext() {}

func NewSenumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SenumContext {
	var p = new(SenumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_senum

	return p
}

func (s *SenumContext) GetParser() antlr.Parser { return s.parser }

func (s *SenumContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *SenumContext) AllEnumMember() []IEnumMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumMemberContext)(nil)).Elem())
	var tst = make([]IEnumMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumMemberContext)
		}
	}

	return tst
}

func (s *SenumContext) EnumMember(i int) IEnumMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumMemberContext)
}

func (s *SenumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SenumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SenumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterSenum(s)
	}
}

func (s *SenumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitSenum(s)
	}
}

func (p *AntlrThriftParser) Senum() (localctx ISenumContext) {
	localctx = NewSenumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AntlrThriftParserRULE_senum)
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
		p.SetState(222)
		p.Match(AntlrThriftParserT__15)
	}
	{
		p.SetState(223)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(224)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AntlrThriftParserIDENTIFIER {
		{
			p.SetState(225)
			p.EnumMember()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
		p.Match(AntlrThriftParserT__11)
	}

	return localctx
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_structDef
	return p
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *StructDefContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *StructDefContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StructDefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (p *AntlrThriftParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AntlrThriftParserRULE_structDef)
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
		p.SetState(233)
		p.Match(AntlrThriftParserT__16)
	}
	{
		p.SetState(234)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(235)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(AntlrThriftParserT__23-24))|(1<<(AntlrThriftParserT__24-24))|(1<<(AntlrThriftParserT__26-24))|(1<<(AntlrThriftParserT__27-24))|(1<<(AntlrThriftParserT__28-24))|(1<<(AntlrThriftParserT__29-24))|(1<<(AntlrThriftParserT__30-24))|(1<<(AntlrThriftParserT__31-24))|(1<<(AntlrThriftParserT__32-24))|(1<<(AntlrThriftParserT__33-24))|(1<<(AntlrThriftParserT__34-24))|(1<<(AntlrThriftParserT__35-24))|(1<<(AntlrThriftParserT__36-24))|(1<<(AntlrThriftParserT__39-24))|(1<<(AntlrThriftParserT__40-24))|(1<<(AntlrThriftParserIDENTIFIER-24))|(1<<(AntlrThriftParserINTEGER-24)))) != 0 {
		{
			p.SetState(236)
			p.Field()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
		p.Match(AntlrThriftParserT__11)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(243)
			p.AnnotationList()
		}

	}

	return localctx
}

// IUnionDefContext is an interface to support dynamic dispatch.
type IUnionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionDefContext differentiates from other interfaces.
	IsUnionDefContext()
}

type UnionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionDefContext() *UnionDefContext {
	var p = new(UnionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_unionDef
	return p
}

func (*UnionDefContext) IsUnionDefContext() {}

func NewUnionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionDefContext {
	var p = new(UnionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_unionDef

	return p
}

func (s *UnionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *UnionDefContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *UnionDefContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *UnionDefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *UnionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterUnionDef(s)
	}
}

func (s *UnionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitUnionDef(s)
	}
}

func (p *AntlrThriftParser) UnionDef() (localctx IUnionDefContext) {
	localctx = NewUnionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AntlrThriftParserRULE_unionDef)
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
		p.SetState(246)
		p.Match(AntlrThriftParserT__17)
	}
	{
		p.SetState(247)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(248)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(AntlrThriftParserT__23-24))|(1<<(AntlrThriftParserT__24-24))|(1<<(AntlrThriftParserT__26-24))|(1<<(AntlrThriftParserT__27-24))|(1<<(AntlrThriftParserT__28-24))|(1<<(AntlrThriftParserT__29-24))|(1<<(AntlrThriftParserT__30-24))|(1<<(AntlrThriftParserT__31-24))|(1<<(AntlrThriftParserT__32-24))|(1<<(AntlrThriftParserT__33-24))|(1<<(AntlrThriftParserT__34-24))|(1<<(AntlrThriftParserT__35-24))|(1<<(AntlrThriftParserT__36-24))|(1<<(AntlrThriftParserT__39-24))|(1<<(AntlrThriftParserT__40-24))|(1<<(AntlrThriftParserIDENTIFIER-24))|(1<<(AntlrThriftParserINTEGER-24)))) != 0 {
		{
			p.SetState(249)
			p.Field()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(255)
		p.Match(AntlrThriftParserT__11)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(256)
			p.AnnotationList()
		}

	}

	return localctx
}

// IExceptionDefContext is an interface to support dynamic dispatch.
type IExceptionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExceptionDefContext differentiates from other interfaces.
	IsExceptionDefContext()
}

type ExceptionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionDefContext() *ExceptionDefContext {
	var p = new(ExceptionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_exceptionDef
	return p
}

func (*ExceptionDefContext) IsExceptionDefContext() {}

func NewExceptionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionDefContext {
	var p = new(ExceptionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_exceptionDef

	return p
}

func (s *ExceptionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *ExceptionDefContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *ExceptionDefContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExceptionDefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *ExceptionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterExceptionDef(s)
	}
}

func (s *ExceptionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitExceptionDef(s)
	}
}

func (p *AntlrThriftParser) ExceptionDef() (localctx IExceptionDefContext) {
	localctx = NewExceptionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AntlrThriftParserRULE_exceptionDef)
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
		p.SetState(259)
		p.Match(AntlrThriftParserT__18)
	}
	{
		p.SetState(260)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(261)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(AntlrThriftParserT__23-24))|(1<<(AntlrThriftParserT__24-24))|(1<<(AntlrThriftParserT__26-24))|(1<<(AntlrThriftParserT__27-24))|(1<<(AntlrThriftParserT__28-24))|(1<<(AntlrThriftParserT__29-24))|(1<<(AntlrThriftParserT__30-24))|(1<<(AntlrThriftParserT__31-24))|(1<<(AntlrThriftParserT__32-24))|(1<<(AntlrThriftParserT__33-24))|(1<<(AntlrThriftParserT__34-24))|(1<<(AntlrThriftParserT__35-24))|(1<<(AntlrThriftParserT__36-24))|(1<<(AntlrThriftParserT__39-24))|(1<<(AntlrThriftParserT__40-24))|(1<<(AntlrThriftParserIDENTIFIER-24))|(1<<(AntlrThriftParserINTEGER-24)))) != 0 {
		{
			p.SetState(262)
			p.Field()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(268)
		p.Match(AntlrThriftParserT__11)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(269)
			p.AnnotationList()
		}

	}

	return localctx
}

// IServiceDefContext is an interface to support dynamic dispatch.
type IServiceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetSuperType returns the superType rule contexts.
	GetSuperType() IFieldTypeContext

	// SetSuperType sets the superType rule contexts.
	SetSuperType(IFieldTypeContext)

	// IsServiceDefContext differentiates from other interfaces.
	IsServiceDefContext()
}

type ServiceDefContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	name      antlr.Token
	superType IFieldTypeContext
}

func NewEmptyServiceDefContext() *ServiceDefContext {
	var p = new(ServiceDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_serviceDef
	return p
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) GetName() antlr.Token { return s.name }

func (s *ServiceDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ServiceDefContext) GetSuperType() IFieldTypeContext { return s.superType }

func (s *ServiceDefContext) SetSuperType(v IFieldTypeContext) { s.superType = v }

func (s *ServiceDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *ServiceDefContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *ServiceDefContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ServiceDefContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *ServiceDefContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *ServiceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterServiceDef(s)
	}
}

func (s *ServiceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitServiceDef(s)
	}
}

func (p *AntlrThriftParser) ServiceDef() (localctx IServiceDefContext) {
	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AntlrThriftParserRULE_serviceDef)
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
		p.SetState(272)
		p.Match(AntlrThriftParserT__19)
	}
	{
		p.SetState(273)

		var _m = p.Match(AntlrThriftParserIDENTIFIER)

		localctx.(*ServiceDefContext).name = _m
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__20 {
		{
			p.SetState(274)
			p.Match(AntlrThriftParserT__20)
		}
		{
			p.SetState(275)

			var _x = p.FieldType()

			localctx.(*ServiceDefContext).superType = _x
		}

	}
	{
		p.SetState(278)
		p.Match(AntlrThriftParserT__10)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AntlrThriftParserT__26-27))|(1<<(AntlrThriftParserT__27-27))|(1<<(AntlrThriftParserT__28-27))|(1<<(AntlrThriftParserT__29-27))|(1<<(AntlrThriftParserT__30-27))|(1<<(AntlrThriftParserT__31-27))|(1<<(AntlrThriftParserT__32-27))|(1<<(AntlrThriftParserT__33-27))|(1<<(AntlrThriftParserT__34-27))|(1<<(AntlrThriftParserT__35-27))|(1<<(AntlrThriftParserT__36-27))|(1<<(AntlrThriftParserT__39-27))|(1<<(AntlrThriftParserT__40-27))|(1<<(AntlrThriftParserVOID-27))|(1<<(AntlrThriftParserONEWAY-27))|(1<<(AntlrThriftParserIDENTIFIER-27)))) != 0 {
		{
			p.SetState(279)
			p.Function()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(285)
		p.Match(AntlrThriftParserT__11)
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(286)
			p.AnnotationList()
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *FunctionContext) FieldList() IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *FunctionContext) VOID() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserVOID, 0)
}

func (s *FunctionContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FunctionContext) ONEWAY() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserONEWAY, 0)
}

func (s *FunctionContext) ThrowsList() IThrowsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrowsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThrowsListContext)
}

func (s *FunctionContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *FunctionContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *AntlrThriftParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AntlrThriftParserRULE_function)
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
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserONEWAY {
		{
			p.SetState(289)
			p.Match(AntlrThriftParserONEWAY)
		}

	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserVOID:
		{
			p.SetState(292)
			p.Match(AntlrThriftParserVOID)
		}

	case AntlrThriftParserT__26, AntlrThriftParserT__27, AntlrThriftParserT__28, AntlrThriftParserT__29, AntlrThriftParserT__30, AntlrThriftParserT__31, AntlrThriftParserT__32, AntlrThriftParserT__33, AntlrThriftParserT__34, AntlrThriftParserT__35, AntlrThriftParserT__36, AntlrThriftParserT__39, AntlrThriftParserT__40, AntlrThriftParserIDENTIFIER:
		{
			p.SetState(293)
			p.FieldType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(296)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	{
		p.SetState(297)
		p.FieldList()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__25 {
		{
			p.SetState(298)
			p.ThrowsList()
		}

	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(301)
			p.AnnotationList()
		}

	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(304)
			p.Separator()
		}

	}

	return localctx
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_fieldList
	return p
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FieldListContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (p *AntlrThriftParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AntlrThriftParserRULE_fieldList)
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
		p.SetState(307)
		p.Match(AntlrThriftParserT__21)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(AntlrThriftParserT__23-24))|(1<<(AntlrThriftParserT__24-24))|(1<<(AntlrThriftParserT__26-24))|(1<<(AntlrThriftParserT__27-24))|(1<<(AntlrThriftParserT__28-24))|(1<<(AntlrThriftParserT__29-24))|(1<<(AntlrThriftParserT__30-24))|(1<<(AntlrThriftParserT__31-24))|(1<<(AntlrThriftParserT__32-24))|(1<<(AntlrThriftParserT__33-24))|(1<<(AntlrThriftParserT__34-24))|(1<<(AntlrThriftParserT__35-24))|(1<<(AntlrThriftParserT__36-24))|(1<<(AntlrThriftParserT__39-24))|(1<<(AntlrThriftParserT__40-24))|(1<<(AntlrThriftParserIDENTIFIER-24))|(1<<(AntlrThriftParserINTEGER-24)))) != 0 {
		{
			p.SetState(308)
			p.Field()
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(314)
		p.Match(AntlrThriftParserT__22)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *FieldContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserINTEGER, 0)
}

func (s *FieldContext) Requiredness() IRequirednessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequirednessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequirednessContext)
}

func (s *FieldContext) ConstValue() IConstValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstValueContext)
}

func (s *FieldContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *FieldContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *AntlrThriftParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AntlrThriftParserRULE_field)
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
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserINTEGER {
		{
			p.SetState(316)
			p.Match(AntlrThriftParserINTEGER)
		}
		{
			p.SetState(317)
			p.Match(AntlrThriftParserT__12)
		}

	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__23 || _la == AntlrThriftParserT__24 {
		{
			p.SetState(320)
			p.Requiredness()
		}

	}
	{
		p.SetState(323)
		p.FieldType()
	}
	{
		p.SetState(324)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__7 {
		{
			p.SetState(325)
			p.Match(AntlrThriftParserT__7)
		}
		{
			p.SetState(326)
			p.ConstValue()
		}

	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__21 {
		{
			p.SetState(329)
			p.AnnotationList()
		}

	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(332)
			p.Separator()
		}

	}

	return localctx
}

// IRequirednessContext is an interface to support dynamic dispatch.
type IRequirednessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequirednessContext differentiates from other interfaces.
	IsRequirednessContext()
}

type RequirednessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequirednessContext() *RequirednessContext {
	var p = new(RequirednessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_requiredness
	return p
}

func (*RequirednessContext) IsRequirednessContext() {}

func NewRequirednessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequirednessContext {
	var p = new(RequirednessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_requiredness

	return p
}

func (s *RequirednessContext) GetParser() antlr.Parser { return s.parser }
func (s *RequirednessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequirednessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequirednessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterRequiredness(s)
	}
}

func (s *RequirednessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitRequiredness(s)
	}
}

func (p *AntlrThriftParser) Requiredness() (localctx IRequirednessContext) {
	localctx = NewRequirednessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AntlrThriftParserRULE_requiredness)
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
	p.SetState(335)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AntlrThriftParserT__23 || _la == AntlrThriftParserT__24) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IThrowsListContext is an interface to support dynamic dispatch.
type IThrowsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrowsListContext differentiates from other interfaces.
	IsThrowsListContext()
}

type ThrowsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowsListContext() *ThrowsListContext {
	var p = new(ThrowsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_throwsList
	return p
}

func (*ThrowsListContext) IsThrowsListContext() {}

func NewThrowsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowsListContext {
	var p = new(ThrowsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_throwsList

	return p
}

func (s *ThrowsListContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowsListContext) FieldList() IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *ThrowsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterThrowsList(s)
	}
}

func (s *ThrowsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitThrowsList(s)
	}
}

func (p *AntlrThriftParser) ThrowsList() (localctx IThrowsListContext) {
	localctx = NewThrowsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AntlrThriftParserRULE_throwsList)

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
		p.SetState(337)
		p.Match(AntlrThriftParserT__25)
	}
	{
		p.SetState(338)
		p.FieldList()
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *FieldTypeContext) AnnotationList() IAnnotationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationListContext)
}

func (s *FieldTypeContext) ContainerType() IContainerTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainerTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainerTypeContext)
}

func (s *FieldTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *AntlrThriftParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AntlrThriftParserRULE_fieldType)
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

	p.SetState(352)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserT__26, AntlrThriftParserT__27, AntlrThriftParserT__28, AntlrThriftParserT__29, AntlrThriftParserT__30, AntlrThriftParserT__31, AntlrThriftParserT__32, AntlrThriftParserT__33, AntlrThriftParserT__34, AntlrThriftParserT__35:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.BaseType()
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AntlrThriftParserT__21 {
			{
				p.SetState(341)
				p.AnnotationList()
			}

		}

	case AntlrThriftParserT__36, AntlrThriftParserT__39, AntlrThriftParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.ContainerType()
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AntlrThriftParserT__21 {
			{
				p.SetState(345)
				p.AnnotationList()
			}

		}

	case AntlrThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.Match(AntlrThriftParserIDENTIFIER)
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AntlrThriftParserT__21 {
			{
				p.SetState(349)
				p.AnnotationList()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (p *AntlrThriftParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AntlrThriftParserRULE_baseType)
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
	p.SetState(354)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AntlrThriftParserT__26-27))|(1<<(AntlrThriftParserT__27-27))|(1<<(AntlrThriftParserT__28-27))|(1<<(AntlrThriftParserT__29-27))|(1<<(AntlrThriftParserT__30-27))|(1<<(AntlrThriftParserT__31-27))|(1<<(AntlrThriftParserT__32-27))|(1<<(AntlrThriftParserT__33-27))|(1<<(AntlrThriftParserT__34-27))|(1<<(AntlrThriftParserT__35-27)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IContainerTypeContext is an interface to support dynamic dispatch.
type IContainerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainerTypeContext differentiates from other interfaces.
	IsContainerTypeContext()
}

type ContainerTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainerTypeContext() *ContainerTypeContext {
	var p = new(ContainerTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_containerType
	return p
}

func (*ContainerTypeContext) IsContainerTypeContext() {}

func NewContainerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainerTypeContext {
	var p = new(ContainerTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_containerType

	return p
}

func (s *ContainerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainerTypeContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *ContainerTypeContext) SetType() ISetTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetTypeContext)
}

func (s *ContainerTypeContext) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *ContainerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterContainerType(s)
	}
}

func (s *ContainerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitContainerType(s)
	}
}

func (p *AntlrThriftParser) ContainerType() (localctx IContainerTypeContext) {
	localctx = NewContainerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AntlrThriftParserRULE_containerType)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AntlrThriftParserT__36:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.MapType()
		}

	case AntlrThriftParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.SetType()
		}

	case AntlrThriftParserT__39:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.ListType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IFieldTypeContext

	// GetValue returns the value rule contexts.
	GetValue() IFieldTypeContext

	// SetKey sets the key rule contexts.
	SetKey(IFieldTypeContext)

	// SetValue sets the value rule contexts.
	SetValue(IFieldTypeContext)

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IFieldTypeContext
	value  IFieldTypeContext
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) GetKey() IFieldTypeContext { return s.key }

func (s *MapTypeContext) GetValue() IFieldTypeContext { return s.value }

func (s *MapTypeContext) SetKey(v IFieldTypeContext) { s.key = v }

func (s *MapTypeContext) SetValue(v IFieldTypeContext) { s.value = v }

func (s *MapTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserCOMMA, 0)
}

func (s *MapTypeContext) AllFieldType() []IFieldTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem())
	var tst = make([]IFieldTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldTypeContext)
		}
	}

	return tst
}

func (s *MapTypeContext) FieldType(i int) IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *MapTypeContext) CppType() ICppTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICppTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICppTypeContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (p *AntlrThriftParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AntlrThriftParserRULE_mapType)
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
		p.SetState(361)
		p.Match(AntlrThriftParserT__36)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__41 {
		{
			p.SetState(362)
			p.CppType()
		}

	}
	{
		p.SetState(365)
		p.Match(AntlrThriftParserT__37)
	}
	{
		p.SetState(366)

		var _x = p.FieldType()

		localctx.(*MapTypeContext).key = _x
	}
	{
		p.SetState(367)
		p.Match(AntlrThriftParserCOMMA)
	}
	{
		p.SetState(368)

		var _x = p.FieldType()

		localctx.(*MapTypeContext).value = _x
	}
	{
		p.SetState(369)
		p.Match(AntlrThriftParserT__38)
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *ListTypeContext) CppType() ICppTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICppTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICppTypeContext)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *AntlrThriftParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AntlrThriftParserRULE_listType)
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
		p.SetState(371)
		p.Match(AntlrThriftParserT__39)
	}
	{
		p.SetState(372)
		p.Match(AntlrThriftParserT__37)
	}
	{
		p.SetState(373)
		p.FieldType()
	}
	{
		p.SetState(374)
		p.Match(AntlrThriftParserT__38)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__41 {
		{
			p.SetState(375)
			p.CppType()
		}

	}

	return localctx
}

// ISetTypeContext is an interface to support dynamic dispatch.
type ISetTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetTypeContext differentiates from other interfaces.
	IsSetTypeContext()
}

type SetTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetTypeContext() *SetTypeContext {
	var p = new(SetTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_setType
	return p
}

func (*SetTypeContext) IsSetTypeContext() {}

func NewSetTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetTypeContext {
	var p = new(SetTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_setType

	return p
}

func (s *SetTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SetTypeContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *SetTypeContext) CppType() ICppTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICppTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICppTypeContext)
}

func (s *SetTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterSetType(s)
	}
}

func (s *SetTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitSetType(s)
	}
}

func (p *AntlrThriftParser) SetType() (localctx ISetTypeContext) {
	localctx = NewSetTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AntlrThriftParserRULE_setType)
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
		p.SetState(378)
		p.Match(AntlrThriftParserT__40)
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__41 {
		{
			p.SetState(379)
			p.CppType()
		}

	}
	{
		p.SetState(382)
		p.Match(AntlrThriftParserT__37)
	}
	{
		p.SetState(383)
		p.FieldType()
	}
	{
		p.SetState(384)
		p.Match(AntlrThriftParserT__38)
	}

	return localctx
}

// ICppTypeContext is an interface to support dynamic dispatch.
type ICppTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCppTypeContext differentiates from other interfaces.
	IsCppTypeContext()
}

type CppTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCppTypeContext() *CppTypeContext {
	var p = new(CppTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_cppType
	return p
}

func (*CppTypeContext) IsCppTypeContext() {}

func NewCppTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CppTypeContext {
	var p = new(CppTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_cppType

	return p
}

func (s *CppTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CppTypeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *CppTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CppTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CppTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterCppType(s)
	}
}

func (s *CppTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitCppType(s)
	}
}

func (p *AntlrThriftParser) CppType() (localctx ICppTypeContext) {
	localctx = NewCppTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AntlrThriftParserRULE_cppType)

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
		p.SetState(386)
		p.Match(AntlrThriftParserT__41)
	}
	{
		p.SetState(387)
		p.Match(AntlrThriftParserLITERAL)
	}

	return localctx
}

// IAnnotationListContext is an interface to support dynamic dispatch.
type IAnnotationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationListContext differentiates from other interfaces.
	IsAnnotationListContext()
}

type AnnotationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationListContext() *AnnotationListContext {
	var p = new(AnnotationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_annotationList
	return p
}

func (*AnnotationListContext) IsAnnotationListContext() {}

func NewAnnotationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationListContext {
	var p = new(AnnotationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_annotationList

	return p
}

func (s *AnnotationListContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationListContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *AnnotationListContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *AnnotationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterAnnotationList(s)
	}
}

func (s *AnnotationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitAnnotationList(s)
	}
}

func (p *AntlrThriftParser) AnnotationList() (localctx IAnnotationListContext) {
	localctx = NewAnnotationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, AntlrThriftParserRULE_annotationList)
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
		p.SetState(389)
		p.Match(AntlrThriftParserT__21)
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AntlrThriftParserIDENTIFIER {
		{
			p.SetState(390)
			p.Annotation()
		}

		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(396)
		p.Match(AntlrThriftParserT__22)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserIDENTIFIER, 0)
}

func (s *AnnotationContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserLITERAL, 0)
}

func (s *AnnotationContext) Separator() ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *AntlrThriftParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, AntlrThriftParserRULE_annotation)
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
		p.SetState(398)
		p.Match(AntlrThriftParserIDENTIFIER)
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserT__7 {
		{
			p.SetState(399)
			p.Match(AntlrThriftParserT__7)
		}
		{
			p.SetState(400)
			p.Match(AntlrThriftParserLITERAL)
		}

	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON {
		{
			p.SetState(403)
			p.Separator()
		}

	}

	return localctx
}

// ISeparatorContext is an interface to support dynamic dispatch.
type ISeparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeparatorContext differentiates from other interfaces.
	IsSeparatorContext()
}

type SeparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeparatorContext() *SeparatorContext {
	var p = new(SeparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AntlrThriftParserRULE_separator
	return p
}

func (*SeparatorContext) IsSeparatorContext() {}

func NewSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeparatorContext {
	var p = new(SeparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AntlrThriftParserRULE_separator

	return p
}

func (s *SeparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *SeparatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserCOMMA, 0)
}

func (s *SeparatorContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(AntlrThriftParserSEMICOLON, 0)
}

func (s *SeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.EnterSeparator(s)
	}
}

func (s *SeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AntlrThriftListener); ok {
		listenerT.ExitSeparator(s)
	}
}

func (p *AntlrThriftParser) Separator() (localctx ISeparatorContext) {
	localctx = NewSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, AntlrThriftParserRULE_separator)
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
	p.SetState(406)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AntlrThriftParserCOMMA || _la == AntlrThriftParserSEMICOLON) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
