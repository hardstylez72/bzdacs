package route

type insertRequest struct {
	Route       string   `json:"route" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	NamespaceId int      `json:"namespaceId" validate:"required"`
	Tags        []string `json:"tags"`
}

type updateRequest struct {
	Route       string   `json:"route" validate:"required"`
	Id          int      `json:"id" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	NamespaceId int      `json:"namespaceId" validate:"required"`
	Tags        []string `json:"tags"`
}

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

type listResponse []Route

func newListResponse(routes []RouteWithTags) listResponse {
	return nil
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
}

type getByIdRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
}

type getResponse *Route

func newGetResponse(route *Route) getResponse {
	return route
}

type listRequest struct {
	Filter filter `json:"filter"`
}

type filter struct {
	Tags     tags  `json:"tags"`
	SystemId int64 `json:"systemId" validate:"required"`
}

type tags struct {
	Names   []string `json:"names"`
	Exclude bool     `json:"exclude"`
}
