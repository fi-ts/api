"""Tests mirroring go/client/ping_test.go.

Tests the ComponentService.Ping RPC which is used by the Go client's Ping() method.
The Python client doesn't have a built-in Ping() method, but we test the underlying
ComponentService call that would be used to implement one.
"""

import threading
import time
from datetime import timedelta

import pytest
from google.protobuf import duration_pb2, timestamp_pb2

from fits.api.v2 import component_pb2 as api_component_pb2
from fits.api.v2 import version_pb2
from fits.infra.v2 import component_pb2 as infra_component_pb2


class TestPing:
    """Mirrors Test_Ping from ping_test.go.

    Verifies that ComponentServicePingRequest is correctly constructed and sent,
    and that the mock service receives the expected fields.
    """

    def test_ping_request_fields(self, test_client, mock_component_service):
        """Ping request contains correct component type, identifier, started_at, interval, version."""
        started_at = timestamp_pb2.Timestamp(seconds=1000)
        interval = duration_pb2.Duration(seconds=5)
        version = version_pb2.Version(version="1.0.0", git_sha1="abc123")

        resp = test_client.infrav2().component().ping(
            request=infra_component_pb2.ComponentServicePingRequest(
                type=api_component_pb2.COMPONENT_TYPE_METAL_CORE,
                identifier="server01",
                started_at=started_at,
                interval=interval,
                version=version,
            ),
            headers={"authorization": "Bearer test-token"},
        )

        assert resp is not None
        assert len(mock_component_service.pings) == 1

        ping = mock_component_service.pings[0]
        assert ping.type == api_component_pb2.COMPONENT_TYPE_METAL_CORE
        assert ping.identifier == "server01"
        assert ping.started_at.seconds == 1000
        assert ping.interval.seconds == 5
        assert ping.version.version == "1.0.0"
        assert ping.version.git_sha1 == "abc123"

    def test_ping_multiple_requests(self, test_client, mock_component_service):
        """Multiple ping calls are recorded correctly."""
        started_at = timestamp_pb2.Timestamp(seconds=2000)
        interval = duration_pb2.Duration(seconds=10)

        for _ in range(3):
            test_client.infrav2().component().ping(
                request=infra_component_pb2.ComponentServicePingRequest(
                    type=api_component_pb2.COMPONENT_TYPE_METAL_CORE,
                    identifier="server02",
                    started_at=started_at,
                    interval=interval,
                ),
                headers={"authorization": "Bearer test-token"},
            )

        assert len(mock_component_service.pings) == 3
        for ping in mock_component_service.pings:
            assert ping.identifier == "server02"
            assert ping.type == api_component_pb2.COMPONENT_TYPE_METAL_CORE
            assert ping.started_at.seconds == 2000
            assert ping.interval.seconds == 10

    def test_ping_background_simulation(self, test_client, mock_component_service):
        """Simulate background pinging from a thread, similar to Go's ticker goroutine."""
        started_at = timestamp_pb2.Timestamp(seconds=3000)
        interval = duration_pb2.Duration(seconds=1)
        ping_count = 5

        def do_pings():
            for _ in range(ping_count):
                test_client.infrav2().component().ping(
                    request=infra_component_pb2.ComponentServicePingRequest(
                        type=api_component_pb2.COMPONENT_TYPE_METAL_CORE,
                        identifier="background-worker",
                        started_at=started_at,
                        interval=interval,
                    ),
                    headers={"authorization": "Bearer test-token"},
                )
                time.sleep(0.05)

        thread = threading.Thread(target=do_pings)
        thread.start()
        thread.join()

        assert len(mock_component_service.pings) == ping_count
        for ping in mock_component_service.pings:
            assert ping.identifier == "background-worker"
