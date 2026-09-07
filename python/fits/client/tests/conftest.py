from typing import Any

import pyqwest
import pytest
from connectrpc.request import RequestContext
from pyqwest.testing import WSGITransport

from fits.api.v1 import ip_pb2
from fits.api.v1 import version_pb2
from fits.api.v1.ip_connect import (
    IPServiceWSGIApplication,
)
from fits.api.v1.version_connect import (
    VersionServiceWSGIApplication,
)
from fits.client.client import Client


class MockVersionService:
    def __init__(self, version: str = "1.0"):
        self._version = version

    def get(
        self, request: version_pb2.VersionServiceGetRequest, ctx: RequestContext
    ) -> version_pb2.VersionServiceGetResponse:
        return version_pb2.VersionServiceGetResponse(
            version=version_pb2.Version(version=self._version)
        )


class MockIPService:
    def __init__(self):
        self.received_auth: str | None = None

    def get(
        self, request: ip_pb2.IPServiceGetRequest, ctx: RequestContext
    ) -> ip_pb2.IPServiceGetResponse:
        self.received_auth = ctx.request_headers().get("authorization")
        return ip_pb2.IPServiceGetResponse(
            ip=ip_pb2.IP(ip=request.ip, project=request.project)
        )

    def create(
        self, request: ip_pb2.IPServiceCreateRequest, ctx: RequestContext
    ) -> ip_pb2.IPServiceCreateResponse:
        return ip_pb2.IPServiceCreateResponse()

    def update(
        self, request: ip_pb2.IPServiceUpdateRequest, ctx: RequestContext
    ) -> ip_pb2.IPServiceUpdateResponse:
        return ip_pb2.IPServiceUpdateResponse()

    def list(
        self, request: ip_pb2.IPServiceListRequest, ctx: RequestContext
    ) -> ip_pb2.IPServiceListResponse:
        return ip_pb2.IPServiceListResponse()

    def delete(
        self, request: ip_pb2.IPServiceDeleteRequest, ctx: RequestContext
    ) -> ip_pb2.IPServiceDeleteResponse:
        return ip_pb2.IPServiceDeleteResponse()


def _build_combined_wsgi_app(services: dict[str, Any]):
    """Build a combined WSGI app that dispatches by path prefix."""

    def application(environ, start_response):
        path = environ.get("PATH_INFO", "/")
        for prefix, app in services.items():
            if path.startswith(prefix):
                return app(environ, start_response)
        from werkzeug.exceptions import NotFound

        return NotFound()(environ, start_response)

    return application


@pytest.fixture
def mock_version_service():
    return MockVersionService()


@pytest.fixture
def mock_ip_service():
    return MockIPService()


@pytest.fixture
def test_client(
    mock_version_service,
    mock_ip_service,
):
    """Create a Client with WSGI-backed transport for testing."""
    services = {
        "/fits.api.v1.VersionService": VersionServiceWSGIApplication(
            mock_version_service
        ),
        "/fits.api.v1.IPService": IPServiceWSGIApplication(mock_ip_service),
    }
    transport = WSGITransport(_build_combined_wsgi_app(services))
    c = Client(baseurl="http://test", timeout=10)
    c._client = pyqwest.SyncClient(transport=transport)
    return c
