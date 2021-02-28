package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests/group"
	"github.com/hardstylez72/bzdacs/tests/relations/usergroup"
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

func resolveUserAndGroup(ctx context.Context, c *client.BZDACS, g *group.Group, u *user.User) error {

	currentGroups, err := usergroup.List(ctx, c, u.NamespaceId, u.Id)
	if err != nil {
		return err
	}

	if !contains(currentGroups, g.Code) {
		err = usergroup.Create(ctx, c, []usergroup.Pair{{
			GroupID: g.Id,
			UserID:  u.Id,
		}})
		if err != nil {
			return err
		}
	}

	return nil
}

func contains(routes []string, target string) bool {
	for _, r := range routes {
		if r == target {
			return true
		}
	}
	return false
}
