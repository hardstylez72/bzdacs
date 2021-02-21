package tag

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/tag"
	"github.com/hardstylez72/bzdacs/models"
)

func Delete(ctx context.Context, client *client.BZDACS, tagId, namespaceId int64) error {
	_, err := client.Tag.TagDelete(
		&tag.TagDeleteParams{
			Req: &models.TagDeleteRequest{
				NamespaceID: &namespaceId,
				ID:          &tagId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

//func GetByName(ctx context.Context, client *client.BZDACS, name string, SystemID int64) (*models.NamespaceGetResponse, error) {
//	res, err := client.Namespace.NamespaceGet(
//		&namespace.NamespaceGetParams{
//			Req: &models.NamespaceGetRequest{
//				Name:     name,
//				SystemID: SystemID,
//			},
//			Context: ctx,
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//	return res.GetPayload(), nil
//}

func GetById(ctx context.Context, client *client.BZDACS, id, namespaceId int64) (*models.TagGetResponse, error) {
	res, err := client.Tag.TagGet(
		&tag.TagGetParams{
			Req: &models.TagGetRequest{
				ID:          &id,
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Update(ctx context.Context, client *client.BZDACS, name string, id, namespaceId int64) (*models.TagUpdateResponse, error) {
	res, err := client.Tag.TagUpdate(
		&tag.TagUpdateParams{
			Req: &models.TagUpdateRequest{
				ID:          &id,
				Name:        &name,
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Create(ctx context.Context, client *client.BZDACS, name string, namespaceId int64) (*models.TagInsertResponse, error) {
	res, err := client.Tag.TagCreate(
		&tag.TagCreateParams{
			Req: &models.TagInsertRequest{
				Name:        &name,
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}
