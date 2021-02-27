package usergroup

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	usergroup "github.com/hardstylez72/bzdacs/client/user_group"
	"github.com/hardstylez72/bzdacs/models"
)

func List(ctx context.Context, client *client.BZDACS, namespaceId, userId int64) ([]string, error) {

	var page int64 = 1

	p, err := client.UserGroup.UserGroupList(
		&usergroup.UserGroupListParams{
			Req: &models.UserGroupListRequest{
				Filter: &models.UserGroupFilter{
					BelongToUser: true,
					NamespaceID:  &namespaceId,
					Page:         &page,
					PageSize:     0,
					Pattern:      "",
					UserID:       &userId,
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
		rs = append(rs, *i.Code)
	}

	return rs, nil
}

type Pair struct {
	GroupID int64 `json:"groupId"`
	UserID  int64 `json:"userId"`
}

func Create(ctx context.Context, client *client.BZDACS, pairs []Pair) error {

	ps := make([]*models.UserGroupPair, 0, len(pairs))
	for i := range pairs {
		ps = append(ps, &models.UserGroupPair{
			GroupID: &pairs[i].GroupID,
			UserID:  &pairs[i].UserID,
		})
	}

	_, err := client.UserGroup.UserGroupCreate(
		&usergroup.UserGroupCreateParams{
			Req: &models.UserGroupInsertRequest{
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
