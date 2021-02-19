package testdata

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func workflowWithTime(ctx workflow.Context) {
	time.Sleep(1)
}
