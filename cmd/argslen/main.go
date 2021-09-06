package main

import (
	"fmt"
	"os"

	"github.com/guerinoni/argslen/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var version string

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Fprintf(os.Stdout, "argslen version: %s\n", version)

		os.Exit(0)
	}

	singlechecker.Main(analyzer.NewAnalyzer())
}
