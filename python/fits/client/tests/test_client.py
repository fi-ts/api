"""Tests mirroring go/client/client_test.go.

Test_Client -> test_client_apiv2_version_get
Test_ClientInterceptors -> test_client_auth_header_unary, test_client_auth_header_streaming
"""

import pyqwest
import pytest
from pyqwest.testing import WSGITransport

from fits.api.v2 import ip_pb2
from fits.api.v2 import version_pb2
from fits.api.v2.ip_connect import (
    IPServiceClientSync,
    IPServiceWSGIApplication,
)
from fits.api.v2.version_connect import (
    VersionServiceClientSync,
    VersionServiceWSGIApplication,
)
from fits.client.client import Client
from fits.infra.v2 import bmc_pb2
from fits.infra.v2.bmc_connect import (
    BMCServiceClientSync,
    BMCServiceWSGIApplication,
)
from fits.infra.v2 import component_pb2 as infra_component_pb2
from fits.infra.v2.component_connect import (
    ComponentServiceClientSync,
    ComponentServiceWSGIApplication,
)

from .conftest import (
    MockBMCService,
    MockComponentService,
    MockIPService,
    MockVersionService,
    _build_combined_wsgi_app,
)


class TestClient:
    """Mirrors Test_Client from client_test.go.

    Tests client instantiation and basic RPC calls through the generated Client wrapper.
    """

    def test_client_creates_service_groups(self, test_client):
        """Client exposes adminv2, apiv2, infrav2 service groups."""
        assert test_client.apiv2() is not None
        assert test_client.infrav2() is not None

    def test_client_apiv2_version_get(self, test_client, mock_version_service):
        """Version.Get returns the expected version through the Client wrapper."""
        resp = test_client.apiv2().version().get(
            request=version_pb2.VersionServiceGetRequest()
        )
        assert resp is not None
        assert resp.version.version == "1.0"

    def test_client_apiv2_version_get_custom_version(self, mock_version_service):
        """Client can be configured with a custom version response."""
        mock_version_service._version = "2.5.3"
        vs = MockVersionService(version="2.5.3")
        ips = MockIPService()
        bmc = MockBMCService()
        cs = MockComponentService()

        services = {
            "/fits.api.v2.VersionService": VersionServiceWSGIApplication(vs),
            "/fits.api.v2.IPService": IPServiceWSGIApplication(ips),
            "/fits.infra.v2.BMCService": BMCServiceWSGIApplication(bmc),
            "/fits.infra.v2.ComponentService": ComponentServiceWSGIApplication(cs),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))
        c = Client(baseurl="http://test", timeout=10)
        c._client = pyqwest.SyncClient(transport=transport)

        resp = c.apiv2().version().get(
            request=version_pb2.VersionServiceGetRequest()
        )
        assert resp.version.version == "2.5.3"


class TestClientInterceptors:
    """Mirrors Test_ClientInterceptors from client_test.go.

    Tests that auth headers are correctly passed for both unary and streaming calls.
    """

    def test_client_auth_header_unary(self, test_client, mock_bmc_service):
        """Unary call has Authorization header set."""
        token = "test-bearer-token"
        resp = test_client.infrav2().bmc().update_b_m_c_info(
            request=bmc_pb2.UpdateBMCInfoRequest(),
            headers={"authorization": f"Bearer {token}"},
        )
        assert resp is not None
        assert mock_bmc_service.received_auth == f"Bearer {token}"

    def test_client_auth_header_streaming(self, test_client, mock_bmc_service):
        """Streaming call has Authorization header set."""
        token = "test-stream-token"
        stream = test_client.infrav2().bmc().wait_for_b_m_c_command(
            request=bmc_pb2.WaitForBMCCommandRequest(),
            headers={"authorization": f"Bearer {token}"},
        )
        assert stream is not None
        for cmd in stream:
            assert cmd.uuid == "test-uuid"
        assert mock_bmc_service.received_auth == f"Bearer {token}"

    def test_client_auth_header_ip_get(self, test_client, mock_ip_service):
        """IP Get call has Authorization header set."""
        token = "ip-service-token"
        resp = test_client.apiv2().ip().get(
            request=ip_pb2.IPServiceGetRequest(ip="10.0.0.1"),
            headers={"authorization": f"Bearer {token}"},
        )
        assert resp is not None
        assert resp.ip.ip == "10.0.0.1"
        assert mock_ip_service.received_auth == f"Bearer {token}"
