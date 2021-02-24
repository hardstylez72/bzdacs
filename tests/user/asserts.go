package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertUsers(t *testing.T, a, b *User) {
	assert.Equal(t, a.Login, b.Login)
	assert.Equal(t, a.NamespaceId, b.NamespaceId)
}
