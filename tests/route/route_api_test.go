package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/namespace"
	"github.com/hardstylez72/bzdacs/tests/system"
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_tag_api(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), tests.GetConfig().Timeout)
	defer func() { _ = cancel }()
	c := tests.GetClient()

	t.Run("acceptance_test", func(t *testing.T) {
		acceptanceTest(ctx, t, c)
	})

	//t.Run("404", func(t *testing.T) {
	//	sysName := "404-h"
	//	_, err := GetNamespaceByName(ctx, c, sysName, 3523425)
	//	_, ok := err.(*namespace.NamespaceGetNotFound)
	//	if !ok {
	//		t.Fatal(err)
	//	}
	//})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) {

	s, err := system.Create(ctx, c, system.GenSystemName())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = system.Delete(ctx, c, s.ID)
	}()

	ns, err := namespace.Create(ctx, c, namespace.GenNamespaceName(), s.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = namespace.Delete(ctx, c, ns.ID, s.ID)
	}()

	r := GenRoute(ns.ID)
	created, err := Create(ctx, c, r)
	if err != nil {
		t.Fatal(err.Error())
	}

	assertRoutes(t, r, created)
	tests.AssertTimesCreated(t, created.Times)

	e := GenRoute(ns.ID)
	e.Id = created.Id
	edited, err := Update(ctx, c, e)
	if err != nil {
		t.Fatal(err.Error())
	}
	assertRoutes(t, e, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited.Id, ns.ID)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetById(ctx, c, edited.Id, ns.ID)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertTimesDeleted(t, deleted.Times)

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.Id, created.Id)
}
