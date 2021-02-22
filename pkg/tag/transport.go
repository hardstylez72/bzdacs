package tag

type insertRequest struct {
	Name        string `json:"name" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
}

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
}

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

type listResponse []Tag

func newListResponse(tags []Tag) listResponse {
	return tags
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
}

type getRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
}

type getResponse *Tag

func newGetResponse(tag *Tag) getResponse {
	return tag
}

type listRequest struct {
	Pattern     string `json:"pattern"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
}

func assertUpdatesEntity() {

}
