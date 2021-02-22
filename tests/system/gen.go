package system

import (
	"github.com/brianvoe/gofakeit"
)

func GenSystem() *System {
	gofakeit.Seed(0)
	return &System{
		Name: gofakeit.Generate("###???"),
	}
}
