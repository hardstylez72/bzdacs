package namespace

import "github.com/brianvoe/gofakeit"

func GenNamespaceName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}
