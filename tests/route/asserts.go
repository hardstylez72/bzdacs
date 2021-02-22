package route

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertRoutes(t *testing.T, a, b *Route) {
	assert.Equal(t, a.Route, b.Route)
	assert.Equal(t, a.Method, b.Method)
	assert.Equal(t, a.Description, b.Description)
	assert.Equal(t, a.NamespaceId, b.NamespaceId)
	assert.EqualValues(t, a.Tags, b.Tags)
}
