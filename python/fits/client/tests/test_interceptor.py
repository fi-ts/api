"""Tests mirroring go/client/test_interceptor_test.go.

Tests the interceptor pattern using connectrpc's MetadataInterceptorSync,
which is the Python equivalent of Go's TestClientInterceptor.
"""

import pyqwest
import pytest
from connectrpc.interceptor import MetadataInterceptorSync
from connectrpc.request import RequestContext
from pyqwest.testing import WSGITransport

from fits.api.v2 import ip_pb2
from fits.api.v2.ip_connect import (
    IPServiceClientSync,
    IPServiceWSGIApplication,
)
from fits.client.client import Client

from .conftest import MockIPService, _build_combined_wsgi_app


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

    def test_interceptor_receives_request(self, mock_ip_service):
        """Interceptor can observe the service and method being called."""
        ips = MockIPService()
        from fits.infra.v2.bmc_connect import BMCServiceWSGIApplication
        from fits.infra.v2.component_connect import ComponentServiceWSGIApplication
        from fits.api.v2.version_connect import VersionServiceWSGIApplication
        from fits.infra.v2 import bmc_pb2
        from fits.infra.v2 import component_pb2 as infra_component_pb2

        class NoopBMC:
            def update_b_m_c_info(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import UpdateBMCInfoResponse
                return UpdateBMCInfoResponse()
            def wait_for_b_m_c_command(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import WaitForBMCCommandResponse
                yield WaitForBMCCommandResponse()
            def b_m_c_command_done(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import BMCCommandDoneResponse
                return BMCCommandDoneResponse()

        class NoopComponent:
            def ping(self, request, ctx):
                from fits.infra.v2.component_pb2 import ComponentServicePingResponse
                return ComponentServicePingResponse()

        class NoopVersion:
            def get(self, request, ctx):
                from fits.api.v2.version_pb2 import VersionServiceGetResponse, Version
                return VersionServiceGetResponse(version=Version(version="0.0"))

        services = {
            "/fits.api.v2.VersionService": VersionServiceWSGIApplication(NoopVersion()),
            "/fits.api.v2.IPService": IPServiceWSGIApplication(ips),
            "/fits.infra.v2.BMCService": BMCServiceWSGIApplication(NoopBMC()),
            "/fits.infra.v2.ComponentService": ComponentServiceWSGIApplication(NoopComponent()),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))

        interceptor = RequestRecordingInterceptor()

        svc = IPServiceClientSync(
            "http://test",
            http_client=pyqwest.SyncClient(transport=transport),
            interceptors=[interceptor],
        )

        resp = svc.get(
            request=ip_pb2.IPServiceGetRequest(ip="1.2.3.4"),
        )

        assert resp.ip.ip == "1.2.3.4"
        assert len(interceptor.calls) == 1
        assert interceptor.calls[0] == ("fits.api.v2.IPService", "Get")

    def test_auth_interceptor_injects_token(self):
        """Auth interceptor injects Authorization header that the server receives."""
        ips = MockIPService()
        from fits.infra.v2.bmc_connect import BMCServiceWSGIApplication
        from fits.infra.v2.component_connect import ComponentServiceWSGIApplication
        from fits.api.v2.version_connect import VersionServiceWSGIApplication

        class NoopBMC:
            def update_b_m_c_info(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import UpdateBMCInfoResponse
                return UpdateBMCInfoResponse()
            def wait_for_b_m_c_command(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import WaitForBMCCommandResponse
                yield WaitForBMCCommandResponse()
            def b_m_c_command_done(self, request, ctx):
                from fits.infra.v2.bmc_pb2 import BMCCommandDoneResponse
                return BMCCommandDoneResponse()

        class NoopComponent:
            def ping(self, request, ctx):
                from fits.infra.v2.component_pb2 import ComponentServicePingResponse
                return ComponentServicePingResponse()

        class NoopVersion:
            def get(self, request, ctx):
                from fits.api.v2.version_pb2 import VersionServiceGetResponse, Version
                return VersionServiceGetResponse(version=Version(version="0.0"))

        services = {
            "/fits.api.v2.VersionService": VersionServiceWSGIApplication(NoopVersion()),
            "/fits.api.v2.IPService": IPServiceWSGIApplication(ips),
            "/fits.infra.v2.BMCService": BMCServiceWSGIApplication(NoopBMC()),
            "/fits.infra.v2.ComponentService": ComponentServiceWSGIApplication(NoopComponent()),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))

        svc = IPServiceClientSync(
            "http://test",
            http_client=pyqwest.SyncClient(transport=transport),
            interceptors=[AuthInterceptor("my-secret-token")],
        )

        resp = svc.get(request=ip_pb2.IPServiceGetRequest(ip="1.2.3.4"))

        assert resp.ip.ip == "1.2.3.4"
        assert ips.received_auth == "Bearer my-secret-token"

    def test_interceptor_on_client_wrapper(self, test_client, mock_ip_service):
        """Client wrapper passes headers through to the underlying service."""
        token = "wrapper-test-token"
        resp = test_client.apiv2().ip().get(
            request=ip_pb2.IPServiceGetRequest(ip="5.6.7.8"),
            headers={"authorization": f"Bearer {token}"},
        )
        assert resp.ip.ip == "5.6.7.8"
        assert mock_ip_service.received_auth == f"Bearer {token}"
