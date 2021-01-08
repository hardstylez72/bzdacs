package group

type insertRequest struct {
	BaseGroupId int    `json:"baseGroupId"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func insertRequestConvert(r *insertRequest) *Group {
	return &Group{
		Code:        r.Code,
		Description: r.Description,
	}
}

type insertResponse Group

func newInsertResponse(group *Group) *insertResponse {
	return (*insertResponse)(group)
}

type updateRequest struct {
	Id          int    `json:"id"  validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func updateRequestConvert(r *updateRequest) *Group {
	return &Group{
		Id:          r.Id,
		Code:        r.Code,
		Description: r.Description,
	}
}

type updateResponse Group

func newUpdateResponse(group *Group) *updateResponse {
	return (*updateResponse)(group)
}

type listResponse []Group

func newListResponse(groups []Group) listResponse {
	return groups
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type getRequest struct {
	Id int `json:"id" validate:"required"`
}
