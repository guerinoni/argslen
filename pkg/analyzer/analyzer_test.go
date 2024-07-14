package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/guerinoni/argslen/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "default.go")
}

func TestSkipTest(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %s", err)
	}

	a := analyzer.NewAnalyzer()
	err = a.Flags.Set("skipTests", "true")
	if err != nil {
		t.Fatalf("failed to set flags %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, a, "skipTest.go")
}

func TestMaxExtended(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %s", err)
	}

	a := analyzer.NewAnalyzer()
	err = a.Flags.Set("maxArguments", "10")
	if err != nil {
		t.Fatalf("failed to set flags %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, a, "manyArgs.go")
}

func TestFolder(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "folder")
}
