package user

type insertRequest struct {
	ExternalId string `json:"externalId" validate:"required"`
}

func insertRequestConvert(r *insertRequest) *User {
	return &User{
		ExternalId: r.ExternalId,
	}
}

type insertResponse User

func newInsertResponse(group *User) *insertResponse {
	return (*insertResponse)(group)
}

type listResponse []User

func newListResponse(groups []User) listResponse {
	return groups
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}
type getRequest struct {
	Id int `json:"id" validate:"required"`
}
