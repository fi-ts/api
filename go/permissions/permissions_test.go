package permissions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetServicePermissions(t *testing.T) {
	perms := GetServicePermissions()
	require.NotNil(t, perms)
	// TODO more coverage
	require.Contains(t, perms.Methods, "/fits.api.v1.IPService/List")
	require.Contains(t, perms.Visibility.Self, "/fits.api.v1.TokenService/Create")
	require.Contains(t, perms.Visibility.Project, "/fits.api.v1.IPService/List")
}
