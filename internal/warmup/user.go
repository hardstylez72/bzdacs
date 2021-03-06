package warmup

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/user"
)

func resolveUser(ctx context.Context, driver storage.SqlDriver, login string, namespaceId int) (*user.User, error) {
	u, err := user.GetByExternalId(ctx, driver, login, namespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			u, err = user.Insert(ctx, driver, &user.User{ExternalId: login, NamespaceId: namespaceId})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return u, nil
}

func resolveUserAndGroup(ctx context.Context, driver storage.SqlDriver, g *group.Group, u *user.User) error {

	currentGroups, _, err := usergroup.List(ctx, driver, usergroup.Filter{
		Page:         1,
		PageSize:     usergroup.NoPageSizeLimit,
		Pattern:      usergroup.NoPattern,
		BelongToUser: true,
		NamespaceId:  u.NamespaceId,
		UserId:       u.Id,
	})
	if err != nil {
		return err
	}

	if !contains(currentGroups, *g) {
		_, err = usergroup.InsertPair(ctx, driver, g.Id, u.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func contains(groups []usergroup.Group, target group.Group) bool {
	for _, group := range groups {
		if group.Code == target.Code {
			return true
		}
	}
	return false
}
