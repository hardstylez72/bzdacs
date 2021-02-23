package grouproute

type insertRequest struct {
	Pairs []Pair `json:"pairs" validate:"required,dive"`
} // @name groupRouteInsertRequest

type insertResponse []Route // @name groupRouteInsertResponse

type listRequest struct {
	Filter Filter `json:"filter"`
} // @name groupRouteListRequest

type deleteRequest struct {
	Pairs []Pair `json:"pairs" validate:"required,dive"`
} // @name groupRouteDeleteRequest

type Filter struct {
	Page          int    `json:"page"  validate:"required,gte=1"`
	PageSize      int    `json:"pageSize"`
	RoutePattern  string `json:"routePattern"`
	BelongToGroup bool   `json:"belongToGroup"`
	NamespaceId   int    `json:"namespaceId" validate:"required,gte=1"`
	GroupId       int    `json:"groupId" validate:"required,gte=1"`
} // @name groupRouteFilter

type listResponse struct {
	Items []Route `json:"items" validate:"required"`
	Total int     `json:"total" validate:"required"`
} // @name groupRouteListResponse

func newListResponse(routes []Route, total int) listResponse {
	out := make([]Route, 0, len(routes))
	for i := range routes {
		out = append(out, routes[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}
