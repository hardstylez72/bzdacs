package namespace

import "github.com/brianvoe/gofakeit"

func GenNamespace(systemId int64) *Namespace {
	gofakeit.Seed(0)
	return &Namespace{
		Name:     gofakeit.Generate("###???"),
		SystemId: systemId,
	}
}
