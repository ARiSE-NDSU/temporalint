package deterministic

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDeterministic(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
