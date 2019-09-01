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

// EnterImportR is called when production importR is entered.
func (s *BaseEuclidListener) EnterImportR(ctx *ImportRContext) {}

// ExitImportR is called when production importR is exited.
func (s *BaseEuclidListener) ExitImportR(ctx *ImportRContext) {}
