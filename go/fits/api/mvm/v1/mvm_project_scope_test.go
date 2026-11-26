package mvmv1

import "testing"

// projectGetter mirrors the interface the authorization layer type-asserts
// against in permissions.GetProjectFromRequest. If a Validate* request stops
// satisfying it, project-scoped authorization breaks with "no project found in
// project scoped request", so we assert it here at compile and run time.
type projectGetter interface {
	GetProjectUuid() string
}

func TestValidateRequestsExposeProjectUuid(t *testing.T) {
	const project = "11111111-1111-1111-1111-111111111111"

	tests := []struct {
		name string
		req  projectGetter
		want string
	}{
		{
			name: "validate create forwards nested project",
			req: &MVMServiceValidateCreateRequest{
				Create: &MVMServiceCreateRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate add disk forwards nested project",
			req: &MVMServiceValidateAddDisksRequest{
				AddDisks: &MVMServiceAddDisksRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate update disk forwards nested project",
			req: &MVMServiceValidateUpdateDiskRequest{
				UpdateDisk: &MVMServiceUpdateDiskRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate reinstall forwards nested project",
			req: &MVMServiceValidateReinstallRequest{
				Reinstall: &MVMServiceReinstallRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate update forwards nested project",
			req: &MVMServiceValidateUpdateRequest{
				Update: &MVMServiceUpdateRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate update service class forwards nested project",
			req: &MVMServiceValidateUpdateServiceClassRequest{
				UpdateServiceClass: &MVMServiceUpdateServiceClassRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "on behalf get forwards nested project",
			req: &OnBehalfMVMServiceGetRequest{
				Request: &MVMServiceGetRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "on behalf list forwards nested project",
			req: &OnBehalfMVMServiceListRequest{
				Request: &MVMServiceListRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "on behalf create forwards nested project",
			req: &OnBehalfMVMServiceCreateRequest{
				Request: &MVMServiceCreateRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "on behalf delete forwards nested project",
			req: &OnBehalfMVMServiceDeleteRequest{
				Request: &MVMServiceDeleteRequest{ProjectUuid: project},
			},
			want: project,
		},
		{
			name: "validate create without nested request yields empty",
			req:  &MVMServiceValidateCreateRequest{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.GetProjectUuid(); got != tt.want {
				t.Errorf("GetProjectUuid() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidateRequestsNilReceiver(t *testing.T) {
	var (
		create             *MVMServiceValidateCreateRequest
		addDisks           *MVMServiceValidateAddDisksRequest
		updateDisk         *MVMServiceValidateUpdateDiskRequest
		reinstall          *MVMServiceValidateReinstallRequest
		update             *MVMServiceValidateUpdateRequest
		updateServiceClass *MVMServiceValidateUpdateServiceClassRequest
		onBehalfGet        *OnBehalfMVMServiceGetRequest
		onBehalfList       *OnBehalfMVMServiceListRequest
		onBehalfCreate     *OnBehalfMVMServiceCreateRequest
		onBehalfDelete     *OnBehalfMVMServiceDeleteRequest
	)
	// A nil receiver must not panic; it simply has no project.
	for _, got := range []string{
		create.GetProjectUuid(),
		addDisks.GetProjectUuid(),
		updateDisk.GetProjectUuid(),
		reinstall.GetProjectUuid(),
		update.GetProjectUuid(),
		updateServiceClass.GetProjectUuid(),
		onBehalfGet.GetProjectUuid(),
		onBehalfList.GetProjectUuid(),
		onBehalfCreate.GetProjectUuid(),
		onBehalfDelete.GetProjectUuid(),
	} {
		if got != "" {
			t.Errorf("GetProjectUuid() on nil receiver = %q, want empty", got)
		}
	}
}
