package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-start-workflow
func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "apex-signup-task-flow",
		TaskQueue: app.SignUpTaskQueue,
	}
	signUpDetails := app.SignUpDetails{
		Email: "",
		ReferenceID: uuid.New().String(),
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, app.TransferMoney, signUpDetails)
	if err != nil {
		log.Fatalln("error starting TransferMoney workflow", err)
	}
	printResults(signUpDetails, we.GetID(), we.GetRunID())
}
// @@@SNIPEND

func printResults(signUpDetails app.SignUpDetails, workflowID, runID string) {
	log.Printf(
		"\ndefault email set to:%s, referenceId: %s\n",
		signUpDetails.Email,
		signUpDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
