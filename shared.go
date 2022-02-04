package app

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
const SignUpTaskQueue = "SIGN_UP_TASK_QUEUE"
// @@@SNIPEND

type SignUpDetails struct {
	EmailAddress string
	PhoneNumber string
	ReferenceID string
}
