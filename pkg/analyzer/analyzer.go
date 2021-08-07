package analyzer

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var flagSet flag.FlagSet

var (
	maxArguments uint
	skipTests    bool
)

const defaultMaxArgs = uint(5)

func init() {
	flagSet.UintVar(&maxArguments, "maxArguments", defaultMaxArgs, "max arguments allowed for function.")
	flagSet.BoolVar(&skipTests, "skipTests", false, "should the linter execute on test files as well.")
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "argsNum",
		Doc:   "linter that warns for long argument list in function.",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			f, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}

			if skipTests && strings.HasPrefix(f.Name.Name, "Test") {
				return true
			}

			numberArgs := argsLen(f)
			if numberArgs > maxArguments {
				pass.Reportf(node.Pos(), "args number for function %q is too high (%d)", f.Name.Name, numberArgs)
			}

			return true
		})
	}

	return nil, nil
}

func argsLen(fn *ast.FuncDecl) uint {
	v := argsLenVisitor{}
	ast.Walk(&v, fn)

	return v.argsCounter
}

type argsLenVisitor struct {
	argsCounter uint
}

func (v *argsLenVisitor) Visit(n ast.Node) ast.Visitor {
	list, ok := n.(*ast.FieldList)
	if ok {
		v.argsCounter = uint(list.NumFields())
	}

	return v
}
