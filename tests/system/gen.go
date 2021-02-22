package system

import "github.com/brianvoe/gofakeit"

func GenSystemName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}
