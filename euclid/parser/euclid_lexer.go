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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 26, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 6, 3, 18, 10, 3, 13, 3, 14, 3, 19, 3, 4, 6, 4, 23, 10, 4,
	13, 4, 14, 4, 24, 2, 2, 5, 3, 3, 5, 4, 7, 5, 3, 2, 4, 3, 2, 102, 102, 3,
	2, 121, 121, 2, 27, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 3, 9, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 22, 3, 2, 2, 2, 9, 10, 7, 107,
	2, 2, 10, 11, 7, 111, 2, 2, 11, 12, 7, 114, 2, 2, 12, 13, 7, 113, 2, 2,
	13, 14, 7, 116, 2, 2, 14, 15, 7, 118, 2, 2, 15, 4, 3, 2, 2, 2, 16, 18,
	9, 2, 2, 2, 17, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2,
	19, 20, 3, 2, 2, 2, 20, 6, 3, 2, 2, 2, 21, 23, 9, 3, 2, 2, 22, 21, 3, 2,
	2, 2, 23, 24, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 8,
	3, 2, 2, 2, 5, 2, 19, 24, 2,
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
	"", "'import'",
}

var lexerSymbolicNames = []string{
	"", "IMPORT", "DIGIT", "WORD",
}

var lexerRuleNames = []string{
	"IMPORT", "DIGIT", "WORD",
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
	EuclidLexerIMPORT = 1
	EuclidLexerDIGIT  = 2
	EuclidLexerWORD   = 3
)