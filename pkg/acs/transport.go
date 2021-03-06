package acs

type checkAccessResponse struct {
	Login         string `json:"login"`
	AccessAllowed bool   `json:"accessAllowed"`
}

func newCheckAccessResponse(isAllowed bool, login string) *checkAccessResponse {
	return &checkAccessResponse{
		Login:         login,
		AccessAllowed: isAllowed,
	}
}
