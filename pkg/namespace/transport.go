package namespace

type insertRequest struct {
	Name     string `json:"name" validate:"required"`
	SystemId int    `json:"systemId" validate:"required"`
}

func insertRequestConvert(r *insertRequest) *Namespace {
	return &Namespace{
		Name:     r.Name,
		SystemId: r.SystemId,
	}
}

type insertResponse Namespace

func newInsertResponse(namespace *Namespace) *insertResponse {
	return (*insertResponse)(namespace)
}

type updateRequest struct {
	Id   int    `json:"id"  validate:"required"`
	Name string `json:"name" validate:"required"`
}

func updateRequestConvert(r *updateRequest) *Namespace {
	return &Namespace{
		Id:   r.Id,
		Name: r.Name,
	}
}

type updateResponse Namespace

func newUpdateResponse(namespace *Namespace) *updateResponse {
	return (*updateResponse)(namespace)
}

type listResponse []Namespace

func newListResponse(namespaces []Namespace) listResponse {
	return namespaces
}

type deleteRequest struct {
	NamespaceId int `json:"namespaceId" validate:"required"`
}

type getRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	SystemId int    `json:"systemId"`
}

type getResponse *Namespace

func newGetResponse(namespace *Namespace) getResponse {
	return namespace
}

type listRequest struct {
	Id int `json:"id" validate:"required"`
}
