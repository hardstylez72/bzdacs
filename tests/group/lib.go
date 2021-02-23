package group

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/group"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
)

type Group struct {
	Id          int64  `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Description string `json:"description" db:"description"`
	NamespaceId int64  `json:"namespaceId" db:"namespace_id"`
	tests.Times
}

var (
	ErrNotFound = errors.New("entity not found")
)

func Delete(ctx context.Context, client *client.BZDACS, groupId, namespaceId int64) error {
	_, err := client.Group.GroupDelete(
		&group.GroupDeleteParams{
			Req: &models.GroupDeleteRequest{
				NamespaceID: &namespaceId,
				ID:          &groupId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return wrapErr(err)
	}
	return nil
}

func GetByParams(ctx context.Context, client *client.BZDACS, g *Group) (*Group, error) {
	res, err := client.Group.GroupGetByCode(
		&group.GroupGetByCodeParams{
			Req: &models.GroupGetByCodeRequest{
				Code:        &g.Code,
				NamespaceID: &g.NamespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, wrapErr(err)
	}
	p := res.GetPayload()
	return routeDTO(p), nil
}

func GetById(ctx context.Context, client *client.BZDACS, id, namespaceId int64) (*Group, error) {
	res, err := client.Group.GroupGetByID(
		&group.GroupGetByIDParams{
			Req: &models.GroupGetByIDRequest{
				ID:          &id,
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, wrapErr(err)
	}
	p := res.GetPayload()
	return routeDTO(p), nil
}

func Update(ctx context.Context, client *client.BZDACS, r *Group) (*Group, error) {
	res, err := client.Group.GroupUpdate(
		&group.GroupUpdateParams{
			Req: &models.GroupUpdateRequest{
				Description: &r.Description,
				ID:          &r.Id,
				Code:        &r.Code,
				NamespaceID: &r.NamespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, wrapErr(err)
	}
	p := res.GetPayload()

	return routeDTO(p), nil
}

func Create(ctx context.Context, client *client.BZDACS, r *Group) (*Group, error) {
	res, err := client.Group.GroupCreate(
		&group.GroupCreateParams{
			Req: &models.GroupInsertRequest{
				Description: &r.Description,
				Code:        &r.Code,
				NamespaceID: &r.NamespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, wrapErr(err)
	}
	p := res.GetPayload()
	return routeDTO(p), nil
}

func routeDTO(p *models.GroupGetResponse) *Group {
	out := &Group{
		Id:          *p.ID,
		Code:        *p.Code,
		Description: *p.Description,
		Times: tests.Times{
			CreatedAt: tests.ParseTime(*p.CreatedAt),
			UpdatedAt: tests.ParseTime(*p.UpdatedAt),
			DeletedAt: nil,
		},

		NamespaceId: *p.NamespaceID,
	}
	if p.DeletedAt != nil {
		t := tests.ParseTime(*p.DeletedAt)
		out.DeletedAt = &t
	}
	return out
}

func wrapErr(err error) error {
	_, ok := err.(*group.GroupGetByCodeNotFound)
	if ok {
		return ErrNotFound
	}
	_, ok = err.(*group.GroupGetByIDNotFound)
	if ok {
		return ErrNotFound
	}
	return err
}
