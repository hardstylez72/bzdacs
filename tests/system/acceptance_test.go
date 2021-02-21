package system

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/client/system"
	"github.com/hardstylez72/bzdacs/models"
	"github.com/hardstylez72/bzdacs/tests"
	"github.com/stretchr/testify/assert"

	"net/http"
	"testing"
)

func Test_system_api(t *testing.T) {
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
		sysName := genSystemName()
		_, err := GetSystemByName(ctx, c, sysName)
		_, ok := err.(*system.SystemGetNotFound)
		if !ok {
			t.Fatal(err)
		}
	})
}

func acceptanceTest(ctx context.Context, t *testing.T, c *client.BZDACS) error {
	sysName := genSystemName()
	created, err := CreateSystem(ctx, c, sysName)
	if err != nil {
		return err
	}

	newSysName := genSystemName()
	edited, err := UpdateSystem(ctx, c, newSysName, created.ID)
	if err != nil {
		return err
	}

	err = DeleteSystem(ctx, c, edited.ID)
	if err != nil {
		return err
	}

	deleted, err := GetSystemById(ctx, c, edited.ID)
	if err != nil {
		return err
	}

	assert.NotNil(t, deleted.DeletedAt)
	assert.Equal(t, deleted.ID, created.ID)
	return nil
}

func DeleteSystem(ctx context.Context, client *client.BZDACS, id int64) error {
	_, err := client.System.SystemDelete(
		&system.SystemDeleteParams{
			Req:        &models.SystemDeleteRequest{ID: &id},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetSystemByName(ctx context.Context, client *client.BZDACS, name string) (*models.SystemGetResponse, error) {
	res, err := client.System.SystemGet(
		&system.SystemGetParams{
			Req: &models.SystemGetRequest{
				Name: name,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func GetSystemById(ctx context.Context, client *client.BZDACS, id int64) (*models.SystemGetResponse, error) {
	res, err := client.System.SystemGet(
		&system.SystemGetParams{
			Req: &models.SystemGetRequest{
				ID: id,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func UpdateSystem(ctx context.Context, client *client.BZDACS, name string, id int64) (*models.SystemUpdateResponse, error) {
	res, err := client.System.SystemUpdate(
		&system.SystemUpdateParams{
			Req: &models.SystemUpdateRequest{
				ID:   &id,
				Name: &name,
			},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func CreateSystem(ctx context.Context, client *client.BZDACS, name string) (*models.SystemInsertResponse, error) {
	res, err := client.System.SystemCreate(
		&system.SystemCreateParams{
			Req:        &models.SystemInsertRequest{Name: &name},
			Context:    ctx,
			HTTPClient: &http.Client{},
		},
	)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}

func genSystemName() string {
	return gofakeit.DomainName()
}
