package testdata

import "go.temporal.io/sdk/workflow"

func doNothing() {

}

func workflowWithGoRoutine(ctx workflow.Context) {
	go doNothing()
}
