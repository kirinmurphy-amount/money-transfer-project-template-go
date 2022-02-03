package app

// @@@SNIPSTART money-transfer-project-template-go-shared-task-queue
const SignUpTaskQueue = "SIGN_UP_TASK_QUEUE"
// @@@SNIPEND

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}

type SignUpDetails struct {
	Email string
	ReferenceID string
}