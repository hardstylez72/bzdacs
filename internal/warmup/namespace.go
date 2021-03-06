package warmup

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
)

func resolveNamespace(ctx context.Context, driver storage.SqlDriver, name string, systemId int) (*namespace.Namespace, error) {
	s, err := namespace.Get(ctx, driver, systemId, 0, name)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			s, err = namespace.Insert(ctx, driver, &namespace.Namespace{SystemId: systemId, Name: name})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return s, nil
}
