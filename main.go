package main

import (
	"github.com/ansnadeem/temporalint/analysis/passes/deterministic"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(deterministic.Analyzer)
}
