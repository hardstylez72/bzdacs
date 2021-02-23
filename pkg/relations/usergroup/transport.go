package usergroup

type insertRequest []Pair

func insertRequestConvert(r insertRequest) []Pair {
	return r
}

type listRequest struct {
	UserId       int  `json:"userId" validate:"required"`
	BelongToUser bool `json:"belongToUser"`
}

type listResponse []Group

func newListResponse(routes []Group) listResponse {
	return routes
}

type deleteRequest struct {
	Params []Pair `json:"groupId" validate:"dive"`
}
