package tag

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

func MergeLL(ctx context.Context, driver storage.SqlDriver, tagNames []string, namespaceId int) ([]Tag, error) {
	merged := make([]Tag, 0)
	for _, tagName := range tagNames {
		tag, err := GetByNameLL(ctx, driver, tagName, namespaceId)
		if err != nil {
			if errors.Is(err, storage.EntityNotFound) {
				var insertErr error
				tag, insertErr = InsertLL(ctx, driver, &Tag{
					Name:        tagName,
					NamespaceId: namespaceId,
				})
				if insertErr != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
		merged = append(merged, *tag)
	}
	return merged, nil
}

func GetByNameLL(ctx context.Context, driver storage.SqlDriver, name string, namespaceId int) (*Tag, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from tags
	   where deleted_at is null
		 and name = $1
		 and namespace_id = $2
`
	var tag Tag
	err := driver.GetContext(ctx, &tag, query, name, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &tag, nil
}
