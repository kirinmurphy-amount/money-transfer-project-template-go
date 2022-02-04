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
		ID:        "ss-sign-up-workflow",
		TaskQueue: app.SignUpTaskQueue,
	}

	signUpDetails := app.SignUpDetails{
		EmailAddress: "",
		PhoneNumber: "",
		ReferenceID: uuid.New().String(),
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.SignUpCustomer, signUpDetails)

	if err != nil {
		log.Fatalln("error starting SignUpCustomer workflow", err)
	}

	printResults(signUpDetails, we.GetID(), we.GetRunID())
}
// @@@SNIPEND
