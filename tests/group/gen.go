package group

import (
	"github.com/brianvoe/gofakeit"
	"github.com/hardstylez72/bzdacs/tests"
)

func GenGroup(namespaceId int64) *Group {
	gofakeit.Seed(0)
	return &Group{
		Code:        gofakeit.Job().Company,
		Description: gofakeit.Job().Descriptor,
		NamespaceId: namespaceId,
		Times:       tests.Times{},
	}
}
