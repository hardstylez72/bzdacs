package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests/system"
	"testing"
)

func SetupNamespace(ctx context.Context, t *testing.T, c *client.BZDACS) (*Namespace, error) {
	s, err := system.Create(ctx, c, system.GenSystem())
	if err != nil {
		t.Fatal(err)
	}

	ns, err := Create(ctx, c, GenNamespace(s.Id).Name, s.Id)
	if err != nil {
		t.Fatal(err)
	}
	return ns, nil
}

func TeardownNamespace(ctx context.Context, c *client.BZDACS, ns *Namespace) {
	_ = system.Delete(ctx, c, ns.SystemId)
	_ = Delete(ctx, c, ns.Id, ns.Id)
}
