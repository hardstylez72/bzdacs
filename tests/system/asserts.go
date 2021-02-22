package system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertSystems(t *testing.T, a, b *System) {
	assert.Equal(t, a.Name, b.Name)
}
