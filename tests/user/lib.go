package user

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/generated/client/user"
	"github.com/hardstylez72/bzdacs/generated/models"
	"github.com/hardstylez72/bzdacs/tests"
)

type User struct {
	Id          int64  `json:"id" db:"id"`
	Login       string `json:"route" db:"route"`
	NamespaceId int64  `json:"namespaceId" db:"namespace_id"`
	tests.Times
}

var (
	ErrNotFound = errors.New("entity not found")
)

func Delete(ctx context.Context, client *client.BZDACS, userId, namespaceId int64) error {
	_, err := client.User.UserDelete(
		&user.UserDeleteParams{
			Req: &models.UserDeleteRequest{
				NamespaceID: &namespaceId,
				ID:          &userId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return wrapErr(err)
	}
	return nil
}

func GetByLogin(ctx context.Context, client *client.BZDACS, r *User) (*User, error) {
	res, err := client.User.UserGetByLogin(
		&user.UserGetByLoginParams{
			Req: &models.UserGetByLoginRequest{
				Login:       &r.Login,
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

func GetById(ctx context.Context, client *client.BZDACS, id, namespaceId int64) (*User, error) {
	res, err := client.User.UserGetByID(
		&user.UserGetByIDParams{
			Req: &models.UserGetByIDRequest{
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

func Update(ctx context.Context, client *client.BZDACS, r *User) (*User, error) {
	res, err := client.User.UserUpdate(
		&user.UserUpdateParams{
			Req: &models.UserUpdateRequest{
				ID:          &r.Id,
				ExternalID:  &r.Login,
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

func Create(ctx context.Context, client *client.BZDACS, r *User) (*User, error) {
	res, err := client.User.UserCreate(
		&user.UserCreateParams{
			Req: &models.UserCreateRequest{
				NamespaceID: &r.NamespaceId,
				ExternalID:  &r.Login,
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

func routeDTO(p *models.UserGetResponse) *User {
	out := &User{
		Id:    *p.ID,
		Login: *p.ExternalID,
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
	_, ok := err.(*user.UserGetByIDNotFound)
	if ok {
		return ErrNotFound
	}
	_, ok = err.(*user.UserGetByLoginNotFound)
	if ok {
		return ErrNotFound
	}
	return err
}
