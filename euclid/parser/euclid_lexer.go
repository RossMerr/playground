// Code generated from Euclid.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 50, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 6, 7, 35, 10, 7, 13, 7, 14,
	7, 36, 3, 8, 6, 8, 40, 10, 8, 13, 8, 14, 8, 41, 3, 8, 3, 8, 3, 9, 6, 9,
	47, 10, 9, 13, 9, 14, 9, 48, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 3, 2, 5, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34,
	3, 2, 99, 124, 2, 52, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3,
	2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 34, 3, 2, 2, 2, 15,
	39, 3, 2, 2, 2, 17, 46, 3, 2, 2, 2, 19, 20, 7, 106, 2, 2, 20, 21, 7, 103,
	2, 2, 21, 22, 7, 110, 2, 2, 22, 23, 7, 110, 2, 2, 23, 24, 7, 113, 2, 2,
	24, 4, 3, 2, 2, 2, 25, 26, 7, 44, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 49,
	2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 45, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32,
	7, 47, 2, 2, 32, 12, 3, 2, 2, 2, 33, 35, 9, 2, 2, 2, 34, 33, 3, 2, 2, 2,
	35, 36, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 14, 3,
	2, 2, 2, 38, 40, 9, 3, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41,
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 8, 8, 2,
	2, 44, 16, 3, 2, 2, 2, 45, 47, 9, 4, 2, 2, 46, 45, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 18, 3, 2, 2, 2,
	6, 2, 36, 41, 48, 3, 8, 2, 2,
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
	"", "'hello'", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WS", "ID",
}

var lexerRuleNames = []string{
	"T__0", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WS", "ID",
}

type EuclidLexer struct {
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

func NewEuclidLexer(input antlr.CharStream) *EuclidLexer {

	l := new(EuclidLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Euclid.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EuclidLexer tokens.
const (
	EuclidLexerT__0   = 1
	EuclidLexerMUL    = 2
	EuclidLexerDIV    = 3
	EuclidLexerADD    = 4
	EuclidLexerSUB    = 5
	EuclidLexerNUMBER = 6
	EuclidLexerWS     = 7
	EuclidLexerID     = 8
)
