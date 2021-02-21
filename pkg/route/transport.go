package route

type insertRequest struct {
	Route       string   `json:"route" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	Tags        []string `json:"tags"`
}

type updateRequest struct {
	Route       string   `json:"route" validate:"required"`
	Id          int      `json:"id" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	Tags        []string `json:"tags"`
}

func insertRequestConvert(r *insertRequest) *Route {
	return &Route{
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
		Tags:        r.Tags,
	}
}

func updateRequestConvert(r *updateRequest) *Route {
	return &Route{
		Id:          r.Id,
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
	}
}

type insertResponse *Route

func newInsertResponse(route *Route) insertResponse {
	return route
}

type listResponse []RouteWithTags

func newListResponse(routes []RouteWithTags) listResponse {
	return routes
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type getRequest struct {
	Id int `json:"id" validate:"required"`
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
