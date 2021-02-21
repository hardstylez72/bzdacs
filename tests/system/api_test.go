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
		err := acceptanceTest(ctx, t, c)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("404", func(t *testing.T) {
		sysName := "404-h"
		_, err := GetSystemByName(ctx, c, sysName)
		_, ok := err.(*system.SystemGetNotFound)
		if !ok {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) error {
	sysName := tests.GenSystemName()
	created, err := CreateSystem(ctx, c, sysName)
	if err != nil {
		return err
	}

	newSysName := tests.GenSystemName()
	edited, err := UpdateSystem(ctx, c, newSysName, created.ID)
	if err != nil {
		return err
	}

	err = DeleteSystem(ctx, c, edited.ID)
	if err != nil {
		return err
	}

	deleted, err := GetSystemById(ctx, c, edited.ID)
	if err != nil {
		return err
	}

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.ID, created.ID)
	return nil
}
