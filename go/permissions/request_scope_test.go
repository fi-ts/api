package permissions

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/emptypb"

	// Register the request types of all services in GetServicePermissions.
	// A new API package must be added here, otherwise the tests below fail
	// because its services cannot be found.
	_ "github.com/fi-ts/api/go/fits/api/mvm/v1"
	_ "github.com/fi-ts/api/go/fits/api/pcvm/v1"
	_ "github.com/fi-ts/api/go/fits/api/v1"
)

var (
	projectFields = []string{"project_uuid", "project"}
	tenantFields  = []string{"tenant"}
)

// TestProjectScopedRequestsExposeProject checks that the authorization layer
// can resolve the project of every project scoped request.
//
// Authorization only finds the project through a GetProjectUuid (or
// GetProject) getter on the top-level request. A request that carries its
// project only in a nested message (e.g. a Validate* request wrapping the real
// request) needs a hand-written getter in a *_project_scope.go file. Without
// it, every call is rejected with "no project found in project scoped
// request". The request then fails closed, but the endpoint is unusable.
//
// The test also checks that the getter returns the project field of the
// request, so the project that is authorized is the one the handler reads.
func TestProjectScopedRequestsExposeProject(t *testing.T) {
	checkScopedRequests(t, GetServicePermissions().Visibility.Project, projectFields, getProjectFromMessage)
}

// TestTenantScopedRequestsExposeTenant is the tenant counterpart of
// TestProjectScopedRequestsExposeProject.
func TestTenantScopedRequestsExposeTenant(t *testing.T) {
	checkScopedRequests(t, GetServicePermissions().Visibility.Tenant, tenantFields, getTenantFromMessage)
}

// TestScopeLookupWithoutGetter makes sure the lookups report a message
// without a getter, which is what the tests above rely on.
func TestScopeLookupWithoutGetter(t *testing.T) {
	if _, ok := getProjectFromMessage(&emptypb.Empty{}); ok {
		t.Error("getProjectFromMessage found a project in a message without a project getter")
	}
	if _, ok := getTenantFromMessage(&emptypb.Empty{}); ok {
		t.Error("getTenantFromMessage found a tenant in a message without a tenant getter")
	}
}

func checkScopedRequests(t *testing.T, methods map[string]bool, fieldNames []string, lookup func(any) (string, bool)) {
	t.Helper()
	if len(methods) == 0 {
		t.Fatal("no scoped methods found")
	}

	for _, method := range slices.Sorted(maps.Keys(methods)) {
		t.Run(method, func(t *testing.T) {
			msg, err := requestMessage(method)
			if err != nil {
				t.Fatal(err)
			}
			name := msg.Descriptor().FullName()

			paths := scopeFieldPaths(msg.Descriptor(), fieldNames)
			if len(paths) != 1 {
				t.Fatalf("request %s must have exactly one %v field (top-level or nested), found %d: %v", name, fieldNames, len(paths), paths)
			}

			const want = "scope-from-request"
			setPath(msg, paths[0], want)

			got, ok := lookup(msg.Interface())
			if !ok {
				t.Fatalf("request %s has its scope in %s but no getter the authorization layer can use; add one in a *_project_scope.go file next to the generated code", name, paths[0])
			}
			if got != want {
				t.Errorf("getter of request %s returns %q, want the value of %s (%q)", name, got, paths[0], want)
			}
		})
	}
}

// requestMessage returns a new, empty request message of a method in the
// form "/fits.api.mvm.v1.MVMService/Get".
func requestMessage(method string) (protoreflect.Message, error) {
	service, name, ok := strings.Cut(strings.TrimPrefix(method, "/"), "/")
	if !ok {
		return nil, fmt.Errorf("malformed method %q", method)
	}
	d, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(service))
	if err != nil {
		return nil, fmt.Errorf("service %s is not registered, import its Go package in this test: %w", service, err)
	}
	sd, ok := d.(protoreflect.ServiceDescriptor)
	if !ok {
		return nil, fmt.Errorf("%s is not a service", service)
	}
	md := sd.Methods().ByName(protoreflect.Name(name))
	if md == nil {
		return nil, fmt.Errorf("service %s has no method %s", service, name)
	}
	mt, err := protoregistry.GlobalTypes.FindMessageByName(md.Input().FullName())
	if err != nil {
		return nil, fmt.Errorf("request type of %s is not registered: %w", method, err)
	}
	return mt.New(), nil
}

// fieldPath is the path from a request to one of its (possibly nested) fields.
type fieldPath []protoreflect.FieldDescriptor

func (p fieldPath) String() string {
	names := make([]string, 0, len(p))
	for _, fd := range p {
		names = append(names, string(fd.Name()))
	}
	return strings.Join(names, ".")
}

// scopeFieldPaths returns the paths to all string fields with one of the given
// names, reachable from md through singular message fields. Repeated and map
// fields are skipped: authorization can only check a single subject.
func scopeFieldPaths(md protoreflect.MessageDescriptor, names []string) []fieldPath {
	var (
		paths []fieldPath
		seen  = map[protoreflect.FullName]bool{}
		walk  func(md protoreflect.MessageDescriptor, prefix fieldPath)
	)
	walk = func(md protoreflect.MessageDescriptor, prefix fieldPath) {
		if seen[md.FullName()] {
			return
		}
		seen[md.FullName()] = true
		defer delete(seen, md.FullName())

		fields := md.Fields()
		for i := range fields.Len() {
			fd := fields.Get(i)
			if fd.IsList() || fd.IsMap() {
				continue
			}
			path := append(slices.Clone(prefix), fd)
			switch {
			case fd.Kind() == protoreflect.StringKind && slices.Contains(names, string(fd.Name())):
				paths = append(paths, path)
			case fd.Kind() == protoreflect.MessageKind:
				walk(fd.Message(), path)
			}
		}
	}
	walk(md, nil)
	return paths
}

// setPath sets the field at path to v, creating the nested messages on the way.
func setPath(m protoreflect.Message, path fieldPath, v string) {
	for _, fd := range path[:len(path)-1] {
		m = m.Mutable(fd).Message()
	}
	m.Set(path[len(path)-1], protoreflect.ValueOfString(v))
}
