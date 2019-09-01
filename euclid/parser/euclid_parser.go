// Code generated from Euclid.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Euclid

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 7, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 5, 2, 4, 3, 2, 2, 2, 4,
	5, 7, 3, 2, 2, 5, 3, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'import'",
}
var symbolicNames = []string{
	"", "IMPORT", "DIGIT", "WORD",
}

var ruleNames = []string{
	"importR",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type EuclidParser struct {
	*antlr.BaseParser
}

func NewEuclidParser(input antlr.TokenStream) *EuclidParser {
	this := new(EuclidParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Euclid.g4"

	return this
}

// EuclidParser tokens.
const (
	EuclidParserEOF    = antlr.TokenEOF
	EuclidParserIMPORT = 1
	EuclidParserDIGIT  = 2
	EuclidParserWORD   = 3
)

// EuclidParserRULE_importR is the EuclidParser rule.
const EuclidParserRULE_importR = 0

// IImportRContext is an interface to support dynamic dispatch.
type IImportRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportRContext differentiates from other interfaces.
	IsImportRContext()
}

type ImportRContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportRContext() *ImportRContext {
	var p = new(ImportRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EuclidParserRULE_importR
	return p
}

func (*ImportRContext) IsImportRContext() {}

func NewImportRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportRContext {
	var p = new(ImportRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EuclidParserRULE_importR

	return p
}

func (s *ImportRContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportRContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(EuclidParserIMPORT, 0)
}

func (s *ImportRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.EnterImportR(s)
	}
}

func (s *ImportRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.ExitImportR(s)
	}
}

func (p *EuclidParser) ImportR() (localctx IImportRContext) {
	localctx = NewImportRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EuclidParserRULE_importR)

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
		p.SetState(2)
		p.Match(EuclidParserIMPORT)
	}

	return localctx
}
