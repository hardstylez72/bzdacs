package system

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests"
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
		if err != ErrSystemGetNotFound {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) {
	s := GenSystem()
	created, err := Create(ctx, c, s)
	if err != nil {
		t.Fatal(err)
	}
	assertSystems(t, created, s)
	tests.AssertTimesCreated(t, created.Times)

	newSysName := GenSystem()
	newSysName.Id = created.Id
	edited, err := Update(ctx, c, newSysName)
	if err != nil {
		t.Fatal(err)
	}

	assertSystems(t, newSysName, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited.Id)
	if err != nil {
		t.Fatal(err)
	}
	deleted, err := GetById(ctx, c, edited.Id)
	if err != nil {
		t.Fatal(err)
	}

	tests.AssertTimesDeleted(t, deleted.Times)
}
