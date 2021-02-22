package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/namespace"
	"github.com/hardstylez72/bzdacs/tests/system"
	"testing"
)

func Test_tag_api(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), tests.GetConfig().Timeout)
	defer func() { _ = cancel }()
	c := tests.GetClient()

	s, err := system.Create(ctx, c, system.GenSystem())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = system.Delete(ctx, c, s.Id)
	}()

	ns, err := namespace.Create(ctx, c, namespace.GenNamespace(s.Id).Name, s.Id)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = namespace.Delete(ctx, c, ns.Id, s.Id)
	}()

	t.Run("acceptance_test", func(t *testing.T) {
		acceptanceTest(ctx, t, c, ns)
	})

	t.Run("404", func(t *testing.T) {
		r := GenRoute(ns.Id)
		_, err := GetByParams(ctx, c, r)
		if err != ErrNotFound {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS, ns *namespace.Namespace) {

	r := GenRoute(ns.Id)
	created, err := Create(ctx, c, r)
	if err != nil {
		t.Fatal(err.Error())
	}

	assertRoutes(t, r, created)
	tests.AssertTimesCreated(t, created.Times)

	e := GenRoute(ns.Id)
	e.Id = created.Id
	edited, err := Update(ctx, c, e)
	if err != nil {
		t.Fatal(err.Error())
	}
	assertRoutes(t, e, edited)
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
