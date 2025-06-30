package kms_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	kms "github.com/sacloud/kms-api-go"
	v1 "github.com/sacloud/kms-api-go/apis/v1"
	"github.com/sacloud/packages-go/testutil"
)

func TestKeyAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	client, err := kms.NewClient()
	require.NoError(t, err, "failed to create client")

	ctx := context.Background()
	keyOp := kms.NewKeyOp(client)

	resCreate, err := keyOp.Create(ctx, v1.CreateKey{
		Name:        "key gen from go",
		Description: v1.NewOptString("key gen from go client"),
		KeyOrigin:   v1.KeyOriginEnumGenerated,
		Tags:        []string{"tag1", "tag2"},
	})
	require.NoError(t, err, "failed to create key")
	assert.Equal(t, "key gen from go", resCreate.Name)

	resList, err := keyOp.List(ctx)
	assert.NoError(t, err, "failed to list keys")

	found := false
	for _, key := range resList {
		if key.ID == resCreate.ID {
			found = true
			assert.Equal(t, "key gen from go client", key.Description.Value)
		}
	}
	assert.True(t, found, "created key not found in list")

	_, err = keyOp.Update(ctx, resCreate.ID, v1.Key{
		Name:        "key gen from go 2",
		Description: v1.NewOptString("key gen from go client 2"),
		KeyOrigin:   v1.KeyOriginEnumGenerated,
		Tags:        []string{"Test"},
	})
	assert.NoError(t, err, "failed to update key")

	resRead, err := keyOp.Read(ctx, resCreate.ID)
	assert.NoError(t, err, "failed to read key")
	assert.Equal(t, "key gen from go 2", resRead.Name)
	assert.Equal(t, "key gen from go client 2", resRead.Description.Value)

	err = keyOp.Delete(ctx, resCreate.ID)
	require.NoError(t, err, "failed to delete key")
}
