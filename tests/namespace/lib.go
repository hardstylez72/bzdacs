package namespace

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/namespace"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
)

var (
	ErrNamespaceGetNotFound = errors.New("namespace not found")
)

type Namespace struct {
	Id       int64
	Name     string
	SystemId int64
	tests.Times
}

func Delete(ctx context.Context, client *client.BZDACS, namespaceId, systemId int64) error {
	_, err := client.Namespace.NamespaceDelete(
		&namespace.NamespaceDeleteParams{
			Req: &models.NamespaceDeleteRequest{
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetByName(ctx context.Context, client *client.BZDACS, name string, SystemID int64) (*Namespace, error) {
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
		_, ok := err.(*namespace.NamespaceGetNotFound)
		if ok {
			return nil, ErrNamespaceGetNotFound
		}
		return nil, err
	}
	return namespaceDTO(res.GetPayload()), nil
}

func GetById(ctx context.Context, client *client.BZDACS, id int64) (*Namespace, error) {
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
	return namespaceDTO(res.GetPayload()), nil
}

func Update(ctx context.Context, client *client.BZDACS, name string, id int64) (*Namespace, error) {
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
	return namespaceDTO(res.GetPayload()), nil
}

func Create(ctx context.Context, client *client.BZDACS, name string, systemId int64) (*Namespace, error) {
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
	return namespaceDTO(res.GetPayload()), nil
}

func namespaceDTO(p *models.NamespaceGetResponse) *Namespace {
	out := &Namespace{
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
