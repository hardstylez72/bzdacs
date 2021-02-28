package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/system"
	"testing"
)

func Test_namespace_api(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), tests.GetConfig().Timeout)
	defer func() { _ = cancel }()
	c := tests.GetClient()

	t.Run("acceptance_test", func(t *testing.T) {
		acceptanceTest(ctx, t, c)
	})

	t.Run("404", func(t *testing.T) {
		sysName := "404-h"
		_, err := GetByName(ctx, c, sysName, 3523425)
		if err != ErrNamespaceGetNotFound {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) {

	s, err := system.Create(ctx, c, system.GenSystem())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = system.Delete(ctx, c, s.Id)
	}()
	ns := GenNamespace(s.Id)
	created, err := Create(ctx, c, ns.Name, ns.SystemId)
	if err != nil {
		t.Fatal(err)
	}
	assertNamespaces(t, created, ns)
	tests.AssertTimesCreated(t, created.Times)

	newNs := GenNamespace(s.Id)
	newNs.Id = created.Id
	edited, err := Update(ctx, c, newNs.Name, newNs.Id)
	if err != nil {
		t.Fatal(err)
	}

	assertNamespaces(t, newNs, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited.Id, edited.Id)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetById(ctx, c, edited.Id)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertTimesDeleted(t, deleted.Times)
}
