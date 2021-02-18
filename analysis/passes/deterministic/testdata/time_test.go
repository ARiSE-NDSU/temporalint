package testdata

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func testWorkflow(ctx workflow.Context) {
	time.Sleep(1)
}
