package warmup

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/pkg/errors"
)

func resolveSystem(ctx context.Context, driver storage.SqlDriver, name string) (*system.System, error) {
	s, err := system.Get(ctx, driver, 0, name)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			s, err = system.Insert(ctx, driver, &system.System{Name: name})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return s, nil
}
