package routetag

import (
	"context"
	"database/sql"
	"github.com/hardstylez72/bblog/ad/pkg/tag"
	"github.com/jmoiron/sqlx"
)

func Merge(ctx context.Context, conn *sqlx.DB, tx *sqlx.Tx, routeId int, tagNames []string) ([]string, error) {

	currentTagIds, err := GetRouteTags(ctx, conn, routeId)
	if err != nil {
		return nil, err
	}

	newTagIds := make([]int, 0)

	for _, tagName := range tagNames {
		t, err := tag.GetByNameDb(ctx, conn, tagName)
		if err != nil {
			if err == tag.ErrNotFound {
				var creationTagError error
				t, creationTagError = tag.InsertTx(ctx, tx, &tag.Tag{Name: tagName})
				if creationTagError != nil {
					return nil, creationTagError
				}
			} else {
				return nil, err
			}
		}

		if !contains(currentTagIds, t.Id) {
			err = insertPair(ctx, tx, routeId, t.Id)
			if err != nil {
				return nil, err
			}
		}
		newTagIds = append(newTagIds, t.Id)
	}

	for i := range currentTagIds {
		if !contains(newTagIds, currentTagIds[i]) {
			err = deletePair(ctx, tx, routeId, currentTagIds[i])
			if err != nil {
				return nil, err
			}
		}
	}

	return tagNames, nil
}

func contains(set []int, el int) bool {
	for i := range set {
		if set[i] == el {
			return true
		}
	}
	return false
}

func GetRouteTags(ctx context.Context, conn *sqlx.DB, routeId int) ([]int, error) {
	query := `
			select tag_id 
			  from ad.routes_tags 
			 where route_id = $1
`

	ids := make([]int, 0)
	err := conn.SelectContext(ctx, &ids, query, routeId)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func isPairExist(ctx context.Context, conn *sqlx.DB, routeId, tagId int) (bool, error) {
	query := `
			select count(*) 
			  from ad.routes_tags 
			 where tag_id = $1 
			   and route_id = $2
`

	var count int
	err := conn.GetContext(ctx, &count, query, tagId, routeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return count != 0, nil
}

func insertPair(ctx context.Context, tx *sqlx.Tx, routeId, tagId int) error {
	query := `
			insert into ad.routes_tags (
					   tag_id,
					   route_id
					   )
				   values (
					   $1,
					   $2
				   )
`

	_, err := tx.ExecContext(ctx, query, tagId, routeId)
	if err != nil {
		return err
	}

	return nil
}
func deletePair(ctx context.Context, tx *sqlx.Tx, routeId, tagId int) error {
	query := `delete from ad.routes_tags where route_id = $1 and tag_id = $2`

	_, err := tx.ExecContext(ctx, query, routeId, tagId)
	if err != nil {
		return err
	}

	return nil
}
