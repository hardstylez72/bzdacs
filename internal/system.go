package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests/system"
)

func resolveSystem(ctx context.Context, c *client.BZDACS, name string) (*system.System, error) {
	s, err := system.GetByName(ctx, c, name)
	if err != nil {
		if err == system.ErrSystemGetNotFound {
			s, err = system.Create(ctx, c, &system.System{Name: name})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return s, nil
}
