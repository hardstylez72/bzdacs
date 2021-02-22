package tag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertTags(t *testing.T, a, b *Tag) {
	assert.Equal(t, a.Name, b.Name)
	assert.Equal(t, a.NamespaceId, b.NamespaceId)
}
