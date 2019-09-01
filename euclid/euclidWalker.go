package euclid

import (
	"fmt"
	"strings"

	"github.com/RossMerr/playground/euclid/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type euclidWalker struct {
	*parser.BaseEuclidListener

	stack []int

	errors []antlr.ErrorNode
}

func (l *euclidWalker) EnterR(ctx *parser.RContext) {
	fmt.Printf("Entering R : %s\n", ctx.ID().GetText())
}

func (l *euclidWalker) ExitR(ctx *parser.RContext) {
	fmt.Printf("Exiting R\n")
}

func (l *euclidWalker) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("Error : %s\n", node)
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
