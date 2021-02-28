package user

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/namespace"
	"testing"
)

func Test_user_api(t *testing.T) {

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
		r := GenUser(ns.Id)
		_, err := GetByLogin(ctx, c, r)
		if err != ErrNotFound {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS, ns *namespace.Namespace) {

	r := GenUser(ns.Id)
	created, err := Create(ctx, c, r)
	if err != nil {
		t.Fatal(err.Error())
	}

	assertUsers(t, r, created)
	tests.AssertTimesCreated(t, created.Times)

	e := GenUser(ns.Id)
	e.Id = created.Id
	edited, err := Update(ctx, c, e)
	if err != nil {
		t.Fatal(err.Error())
	}
	assertUsers(t, e, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited.Id, ns.Id)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetById(ctx, c, edited.Id, ns.Id)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertTimesDeleted(t, deleted.Times)
}
