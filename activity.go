package app

import (
	"context"
	"fmt"
)

func Withdraw(ctx context.Context, signUpDetails SignUpDetails) error {
	fmt.Printf(
		"\nWithdrawing $%s from account %s. ReferenceId: %s\n",
		signUpDetails.EmailAddress,
		signUpDetails.PhoneNumber,
		signUpDetails.ReferenceID,
	)
	return nil
}

// @@@SNIPSTART money-transfer-project-template-go-activity
func Deposit(ctx context.Context, signUpDetails SignUpDetails) error {
	fmt.Printf(
		"\nDepositing $%s into account %s. ReferenceId: %s\n",
		signUpDetails.EmailAddress,
		signUpDetails.ReferenceID,
	)
	// Switch out comments on the return statements to simulate an error
	//return fmt.Errorf("deposit did not occur due to an issue")
	return nil
}
// @@@SNIPEND"
