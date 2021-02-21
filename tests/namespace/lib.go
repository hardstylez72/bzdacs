package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/namespace"
	"github.com/hardstylez72/bzdacs/models"
)

func Delete(ctx context.Context, client *client.BZDACS, namespaceId, systemId int64) error {
	_, err := client.Namespace.NamespaceDelete(
		&namespace.NamespaceDeleteParams{
			Req: &models.NamespaceDeleteRequest{
				NamespaceID: &namespaceId,
				SystemID:    &systemId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetByName(ctx context.Context, client *client.BZDACS, name string, SystemID int64) (*models.NamespaceGetResponse, error) {
	res, err := client.Namespace.NamespaceGet(
		&namespace.NamespaceGetParams{
			Req: &models.NamespaceGetRequest{
				Name:     name,
				SystemID: SystemID,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func GetById(ctx context.Context, client *client.BZDACS, id int64) (*models.NamespaceGetResponse, error) {
	res, err := client.Namespace.NamespaceGet(
		&namespace.NamespaceGetParams{
			Req: &models.NamespaceGetRequest{
				ID: id,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Update(ctx context.Context, client *client.BZDACS, name string, id int64) (*models.NamespaceUpdateResponse, error) {
	res, err := client.Namespace.NamespaceUpdate(
		&namespace.NamespaceUpdateParams{
			Req: &models.NamespaceUpdateRequest{
				ID:   &id,
				Name: &name,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Create(ctx context.Context, client *client.BZDACS, name string, systemId int64) (*models.NamespaceInsertResponse, error) {
	res, err := client.Namespace.NamespaceCreate(
		&namespace.NamespaceCreateParams{
			Req: &models.NamespaceInsertRequest{
				Name:     &name,
				SystemID: &systemId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}
