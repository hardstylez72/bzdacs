package tag

import "github.com/brianvoe/gofakeit"

func GenTag(namespaceId int64) *Tag {
	gofakeit.Seed(0)
	return &Tag{
		Name:        gofakeit.Generate("###???"),
		NamespaceId: namespaceId,
	}
}
