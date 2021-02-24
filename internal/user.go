package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests/user"
)

func resolveUser(ctx context.Context, c *client.BZDACS, login string, namespaceId int64) (*user.User, error) {
	u, err := user.GetByLogin(ctx, c, &user.User{NamespaceId: namespaceId, Login: login})
	if err != nil {
		if err == user.ErrNotFound {
			u, err = user.Create(ctx, c, &user.User{Login: login, NamespaceId: namespaceId})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return u, nil
}
