package main

import (
	"log"
	"money-transfer-project-template-go/app"
)

func printResults(signUpDetails app.SignUpDetails, workflowID, runID string) {
	log.Printf(
		"\n\nSign up workflow started \nCurrent State \n  Email: %s \n  PhoneNumber: %s \n  ReferenceId: %s\n\n",
		signUpDetails.EmailAddress,
		signUpDetails.PhoneNumber,
		signUpDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
