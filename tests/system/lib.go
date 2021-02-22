package system

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/system"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
	"net/http"
)

var (
	ErrSystemGetNotFound = errors.New("system not found")
)

type System struct {
	Id   int64
	Name string
	tests.Times
}

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

func GetByName(ctx context.Context, client *client.BZDACS, name string) (*System, error) {
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
		_, ok := err.(*system.SystemGetNotFound)
		if ok {
			return nil, ErrSystemGetNotFound
		}

		return nil, err
	}

	return systemDTO(res.GetPayload()), nil
}

func GetById(ctx context.Context, client *client.BZDACS, id int64) (*System, error) {
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
	return systemDTO(res.GetPayload()), nil
}

func Update(ctx context.Context, client *client.BZDACS, s *System) (*System, error) {
	res, err := client.System.SystemUpdate(
		&system.SystemUpdateParams{
			Req: &models.SystemUpdateRequest{
				ID:   &s.Id,
				Name: &s.Name,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return systemDTO(res.GetPayload()), nil
}

func Create(ctx context.Context, client *client.BZDACS, s *System) (*System, error) {
	res, err := client.System.SystemCreate(
		&system.SystemCreateParams{
			Req:        &models.SystemInsertRequest{Name: &s.Name},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return systemDTO(res.GetPayload()), nil
}

func systemDTO(p *models.SystemGetResponse) *System {
	out := &System{
		Id:   p.ID,
		Name: p.Name,
		Times: tests.Times{
			CreatedAt: tests.ParseTime(p.CreatedAt),
			UpdatedAt: tests.ParseTime(p.UpdatedAt),
			DeletedAt: nil,
		},
	}
	if p.DeletedAt != nil {
		t := tests.ParseTime(*p.DeletedAt)
		out.DeletedAt = &t
	}
	return out
}
