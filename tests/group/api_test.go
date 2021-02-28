package group

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/namespace"
	"testing"
)

func Test_tag_api(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), tests.GetConfig().Timeout)
	defer func() { _ = cancel }()
	c := tests.GetClient()

	ns, err := namespace.SetupNamespace(ctx, t, c)
	if err != nil {
		t.Fatal(err)
	}
	defer namespace.TeardownNamespace(ctx, c, ns)

	t.Run("acceptance_test", func(t *testing.T) {
		acceptanceTest(ctx, t, c, ns)
	})

	t.Run("404", func(t *testing.T) {
		r := GenGroup(ns.Id)
		_, err := GetByParams(ctx, c, r)
		if err != ErrNotFound {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS, ns *namespace.Namespace) {

	r := GenGroup(ns.Id)
	created, err := Create(ctx, c, r)
	if err != nil {
		t.Fatal(err.Error())
	}

	assertGroups(t, r, created)
	tests.AssertTimesCreated(t, created.Times)

	e := GenGroup(ns.Id)
	e.Id = created.Id
	edited, err := Update(ctx, c, e)
	if err != nil {
		t.Fatal(err.Error())
	}
	assertGroups(t, e, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited.Id, ns.Id)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetById(ctx, c, edited.Id, ns.Id)
	if err != nil {
		t.Fatal(err)
	}
	assertGroups(t, deleted, edited)
	tests.AssertTimesDeleted(t, deleted.Times)
}
