package euclid

import (
	"github.com/RossMerr/playground/euclid/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type euclidListener struct {
	*parser.BaseEuclidListener

	stack []int
}

// Parser takes a string expression and returns the evaluated result.
func Parser(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewEuclidLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewEuclidParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	// // Finally parse the expression (by walking the tree)
	// var listener euclidListener
	// antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// return listener.pop()
}
