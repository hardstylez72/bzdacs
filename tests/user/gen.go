package user

import (
	"github.com/brianvoe/gofakeit"
)

func GenUser(namespaceId int64) *User {
	gofakeit.Seed(0)
	return &User{
		Login:       gofakeit.Name(),
		NamespaceId: namespaceId,
	}
}
