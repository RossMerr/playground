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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 21, 10, 2, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 30, 10, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2,
	32, 2, 20, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 9, 7, 8,
	2, 2, 9, 10, 7, 6, 2, 2, 10, 21, 7, 8, 2, 2, 11, 12, 7, 8, 2, 2, 12, 13,
	7, 7, 2, 2, 13, 21, 7, 8, 2, 2, 14, 15, 7, 8, 2, 2, 15, 16, 7, 4, 2, 2,
	16, 21, 7, 8, 2, 2, 17, 18, 7, 8, 2, 2, 18, 19, 7, 5, 2, 2, 19, 21, 7,
	8, 2, 2, 20, 8, 3, 2, 2, 2, 20, 11, 3, 2, 2, 2, 20, 14, 3, 2, 2, 2, 20,
	17, 3, 2, 2, 2, 21, 3, 3, 2, 2, 2, 22, 23, 7, 3, 2, 2, 23, 24, 7, 10, 2,
	2, 24, 5, 3, 2, 2, 2, 25, 30, 5, 4, 3, 2, 26, 27, 5, 2, 2, 2, 27, 28, 7,
	2, 2, 3, 28, 30, 3, 2, 2, 2, 29, 25, 3, 2, 2, 2, 29, 26, 3, 2, 2, 2, 30,
	7, 3, 2, 2, 2, 4, 20, 29,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'hello'", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WS", "ID",
}

var ruleNames = []string{
	"operation", "r", "start",
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
	EuclidParserT__0   = 1
	EuclidParserMUL    = 2
	EuclidParserDIV    = 3
	EuclidParserADD    = 4
	EuclidParserSUB    = 5
	EuclidParserNUMBER = 6
	EuclidParserWS     = 7
	EuclidParserID     = 8
)

// EuclidParser rules.
const (
	EuclidParserRULE_operation = 0
	EuclidParserRULE_r         = 1
	EuclidParserRULE_start     = 2
)

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left token.
	GetLeft() antlr.Token

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// GetRight returns the right token.
	GetRight() antlr.Token

	// SetLeft sets the left token.
	SetLeft(antlr.Token)

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// SetRight sets the right token.
	SetRight(antlr.Token)

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     antlr.Token
	operator antlr.Token
	right    antlr.Token
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EuclidParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EuclidParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) GetLeft() antlr.Token { return s.left }

func (s *OperationContext) GetOperator() antlr.Token { return s.operator }

func (s *OperationContext) GetRight() antlr.Token { return s.right }

func (s *OperationContext) SetLeft(v antlr.Token) { s.left = v }

func (s *OperationContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *OperationContext) SetRight(v antlr.Token) { s.right = v }

func (s *OperationContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(EuclidParserNUMBER)
}

func (s *OperationContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(EuclidParserNUMBER, i)
}

func (s *OperationContext) ADD() antlr.TerminalNode {
	return s.GetToken(EuclidParserADD, 0)
}

func (s *OperationContext) SUB() antlr.TerminalNode {
	return s.GetToken(EuclidParserSUB, 0)
}

func (s *OperationContext) MUL() antlr.TerminalNode {
	return s.GetToken(EuclidParserMUL, 0)
}

func (s *OperationContext) DIV() antlr.TerminalNode {
	return s.GetToken(EuclidParserDIV, 0)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *EuclidParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EuclidParserRULE_operation)

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

	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).left = _m
		}
		{
			p.SetState(7)

			var _m = p.Match(EuclidParserADD)

			localctx.(*OperationContext).operator = _m
		}
		{
			p.SetState(8)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).right = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).left = _m
		}
		{
			p.SetState(10)

			var _m = p.Match(EuclidParserSUB)

			localctx.(*OperationContext).operator = _m
		}
		{
			p.SetState(11)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).right = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).left = _m
		}
		{
			p.SetState(13)

			var _m = p.Match(EuclidParserMUL)

			localctx.(*OperationContext).operator = _m
		}
		{
			p.SetState(14)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).right = _m
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(15)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).left = _m
		}
		{
			p.SetState(16)

			var _m = p.Match(EuclidParserDIV)

			localctx.(*OperationContext).operator = _m
		}
		{
			p.SetState(17)

			var _m = p.Match(EuclidParserNUMBER)

			localctx.(*OperationContext).right = _m
		}

	}

	return localctx
}

// IRContext is an interface to support dynamic dispatch.
type IRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRContext differentiates from other interfaces.
	IsRContext()
}

type RContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRContext() *RContext {
	var p = new(RContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EuclidParserRULE_r
	return p
}

func (*RContext) IsRContext() {}

func NewRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RContext {
	var p = new(RContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EuclidParserRULE_r

	return p
}

func (s *RContext) GetParser() antlr.Parser { return s.parser }

func (s *RContext) ID() antlr.TerminalNode {
	return s.GetToken(EuclidParserID, 0)
}

func (s *RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.EnterR(s)
	}
}

func (s *RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.ExitR(s)
	}
}

func (p *EuclidParser) R() (localctx IRContext) {
	localctx = NewRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EuclidParserRULE_r)

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
		p.SetState(20)
		p.Match(EuclidParserT__0)
	}
	{
		p.SetState(21)
		p.Match(EuclidParserID)
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EuclidParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EuclidParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) R() IRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRContext)
}

func (s *StartContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(EuclidParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EuclidListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *EuclidParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EuclidParserRULE_start)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EuclidParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.R()
		}

	case EuclidParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Operation()
		}
		{
			p.SetState(25)
			p.Match(EuclidParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
