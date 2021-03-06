package routetag

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/tag"
)

func Merge(ctx context.Context, driver storage.SqlDriver, routeId int, tagNames []string, namespaceId int) ([]tag.Tag, error) {

	oldRouteTags, err := GetRouteTags(ctx, driver, routeId, namespaceId)
	if err != nil {
		return nil, err
	}
	oldRouteTagIds := getTagIdsFromArray(oldRouteTags)

	tags, err := tag.Merge(ctx, driver, tagNames, namespaceId)
	if err != nil {
		return nil, err
	}

	newTagIds := getTagIdsFromArray(tags)
	tagsToInsert := getTagsToInsert(oldRouteTagIds, newTagIds)

	for _, tagId := range tagsToInsert {
		err = insertPair(ctx, driver, routeId, tagId)
		if err != nil {
			return nil, err
		}
	}

	tagsToDelete := getTagsToDelete(oldRouteTagIds, newTagIds)
	for _, tagId := range tagsToDelete {
		err = deletePair(ctx, driver, routeId, tagId)
		if err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func getTagIdsFromArray(tags []tag.Tag) []int {
	ids := make([]int, 0, len(tags))

	for _, tag := range tags {
		ids = append(ids, tag.Id)
	}
	return ids
}

func getTagsToInsert(oldIds, newIds []int) []int {
	result := make([]int, 0)
	for _, newId := range newIds {
		if !contains(oldIds, newId) {
			result = append(result, newId)
		}
	}
	return result
}

func getTagsToDelete(oldIds, newIds []int) []int {
	result := make([]int, 0)
	for _, oldId := range oldIds {
		if !contains(newIds, oldId) {
			result = append(result, oldId)
		}
	}
	return result
}

func contains(set []int, el int) bool {
	for i := range set {
		if set[i] == el {
			return true
		}
	}
	return false
}

func GetRouteTags(ctx context.Context, driver storage.SqlDriver, routeId, namespaceId int) ([]tag.Tag, error) {
	query := `
			select 	   t.id, 
					   t.name,
                       t.created_at,
                       t.updated_at,
                       t.deleted_at,
					   t.namespace_id
			  from routes_tags rt
			   join tags t on rt.tag_id = t.id
			 where rt.route_id = $1
			   and t.namespace_id = $2
`
	tags := make([]tag.Tag, 0)
	err := driver.SelectContext(ctx, &tags, query, routeId, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return tags, nil
}

func insertPair(ctx context.Context, driver storage.SqlDriver, routeId, tagId int) error {
	query := `
			insert into routes_tags (
					   tag_id,
					   route_id
					   )
				   values (
					   $1,
					   $2
				   )
`

	_, err := driver.ExecContext(ctx, query, tagId, routeId)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}
func deletePair(ctx context.Context, driver storage.SqlDriver, routeId, tagId int) error {
	query := `delete from routes_tags where route_id = $1 and tag_id = $2`

	_, err := driver.ExecContext(ctx, query, routeId, tagId)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}
