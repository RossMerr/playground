package euclid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RossMerr/playground/euclid/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type euclidWalker struct {
	*parser.BaseEuclidListener
	errors []antlr.ErrorNode
	stack  []interface{}
}

func (l *euclidWalker) VisitErrorNode(node antlr.ErrorNode) {
	l.errors = append(l.errors, node)
}

func (l *euclidWalker) ErrorReport() error {
	var errstrings []string

	for _, e := range l.errors {
		errstrings = append(errstrings, e.GetText())
	}

	if len(errstrings) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errstrings, "\n"))
}

func (l *euclidWalker) StackDump() string {
	var stackStrings []string

	for _, e := range l.stack {
		stackStrings = append(stackStrings, fmt.Sprintf("%+v", e))
	}

	return strings.Join(stackStrings, "\n")
}

func (l *euclidWalker) EnterR(ctx *parser.RContext) {
	fmt.Printf("Entering R : %s\n", ctx.ID().GetText())
}

func (l *euclidWalker) ExitR(ctx *parser.RContext) {
	fmt.Printf("Exiting R\n")
}

func (l *euclidWalker) EnterOperation(ctx *parser.OperationContext) {

	operator := ctx.GetOperator().GetText()

	fmt.Printf("Entering Operation : %v\n", operator)

	left, _ := strconv.ParseFloat(ctx.GetLeft().GetText(), 64)
	right, _ := strconv.ParseFloat(ctx.GetRight().GetText(), 64)

	var result float64
	switch ctx.GetOperator().GetTokenType() {
	case parser.EuclidParserADD:
		result = left + right
	case parser.EuclidParserSUB:
		result = left - right
	case parser.EuclidParserDIV:
		result = left / right
	case parser.EuclidParserMUL:
		result = left * right
	default:
		panic(fmt.Sprintf("Unexpected Operation: %s", operator))
	}

	l.stack = append(l.stack, result)

}
