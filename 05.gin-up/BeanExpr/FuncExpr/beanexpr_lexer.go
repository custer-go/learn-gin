// Code generated from /Users/tianxiaoqiang/Work/2020/study/learn-gin/05.gin-up/BeanExpr/BeanExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package FuncExpr
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 88, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 
	6, 3, 6, 7, 6, 40, 10, 6, 12, 6, 14, 6, 43, 11, 6, 3, 6, 3, 6, 3, 7, 5, 
	7, 48, 10, 7, 3, 7, 3, 7, 3, 8, 5, 8, 53, 10, 8, 3, 8, 6, 8, 56, 10, 8, 
	13, 8, 14, 8, 57, 3, 9, 3, 9, 7, 9, 62, 10, 9, 12, 9, 14, 9, 65, 11, 9, 
	3, 10, 3, 10, 3, 11, 6, 11, 70, 10, 11, 13, 11, 14, 11, 71, 5, 11, 74, 
	10, 11, 3, 11, 3, 11, 6, 11, 78, 10, 11, 13, 11, 14, 11, 79, 3, 12, 6, 
	12, 83, 10, 12, 13, 12, 14, 12, 84, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 
	7, 5, 9, 2, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 3, 2, 7, 
	3, 2, 50, 59, 4, 2, 41, 41, 94, 94, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 
	67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 97, 2, 3, 3, 2, 2, 2, 
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 
	2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 
	2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33, 3, 2, 2, 2, 13, 47, 3, 2, 2, 2, 15, 
	52, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19, 66, 3, 2, 2, 2, 21, 73, 3, 2, 2, 
	2, 23, 82, 3, 2, 2, 2, 25, 26, 7, 42, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 
	7, 43, 2, 2, 28, 6, 3, 2, 2, 2, 29, 30, 7, 46, 2, 2, 30, 8, 3, 2, 2, 2, 
	31, 32, 9, 2, 2, 2, 32, 10, 3, 2, 2, 2, 33, 41, 7, 41, 2, 2, 34, 35, 7, 
	94, 2, 2, 35, 40, 11, 2, 2, 2, 36, 37, 7, 41, 2, 2, 37, 40, 7, 41, 2, 2, 
	38, 40, 10, 3, 2, 2, 39, 34, 3, 2, 2, 2, 39, 36, 3, 2, 2, 2, 39, 38, 3, 
	2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 
	44, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 41, 2, 2, 45, 12, 3, 2, 
	2, 2, 46, 48, 7, 47, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 
	49, 3, 2, 2, 2, 49, 50, 5, 21, 11, 2, 50, 14, 3, 2, 2, 2, 51, 53, 7, 47, 
	2, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 56, 
	5, 9, 5, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 
	57, 58, 3, 2, 2, 2, 58, 16, 3, 2, 2, 2, 59, 63, 9, 4, 2, 2, 60, 62, 9, 
	5, 2, 2, 61, 60, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 
	64, 3, 2, 2, 2, 64, 18, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 48, 
	2, 2, 67, 20, 3, 2, 2, 2, 68, 70, 5, 9, 5, 2, 69, 68, 3, 2, 2, 2, 70, 71, 
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 
	73, 69, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 7, 
	48, 2, 2, 76, 78, 5, 9, 5, 2, 77, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 
	77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 22, 3, 2, 2, 2, 81, 83, 9, 6, 2, 
	2, 82, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 
	3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 8, 12, 2, 2, 87, 24, 3, 2, 2, 2, 
	13, 2, 39, 41, 47, 52, 57, 63, 71, 73, 79, 84, 3, 8, 2, 2,
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
	"", "'('", "')'", "','", "", "", "", "", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "StringArg", "FloatArg", "IntArg", "FuncName", "Dot", "Float", 
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "DIGIT", "StringArg", "FloatArg", "IntArg", "FuncName", 
	"Dot", "Float", "WHITESPACE",
}

type BeanExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBeanExprLexer(input antlr.CharStream) *BeanExprLexer {

	l := new(BeanExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeanExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeanExprLexer tokens.
const (
	BeanExprLexerT__0 = 1
	BeanExprLexerT__1 = 2
	BeanExprLexerT__2 = 3
	BeanExprLexerStringArg = 4
	BeanExprLexerFloatArg = 5
	BeanExprLexerIntArg = 6
	BeanExprLexerFuncName = 7
	BeanExprLexerDot = 8
	BeanExprLexerFloat = 9
	BeanExprLexerWHITESPACE = 10
)

