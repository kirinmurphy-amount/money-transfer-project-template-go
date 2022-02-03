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
