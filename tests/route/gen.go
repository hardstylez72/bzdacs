package route

import (
	"github.com/brianvoe/gofakeit"
)

func GenRoute(namespaceId int64) *Route {
	gofakeit.Seed(0)
	return &Route{
		Route:       gofakeit.URL(),
		Method:      gofakeit.HTTPMethod(),
		Description: gofakeit.JobDescriptor(),
		NamespaceId: namespaceId,
		Tags: []string{
			gofakeit.Generate("/api/????????/????/####"),
			gofakeit.Generate("/api/????????/????/####"),
			gofakeit.Generate("/api/????????/????/####"),
		},
	}
}
