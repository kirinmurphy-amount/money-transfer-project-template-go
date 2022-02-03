package app

import (
	"context"
	"fmt"
)

func Withdraw(ctx context.Context, signUpDetails SignUpDetails) error {
	fmt.Printf(
		"\nUpdating email address: %s. ReferenceId: %s\n",
		signUpDetails.Email,
		signUpDetails.ReferenceID,
	)
	return nil
}


// @@@SNIPSTART money-transfer-project-template-go-activity
func Deposit(ctx context.Context, signUpDetails SignUpDetails) error {
	fmt.Printf(
		"\nUpdating email address: %s. ReferenceId: %s\n",
		signUpDetails.Email,
		signUpDetails.ReferenceID,
	)
	// Switch out comments on the return statements to simulate an error
	// return fmt.Errorf("deposit did not occur due to an issue")
	return nil
}
// @@@SNIPEND"
