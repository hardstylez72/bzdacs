package tag

import "github.com/brianvoe/gofakeit"

func GenTagName() string {
	gofakeit.Seed(0)
	return gofakeit.Generate("###???")
}
