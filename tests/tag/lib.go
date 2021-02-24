package tag

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/tag"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
)

type Tag struct {
	Id          int64
	Name        string
	NamespaceId int64
	tests.Times
}

func Delete(ctx context.Context, client *client.BZDACS, t *Tag) error {
	_, err := client.Tag.TagDelete(
		&tag.TagDeleteParams{
			Req: &models.TagDeleteRequest{
				NamespaceID: &t.NamespaceId,
				ID:          &t.Id,
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

func GetById(ctx context.Context, client *client.BZDACS, id int64) (*Tag, error) {
	res, err := client.Tag.TagGet(
		&tag.TagGetParams{
			Req: &models.TagGetRequest{
				ID: &id,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return tagDTO(res.GetPayload()), nil
}

func Update(ctx context.Context, client *client.BZDACS, t *Tag) (*Tag, error) {
	res, err := client.Tag.TagUpdate(
		&tag.TagUpdateParams{
			Req: &models.TagUpdateRequest{
				ID:          &t.Id,
				Name:        &t.Name,
				NamespaceID: &t.NamespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return tagDTO(res.GetPayload()), nil
}

func Create(ctx context.Context, client *client.BZDACS, t *Tag) (*Tag, error) {
	res, err := client.Tag.TagCreate(
		&tag.TagCreateParams{
			Req: &models.TagInsertRequest{
				Name:        &t.Name,
				NamespaceID: &t.NamespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return tagDTO(res.GetPayload()), nil
}

func tagDTO(p *models.TagGetResponse) *Tag {
	out := &Tag{
		Id:          *p.ID,
		Name:        *p.Name,
		NamespaceId: *p.NamespaceID,
		Times: tests.Times{
			CreatedAt: tests.ParseTime(*p.CreatedAt),
			UpdatedAt: tests.ParseTime(*p.UpdatedAt),
			DeletedAt: nil,
		},
	}
	if p.DeletedAt != nil {
		t := tests.ParseTime(*p.DeletedAt)
		out.DeletedAt = &t
	}
	return out
}
