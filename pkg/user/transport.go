package user

type insertRequest struct {
	ExternalId  string `json:"externalId" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name userCreateRequest

func insertRequestConvert(r *insertRequest) *User {
	return &User{
		ExternalId:  r.ExternalId,
		NamespaceId: r.NamespaceId,
	}
}

type insertResponse *User

type getResponse *User // @name userGetResponse

func newInsertResponse(u *User) insertResponse {
	return u
}

type deleteRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name userDeleteRequest

type getByLoginRequest struct {
	Login       string `json:"login" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name userGetByLoginRequest

type getByIdRequest struct {
	Id          int `json:"id" validate:"required"`
	NamespaceId int `json:"namespaceId" validate:"required"`
} // @name userGetByIdRequest

type updateRequest struct {
	Id          int    `json:"id" validate:"required"`
	ExternalId  string `json:"externalId" validate:"required"`
	NamespaceId int    `json:"namespaceId" validate:"required"`
} // @name userUpdateRequest

func updateRequestConvert(r *updateRequest) *User {
	return &User{
		ExternalId:  r.ExternalId,
		Id:          r.Id,
		NamespaceId: r.NamespaceId,
	}
}

type listResponse struct {
	Items []getResponse `json:"items" validate:"required"`
	Total int           `json:"total" validate:"required"`
} // @name userListResponse

func newListResponse(routes []User, total int) listResponse {
	out := make([]getResponse, 0, len(routes))
	for i := range routes {
		out = append(out, &routes[i])
	}
	return listResponse{
		Items: out,
		Total: total,
	}
}

type listRequest struct {
	Filter filter `json:"filter"`
} // @name userListRequest

type filter struct {
	Page        int    `json:"page"  validate:"required,gte=1"`
	PageSize    int    `json:"pageSize"`
	NamespaceId int    `json:"namespaceId" validate:"required,gte=1"`
	Pattern     string `json:"pattern"`
}
