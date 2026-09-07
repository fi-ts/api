"""Tests mirroring go/client/client_test.go.

Test_Client -> test_client_apiv1_version_get
Test_ClientInterceptors -> test_client_auth_header_ip_get
"""

import pyqwest
import pytest
from pyqwest.testing import WSGITransport

from fits.api.v1 import ip_pb2
from fits.api.v1 import version_pb2
from fits.api.v1.ip_connect import (
    IPServiceClientSync,
    IPServiceWSGIApplication,
)
from fits.api.v1.version_connect import (
    VersionServiceClientSync,
    VersionServiceWSGIApplication,
)
from fits.client.client import Client

from .conftest import (
    MockIPService,
    MockVersionService,
    _build_combined_wsgi_app,
)


class TestClient:
    """Mirrors Test_Client from client_test.go.

    Tests client instantiation and basic RPC calls through the generated Client wrapper.
    """

    def test_client_creates_service_groups(self, test_client):
        """Client exposes the apiv1 service group."""
        assert test_client.apiv1() is not None

    def test_client_apiv1_version_get(self, test_client, mock_version_service):
        """Version.Get returns the expected version through the Client wrapper."""
        resp = test_client.apiv1().version().get(
            request=version_pb2.VersionServiceGetRequest()
        )
        assert resp is not None
        assert resp.version.version == "1.0"

    def test_client_apiv1_version_get_custom_version(self, mock_version_service):
        """Client can be configured with a custom version response."""
        mock_version_service._version = "2.5.3"
        vs = MockVersionService(version="2.5.3")
        ips = MockIPService()

        services = {
            "/fits.api.v1.VersionService": VersionServiceWSGIApplication(vs),
            "/fits.api.v1.IPService": IPServiceWSGIApplication(ips),
        }
        transport = WSGITransport(_build_combined_wsgi_app(services))
        c = Client(baseurl="http://test", timeout=10)
        c._client = pyqwest.SyncClient(transport=transport)

        resp = c.apiv1().version().get(
            request=version_pb2.VersionServiceGetRequest()
        )
        assert resp.version.version == "2.5.3"


class TestClientInterceptors:
    """Mirrors Test_ClientInterceptors from client_test.go.

    Tests that auth headers are correctly passed on unary calls.
    """

    def test_client_auth_header_ip_get(self, test_client, mock_ip_service):
        """IP Get call has Authorization header set."""
        token = "ip-service-token"
        resp = test_client.apiv1().ip().get(
            request=ip_pb2.IPServiceGetRequest(ip="10.0.0.1"),
            headers={"authorization": f"Bearer {token}"},
        )
        assert resp is not None
        assert resp.ip.ip == "10.0.0.1"
        assert mock_ip_service.received_auth == f"Bearer {token}"
