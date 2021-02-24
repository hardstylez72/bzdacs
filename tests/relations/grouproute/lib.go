package grouproute

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	grouproute "github.com/hardstylez72/bzdacs/client/group_route"
	"github.com/hardstylez72/bzdacs/models"
)

func List(ctx context.Context, client *client.BZDACS, namespaceId, groupId int64) ([]string, error) {

	var page int64 = 1

	p, err := client.GroupRoute.GroupRoutesList(
		&grouproute.GroupRoutesListParams{
			Req: &models.GroupRouteListRequest{
				Filter: &models.GroupRouteFilter{
					BelongToGroup: true,
					GroupID:       &groupId,
					NamespaceID:   &namespaceId,
					Page:          &page,
					PageSize:      0,
					RoutePattern:  "",
				},
			},
			Context: ctx,
		},
	)
	if err != nil {
		return nil, err
	}
	rs := make([]string, 0, len(p.Payload.Items))
	for _, i := range p.Payload.Items {
		rs = append(rs, *i.Route)
	}

	return rs, nil
}

type Pair struct {
	GroupID int64 `json:"groupId"`
	RouteID int64 `json:"routeId"`
}

func Create(ctx context.Context, client *client.BZDACS, pairs []Pair) error {

	ps := make([]*models.GroupRoutePair, 0, len(pairs))
	for i := range pairs {
		ps = append(ps, &models.GroupRoutePair{
			GroupID: &pairs[i].GroupID,
			RouteID: &pairs[i].RouteID,
		})
	}

	_, err := client.GroupRoute.GroupRoutesCreate(
		&grouproute.GroupRoutesCreateParams{
			Req: &models.GroupRouteInsertRequest{
				Pairs: ps,
			},
			Context: ctx,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
