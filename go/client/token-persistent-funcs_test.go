package client_test

import (
	"os"
	"path"
	"testing"

	"github.com/fi-ts/api/go/client"
	"github.com/stretchr/testify/require"
)

func TestNewFilesystemTokenPersiter(t *testing.T) {
	tokenPath := path.Join(t.TempDir(), "token")

	_, err := os.OpenFile(tokenPath, os.O_RDONLY|os.O_CREATE, 0600)
	require.NoError(t, err)

	filesystemPersister, err := client.NewFilesystemTokenPersister(tokenPath)
	require.NoError(t, err)

	token := "a-token"

	err = filesystemPersister(token)
	require.NoError(t, err)
	tokenBytes, err := os.ReadFile(tokenPath)
	require.NoError(t, err)
	require.Equal(t, token, string(tokenBytes))
}
