package tag

type insertRequest struct {
	Name string `json:"name" validate:"required"`
}

func insertRequestConvert(r *insertRequest) *Tag {
	return &Tag{
		Name: r.Name,
	}
}

type insertResponse Tag

func newInsertResponse(group *Tag) *insertResponse {
	return (*insertResponse)(group)
}

type updateRequest struct {
	Id   int    `json:"id"  validate:"required"`
	Name string `json:"name" validate:"required"`
}

func updateRequestConvert(r *updateRequest) *Tag {
	return &Tag{
		Id:   r.Id,
		Name: r.Name,
	}
}

type updateResponse Tag

func newUpdateResponse(group *Tag) *updateResponse {
	return (*updateResponse)(group)
}

type listResponse []Tag

func newListResponse(groups []Tag) listResponse {
	return groups
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type getRequest struct {
	Id int `json:"id" validate:"required"`
}

type suggestRequest struct {
	Pattern string `json:"pattern" validate:"required"`
}
