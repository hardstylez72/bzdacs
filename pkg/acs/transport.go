package acs

import "github.com/hardstylez72/bzdacs/pkg/usergroup"

type checkAccessResponse struct {
	Login         string   `json:"login"`
	AccessAllowed bool     `json:"accessAllowed"`
	Groups        []string `json:"groups"`
}

func newCheckAccessResponse(groups []usergroup.Group, isAllowed bool, login string) *checkAccessResponse {
	groupNames := make([]string, 0)

	for _, group := range groups {
		groupNames = append(groupNames, group.Code)
	}

	return &checkAccessResponse{
		Login:         login,
		AccessAllowed: isAllowed,
		Groups:        groupNames,
	}
}
