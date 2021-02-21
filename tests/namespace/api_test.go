package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/namespace"

	"github.com/hardstylez72/bzdacs/tests"
	"github.com/hardstylez72/bzdacs/tests/system"
	"github.com/stretchr/testify/assert"

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
		_, err := GetNamespaceByName(ctx, c, sysName, 3523425)
		_, ok := err.(*namespace.NamespaceGetNotFound)
		if !ok {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) {

	s, err := system.CreateSystem(ctx, c, tests.GenSystemName())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = system.DeleteSystem(ctx, c, s.ID)
	}()
	nsName := tests.GenNamespaceName()
	created, err := CreateNamespace(ctx, c, nsName, s.ID)
	if err != nil {
		t.Fatal(err)
	}

	newNsName := tests.GenNamespaceName()
	edited, err := UpdateNamespace(ctx, c, newNsName, created.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteNamespace(ctx, c, edited.ID, s.ID)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := GetNamespaceById(ctx, c, edited.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.ID, created.ID)
}
