package app

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	// Mock activity implementation
	testDetails := SignUpDetails{
		EmailAddress:  "",
		PhoneNumber: "",
		ReferenceID: "1234",
	}
	env.OnActivity(Withdraw, mock.Anything, testDetails).Return(nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return(nil)
	env.ExecuteWorkflow(SignUpCustomer, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}
