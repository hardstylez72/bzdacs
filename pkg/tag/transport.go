package tag

type insertRequest struct {
	Name        string `json:"name" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name tagInsertRequest

func insertRequestConvert(r *insertRequest) *Tag {
	return &Tag{
		Name:        r.Name,
		NamespaceId: r.NamespaceId,
	}
}

type insertResponse *Tag

func newInsertResponse(tag *Tag) insertResponse {
	return tag
}

type updateRequest struct {
	Id          int    `json:"id"  validate:"required"`
	Name        string `json:"name" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name tagUpdateRequest

func updateRequestConvert(r *updateRequest) *Tag {
	return &Tag{
		Id:          r.Id,
		Name:        r.Name,
		NamespaceId: r.NamespaceId,
	}
}

type updateResponse *Tag

func newUpdateResponse(tag *Tag) updateResponse {
	return tag
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name tagDeleteRequest

type getRequest struct {
	Id int `json:"id" validate:"required"`
} // @name tagGetRequest

type getResponse *Tag // @name tagGetResponse

func newGetResponse(tag *Tag) getResponse {
	return tag
}

type listRequest struct {
	Filter filter `json:"filter"`
} // @name tagListRequest

type filter struct {
	Page        int    `json:"page"  validate:"required,gte=1"`
	PageSize    int    `json:"pageSize"  validate:"required,gte=1"`
	Pattern     string `json:"pattern"`
	NamespaceId int    `json:"namespaceId" validate:"required,gte=1"`
}

type listResponse struct {
	Items []getResponse `json:"items"`
	Total int           `json:"total"`
} // @name tagListResponse

func newListResponse(tags []Tag, total int) listResponse {
	out := make([]getResponse, 0, len(tags))
	for i := range tags {
		out = append(out, &tags[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}
