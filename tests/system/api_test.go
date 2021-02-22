package system

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/system"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_system_api(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), tests.GetConfig().Timeout)
	defer func() { _ = cancel }()
	c := tests.GetClient()

	t.Run("acceptance_test", func(t *testing.T) {
		acceptanceTest(ctx, t, c)
	})

	t.Run("404", func(t *testing.T) {
		sysName := "404-h"
		_, err := GetByName(ctx, c, sysName)
		_, ok := err.(*system.SystemGetNotFound)
		if !ok {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) {
	sysName := GenSystemName()
	created, err := Create(ctx, c, sysName)
	if err != nil {
		t.Fatal(err)
	}

	newSysName := GenSystemName()
	edited, err := Update(ctx, c, newSysName, created.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = Delete(ctx, c, edited.ID)
	if err != nil {
		t.Fatal(err)
	}
	deleted, err := GetById(ctx, c, edited.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.ID, created.ID)
}
