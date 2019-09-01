// Code generated from Euclid.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Euclid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EuclidListener is a complete listener for a parse tree produced by EuclidParser.
type EuclidListener interface {
	antlr.ParseTreeListener

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterR is called when entering the r production.
	EnterR(c *RContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitR is called when exiting the r production.
	ExitR(c *RContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)
}
