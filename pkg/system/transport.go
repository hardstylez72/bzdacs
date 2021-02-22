package system

type insertRequest struct {
	Name string `json:"name" validate:"required"`
}

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
}

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

type listResponse []System

func newListResponse(systems []System) listResponse {
	return systems
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type getRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type getResponse *System

func newGetResponse(system *System) getResponse {
	return system
}
