package tests

import "github.com/brianvoe/gofakeit"

func GenSystemName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}

func GenNamespaceName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}

func GenTagName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}
