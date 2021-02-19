package testdata

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func workflowWithChannel(ctx workflow.Context) {
	c := make(chan int)
	x := <-c
	fmt.Print(x)
}
