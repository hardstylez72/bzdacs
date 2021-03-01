package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	user "github.com/hardstylez72/bzdacs/tests/sysuser"
)

func registerUser(ctx context.Context, c *client.BZDACS, login, password string) error {
	err := user.Create(ctx, c, login, password)
	if err != nil {
		if err != user.ErrAlreadyRegistered {
			return err
		}

	}
	return nil
}
