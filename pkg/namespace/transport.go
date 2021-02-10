package namespace

type insertRequest struct {
	Name string `json:"name" validate:"required"`
}

func insertRequestConvert(r *insertRequest) *Namespace {
	return &Namespace{
		Name: r.Name,
	}
}

type insertResponse Namespace

func newInsertResponse(group *Namespace) *insertResponse {
	return (*insertResponse)(group)
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

func newUpdateResponse(group *Namespace) *updateResponse {
	return (*updateResponse)(group)
}

type listResponse []Namespace

func newListResponse(groups []Namespace) listResponse {
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