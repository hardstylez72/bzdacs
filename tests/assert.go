package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Times struct {
	CreatedAt time.Time  `json:"createdAt" db:"created_at" swaggertype:"string"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at" swaggertype:"string"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
}

func AssertTimesUpdated(t *testing.T, times Times) {
	assert.NotZero(t, times.CreatedAt)
	assert.Nil(t, times.DeletedAt)
	assert.True(t, times.UpdatedAt.After(times.CreatedAt))
}

func AssertTimesCreated(t *testing.T, times Times) {
	assert.NotZero(t, times.CreatedAt)
	assert.Nil(t, times.DeletedAt)
	assert.True(t, times.UpdatedAt.Equal(times.CreatedAt))
}

func AssertTimesDeleted(t *testing.T, times Times) {
	assert.NotZero(t, times.CreatedAt)
	assert.NotZero(t, times.UpdatedAt)
	assert.True(t, times.DeletedAt.After(times.UpdatedAt))
}
