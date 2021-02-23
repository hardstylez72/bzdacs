package group

type insertRequest struct {
	BaseGroupId int    `json:"baseGroupId"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name groupInsertRequest

func insertRequestConvert(r *insertRequest) *Group {
	return &Group{
		Code:        r.Code,
		Description: r.Description,
		NamespaceId: r.NamespaceId,
	}
}

type getResponse *Group // @name groupGetResponse

func newInsertResponse(group *Group) getResponse {
	return group
}

type updateRequest struct {
	Id          int    `json:"id"  validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name groupUpdateRequest

func updateRequestConvert(r *updateRequest) *Group {
	return &Group{
		Id:          r.Id,
		Code:        r.Code,
		Description: r.Description,
		NamespaceId: r.NamespaceId,
	}
}

func newUpdateResponse(group *Group) getResponse {
	return group
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name groupDeleteRequest

type getByIdRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name groupGetByIdRequest

type getByCodeRequest struct {
	Code        string `json:"code" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name groupGetByCodeRequest

func newGetResponse(group *Group) getResponse {
	return group
}

type listRequest struct {
	Filter filter `json:"filter"`
} // @name groupListRequest

type filter struct {
	Page        int    `json:"page"  validate:"required,gte=1"`
	PageSize    int    `json:"pageSize"  validate:"required,gte=1"`
	Pattern     string `json:"pattern"`
	NamespaceId int    `json:"namespaceId" validate:"required,gte=1"`
}

type listResponse struct {
	Items []getResponse `json:"items" validate:"required"`
	Total int           `json:"total" validate:"required"`
} // @name groupListResponse

func newListResponse(tags []Group, total int) listResponse {
	out := make([]getResponse, 0, len(tags))
	for i := range tags {
		out = append(out, &tags[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}
