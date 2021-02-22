package namespace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertNamespaces(t *testing.T, a, b *Namespace) {
	assert.Equal(t, a.Name, b.Name)
}
