package group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertGroups(t *testing.T, a, b *Group) {
	assert.Equal(t, a.Description, b.Description)
	assert.Equal(t, a.NamespaceId, b.NamespaceId)
	assert.Equal(t, a.Code, b.Code)
}
