package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests/group"
)

func resolveGroup(ctx context.Context, c *client.BZDACS, code, description string, namespaceId int64) (*group.Group, error) {
	g, err := group.GetByParams(ctx, c, &group.Group{Code: code, Description: description, NamespaceId: namespaceId})
	if err != nil {
		if err == group.ErrNotFound {
			g, err = group.Create(ctx, c, &group.Group{Code: code, Description: description, NamespaceId: namespaceId})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return g, nil
}
