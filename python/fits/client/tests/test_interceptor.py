"""Tests mirroring go/client/test_interceptor_test.go.

Tests the interceptor pattern using connectrpc's MetadataInterceptorSync,
which is the Python equivalent of Go's TestClientInterceptor.
"""

import pyqwest
from connectrpc.request import RequestContext
from pyqwest.testing import WSGITransport

from fits.api.v1 import version_pb2
from fits.api.v1.version_connect import (
    VersionServiceClientSync,
    VersionServiceWSGIApplication,
)

from .conftest import MockVersionService, _build_combined_wsgi_app


class AuthInterceptor:
    """Client-side interceptor that injects an Authorization header.

    Mirrors the authInterceptor from go/client/client-interceptors.go.
    """

    def __init__(self, token: str):
        self.token = token

    def on_start_sync(self, ctx: RequestContext) -> None:
        ctx.request_headers().add("authorization", f"Bearer {self.token}")

    def on_end_sync(self, token, ctx: RequestContext, error) -> None:
        pass


class RequestRecordingInterceptor:
    """Interceptor that records the received request for assertion.

    Mirrors the TestClientInterceptor from go/client/test_interceptor.go.
    """

    def __init__(self):
        self.calls: list[tuple[str, str]] = []

    def on_start_sync(self, ctx: RequestContext) -> None:
        method = ctx.method()
        self.calls.append(
            (method.service_name, method.name)
        )

    def on_end_sync(self, token, ctx: RequestContext, error) -> None:
        pass


class TestInterceptor:
    """Mirrors TestInterceptor from test_interceptor_test.go.

    Tests that an interceptor can intercept a call and verify request/response flow.
    """

    def test_interceptor_receives_request(self, mock_version_service):
        """Interceptor can observe the service and method being called."""
        vs = MockVersionService()

        services = {
            "/fits.api.v1.VersionService": VersionServiceWSGIApplication(vs),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))

        interceptor = RequestRecordingInterceptor()

        svc = VersionServiceClientSync(
            "http://test",
            http_client=pyqwest.SyncClient(transport=transport),
            interceptors=[interceptor],
        )

        resp = svc.get(
            request=version_pb2.VersionServiceGetRequest(),
        )

        assert resp.version.version == "1.0"
        assert len(interceptor.calls) == 1
        assert interceptor.calls[0] == ("fits.api.v1.VersionService", "Get")

    def test_auth_interceptor_injects_token(self):
        """Auth interceptor injects Authorization header that the server receives."""
        vs = MockVersionService()

        services = {
            "/fits.api.v1.VersionService": VersionServiceWSGIApplication(vs),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))

        svc = VersionServiceClientSync(
            "http://test",
            http_client=pyqwest.SyncClient(transport=transport),
            interceptors=[AuthInterceptor("my-secret-token")],
        )

        resp = svc.get(request=version_pb2.VersionServiceGetRequest())

        assert resp.version.version == "1.0"
        assert vs.received_auth == "Bearer my-secret-token"

    def test_interceptor_on_client_wrapper(self, test_client, mock_version_service):
        """Client wrapper passes headers through to the underlying service."""
        token = "wrapper-test-token"
        resp = test_client.apiv1().version().get(
            request=version_pb2.VersionServiceGetRequest(),
            headers={"authorization": f"Bearer {token}"},
        )
        assert resp.version.version == "1.0"
        assert mock_version_service.received_auth == f"Bearer {token}"
