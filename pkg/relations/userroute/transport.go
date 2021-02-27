package userroute

type insertRequest []Pair

func insertRequestConvert(r insertRequest) []Pair {
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
	Params []Pair `json:"groupId" validate:"dive"`
}
