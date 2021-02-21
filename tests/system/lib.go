package system

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/system"
	"github.com/hardstylez72/bzdacs/models"
	"net/http"
)

func Delete(ctx context.Context, client *client.BZDACS, id int64) error {
	_, err := client.System.SystemDelete(
		&system.SystemDeleteParams{
			Req:        &models.SystemDeleteRequest{ID: &id},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetByName(ctx context.Context, client *client.BZDACS, name string) (*models.SystemGetResponse, error) {
	res, err := client.System.SystemGet(
		&system.SystemGetParams{
			Req: &models.SystemGetRequest{
				Name: name,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func GetById(ctx context.Context, client *client.BZDACS, id int64) (*models.SystemGetResponse, error) {
	res, err := client.System.SystemGet(
		&system.SystemGetParams{
			Req: &models.SystemGetRequest{
				ID: id,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Update(ctx context.Context, client *client.BZDACS, name string, id int64) (*models.SystemUpdateResponse, error) {
	res, err := client.System.SystemUpdate(
		&system.SystemUpdateParams{
			Req: &models.SystemUpdateRequest{
				ID:   &id,
				Name: &name,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func Create(ctx context.Context, client *client.BZDACS, name string) (*models.SystemInsertResponse, error) {
	res, err := client.System.SystemCreate(
		&system.SystemCreateParams{
			Req:        &models.SystemInsertRequest{Name: &name},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}
