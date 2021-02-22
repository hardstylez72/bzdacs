package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests/namespace"
)

func resolveNamespace(ctx context.Context, c *client.BZDACS, name string, systemId int64) (*namespace.Namespace, error) {
	s, err := namespace.GetByName(ctx, c, name, systemId)
	if err != nil {
		if err == namespace.ErrNamespaceGetNotFound {
			s, err = namespace.Create(ctx, c, name, systemId)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return s, nil
}
