package userroute

type insertRequest []params

func insertRequestConvert(r insertRequest) []params {
	return r
}

type listRequest struct {
	UserId       int  `json:"userId" validate:"required"`
	BelongToUser bool `json:"belongToUser"`
}

type listResponse []RouteWithGroups

func newListResponse(routes []RouteWithGroups) listResponse {
	return routes
}

type deleteRequest struct {
	Params []params `json:"groupId" validate:"dive"`
}
