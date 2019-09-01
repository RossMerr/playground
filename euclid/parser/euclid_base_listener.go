// Code generated from Euclid.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Euclid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEuclidListener is a complete listener for a parse tree produced by EuclidParser.
type BaseEuclidListener struct{}

var _ EuclidListener = &BaseEuclidListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEuclidListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEuclidListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEuclidListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEuclidListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseEuclidListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseEuclidListener) ExitOperation(ctx *OperationContext) {}

// EnterR is called when production r is entered.
func (s *BaseEuclidListener) EnterR(ctx *RContext) {}

// ExitR is called when production r is exited.
func (s *BaseEuclidListener) ExitR(ctx *RContext) {}

// EnterStart is called when production start is entered.
func (s *BaseEuclidListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseEuclidListener) ExitStart(ctx *StartContext) {}
