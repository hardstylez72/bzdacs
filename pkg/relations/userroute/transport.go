package userroute

type insertRequest struct {
	Pairs []Pair `json:"pairs" validate:"required,dive"`
} // @name userRouteInsertRequest

type insertResponse []Route // @name userRouteInsertResponse
type updateResponse *Route  // @name userRouteUpdateResponse

type deleteRequest struct {
	Pairs []PairToDelete `json:"pairs" validate:"required,dive"`
} // @name userRouteDeleteRequest

type listRequest struct {
	Filter Filter `json:"filter"`
} // @name userRouteListRequest

type Filter struct {
	Page         int    `json:"page"  validate:"required,gte=1"`
	PageSize     int    `json:"pageSize"`
	Pattern      string `json:"pattern"`
	BelongToUser bool   `json:"belongToUser"`
	NamespaceId  int    `json:"namespaceId" validate:"required,gte=1"`
	UserId       int    `json:"userId" validate:"required,gte=1"`
} // @name userRouteFilter

type listResponse struct {
	Items []RouteWithGroups `json:"items" validate:"required"`
	Total int               `json:"total" validate:"required"`
} // @name userRouteListResponse

func newListResponse(groups []RouteWithGroups, total int) listResponse {
	out := make([]RouteWithGroups, 0, len(groups))
	for i := range groups {
		out = append(out, groups[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}
