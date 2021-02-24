package route

type insertRequest struct {
	Route       string   `json:"route" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	NamespaceId int      `json:"namespaceId" validate:"required"`
	Tags        []string `json:"tags"`
} // @name routeInsertRequest

type updateRequest struct {
	Route       string   `json:"route" validate:"required"`
	Id          int      `json:"id" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	NamespaceId int      `json:"namespaceId" validate:"required"`
	Tags        []string `json:"tags"`
} // @name routeUpdateRequest

func insertRequestConvert(r *insertRequest) *Route {
	return &Route{
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
		Tags:        r.Tags,
		NamespaceId: r.NamespaceId,
	}
}

type insertResponse *Route

func newInsertResponse(route *Route) insertResponse {
	return route
}

func updateRequestConvert(r *updateRequest) *Route {
	return &Route{
		Id:          r.Id,
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
		Tags:        r.Tags,
		NamespaceId: r.NamespaceId,
	}
}

type updateResponse *Route

func newUpdateResponse(route *Route) updateResponse {
	return route
}

type listResponse struct {
	Items []getResponse `json:"items"`
	Total int           `json:"total"`
} // @name routeListResponse

func newListResponse(routes []Route, total int) listResponse {
	out := make([]getResponse, 0, len(routes))
	for i := range routes {
		out = append(out, &routes[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name routeDeleteRequest

type getByIdRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name routeGetByIdRequest

type getByParamsRequest struct {
	Route       string `json:"route" validate:"required"`
	Method      string `json:"method" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name routeGetByParamsRequest

type getResponse *Route // @name routeGetResponse

func newGetResponse(route *Route) getResponse {
	return route
}

type listRequest struct {
	Filter filter `json:"filter"`
} // @name routeListRequest

type filter struct {
	Page        int `json:"page"  validate:"required,gte=1"`
	PageSize    int `json:"pageSize"  validate:"required,gte=1"`
	NamespaceId int `json:"namespaceId" validate:"required,gte=1"`
}
