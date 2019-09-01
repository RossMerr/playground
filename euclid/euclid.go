package euclid

import (
	"github.com/RossMerr/playground/euclid/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parser takes a string expression and returns the evaluated result.
func Parser(input string) (error, string)  {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewEuclidLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewEuclidParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//tree := p.R()
	tree := p.Start()
	// Finally parse the expression (by walking the tree)
	var walker euclidWalker
	antlr.ParseTreeWalkerDefault.Walk(&walker, tree)

	return walker.ErrorReport(), walker.StackDump()
}