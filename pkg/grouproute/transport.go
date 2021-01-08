package grouproute

type insertRequest []Pair

func insertRequestConvert(r insertRequest) []Pair {
	return r
}

type listRequest struct {
	GroupId       int  `json:"groupId" validate:"required"`
	BelongToGroup bool `json:"belongToGroup"`
}

type listResponse []Route

func newListResponse(routes []Route) listResponse {
	return routes
}

type deleteRequest struct {
	Params []Pair `json:"groupId" validate:"dive"`
}
