package system

type insertRequest struct {
	Name string `json:"name" validate:"required"`
} // @name systemInsertRequest

func insertRequestConvert(r *insertRequest) *System {
	return &System{
		Name: r.Name,
	}
}

type insertResponse System

func newInsertResponse(group *System) *insertResponse {
	return (*insertResponse)(group)
}

type updateRequest struct {
	Id   int    `json:"id"  validate:"required"`
	Name string `json:"name" validate:"required"`
} // @name systemUpdateRequest

func updateRequestConvert(r *updateRequest) *System {
	return &System{
		Id:   r.Id,
		Name: r.Name,
	}
}

type updateResponse System

func newUpdateResponse(group *System) *updateResponse {
	return (*updateResponse)(group)
}

type listResponse []System // @name systemListResponse

func newListResponse(systems []System) listResponse {
	return systems
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
} // @name systemDeleteRequest

type getRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
} // @name systemGetRequest

type getResponse *System // @name systemGetResponse

func newGetResponse(system *System) getResponse {
	return system
}
