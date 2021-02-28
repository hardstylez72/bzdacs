package tag

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

	//t.Run("404", func(t *testing.T) {
	//	sysName := "404-h"
	//	_, err := GetNamespaceByName(ctx, c, sysName, 3523425)
	//	_, ok := err.(*namespace.NamespaceGetNotFound)
	//	if !ok {
	//		t.Fatal(err)
	//	}
	//})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS, ns *namespace.Namespace) {

	tag := GenTag(ns.Id)
	created, err := Create(ctx, c, tag)
	if err != nil {
		t.Fatal(err)
	}
	assertTags(t, created, tag)
	tests.AssertTimesCreated(t, created.Times)

	newTag := GenTag(ns.Id)
	newTag.Id = created.Id
	edited, err := Update(ctx, c, newTag)
	if err != nil {
		t.Fatal(err)
	}
	assertTags(t, newTag, edited)
	tests.AssertTimesUpdated(t, edited.Times)

	err = Delete(ctx, c, edited)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetById(ctx, c, edited.Id)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertTimesDeleted(t, deleted.Times)

}
