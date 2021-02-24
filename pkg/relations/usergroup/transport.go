package usergroup

type insertRequest struct {
	Pairs []Pair `json:"pairs" validate:"required,dive"`
} // @name userGroupInsertRequest

type insertResponse []Group // @name userGroupInsertResponse

type deleteRequest struct {
	Pairs []Pair `json:"pairs" validate:"required,dive"`
} // @name userGroupDeleteRequest

type listRequest struct {
	Filter Filter `json:"filter"`
} // @name userGroupListRequest

type Filter struct {
	Page         int    `json:"page"  validate:"required,gte=1"`
	PageSize     int    `json:"pageSize"`
	Pattern      string `json:"pattern"`
	BelongToUser bool   `json:"belongToUser"`
	NamespaceId  int    `json:"namespaceId" validate:"required,gte=1"`
	UserId       int    `json:"userId" validate:"required,gte=1"`
} // @name userGroupFilter

type listResponse struct {
	Items []Group `json:"items" validate:"required"`
	Total int     `json:"total" validate:"required"`
} // @name userGroupListResponse

func newListResponse(groups []Group, total int) listResponse {
	out := make([]Group, 0, len(groups))
	for i := range groups {
		out = append(out, groups[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}
