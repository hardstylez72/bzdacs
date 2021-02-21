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
		err := acceptanceTest(ctx, t, c)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("404", func(t *testing.T) {
		sysName := "404-h"
		_, err := GetNamespaceByName(ctx, c, sysName)
		_, ok := err.(*namespace.NamespaceGetNotFound)
		if !ok {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) error {

	s, err := system.CreateSystem(ctx, c, tests.GenSystemName())
	if err != nil {
		return err
	}
	defer func() {
		_ = system.DeleteSystem(ctx, c, s.ID)
	}()
	nsName := tests.GenNamespaceName()
	created, err := CreateNamespace(ctx, c, nsName, s.ID)
	if err != nil {
		return err
	}

	newNsName := tests.GenNamespaceName()
	edited, err := UpdateNamespace(ctx, c, newNsName, created.ID)
	if err != nil {
		return err
	}

	err = DeleteNamespace(ctx, c, edited.ID, s.ID)
	if err != nil {
		return err
	}

	deleted, err := GetNamespaceById(ctx, c, edited.ID)
	if err != nil {
		return err
	}

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.ID, created.ID)
	return nil
}
