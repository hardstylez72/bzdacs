package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/route"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
)

type Route struct {
	Id          int64    `json:"id" db:"id"`
	Route       string   `json:"route" db:"route"`
	Method      string   `json:"method" db:"method"`
	Description string   `json:"description" db:"description"`
	NamespaceId int64    `json:"namespaceId" db:"namespace_id"`
	Tags        []string `json:"tags" `
	tests.Times
}

func Delete(ctx context.Context, client *client.BZDACS, routeId, namespaceId int64) error {
	_, err := client.Route.RouteDelete(
		&route.RouteDeleteParams{
			Req: &models.RouteDeleteRequest{
				NamespaceID: &namespaceId,
				ID:          &routeId,
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

func GetById(ctx context.Context, client *client.BZDACS, id, namespaceId int64) (*Route, error) {
	res, err := client.Route.RouteGetByID(
		&route.RouteGetByIDParams{
			Req: &models.RouteGetByIDRequest{
				ID:          &id,
				NamespaceID: &namespaceId,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	p := res.GetPayload()
	return routeDTO(p), nil
}

func Update(ctx context.Context, client *client.BZDACS, r *Route) (*Route, error) {
	res, err := client.Route.RouteUpdate(
		&route.RouteUpdateParams{
			Req: &models.RouteUpdateRequest{
				Description: &r.Description,
				ID:          &r.Id,
				Method:      &r.Method,
				NamespaceID: &r.NamespaceId,
				Route:       &r.Route,
				Tags:        r.Tags,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	p := res.GetPayload()

	return routeDTO(p), nil
}

func routeDTO(p *models.RouteGetResponse) *Route {
	out := &Route{
		Id:          p.ID,
		Route:       p.Route,
		Method:      p.Method,
		Description: p.Description,
		Times: tests.Times{
			CreatedAt: tests.ParseTime(p.CreatedAt),
			UpdatedAt: tests.ParseTime(p.UpdatedAt),
			DeletedAt: nil,
		},

		NamespaceId: p.NamespaceID,
		Tags:        p.Tags,
	}
	if p.DeletedAt != nil {
		t := tests.ParseTime(*p.DeletedAt)
		out.DeletedAt = &t
	}
	return out
}

func Create(ctx context.Context, client *client.BZDACS, r *Route) (*Route, error) {
	res, err := client.Route.RouteCreate(
		&route.RouteCreateParams{
			Req: &models.RouteInsertRequest{
				Description: &r.Description,
				Method:      &r.Method,
				NamespaceID: &r.NamespaceId,
				Route:       &r.Route,
				Tags:        r.Tags,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	p := res.GetPayload()
	return routeDTO(p), nil
}
