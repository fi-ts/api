"""Tests mirroring go/client/token-persistent-funcs_test.go.

Tests filesystem-based token persistence, which is the Python equivalent of
Go's NewFilesystemTokenPersister.
"""

import os
import tempfile
from pathlib import Path

import pytest


class FilesystemTokenPersister:
    """Persists a token to a file on disk.

    Mirrors the filesystem struct and persist method from
    go/client/token-persistent-funcs.go.
    """

    def __init__(self, token_path: str):
        self._token_path = token_path

    def persist(self, token: str) -> None:
        Path(self._token_path).write_text(token)


def new_filesystem_token_persister(token_file: str) -> FilesystemTokenPersister:
    """Create a new filesystem token persister.

    Mirrors NewFilesystemTokenPersister from go/client/token-persistent-funcs.go.
    """
    path = Path(token_file)
    if not path.exists():
        raise FileNotFoundError(f"token file {token_file} does not exist")
    if os.access(token_file, os.W_OK) is False:
        raise ValueError(f"tokenfile {token_file} is not writable")
    return FilesystemTokenPersister(token_file)


class TestFilesystemTokenPersister:
    """Mirrors TestNewFilesystemTokenPersiter from token-persistent-funcs_test.go."""

    def test_persist_and_read_back(self, tmp_path):
        """Write a token to a file and read it back."""
        token_path = str(tmp_path / "token")
        Path(token_path).touch(mode=0o600)

        persister = new_filesystem_token_persister(token_path)
        token = "a-token"

        persister.persist(token)

        content = Path(token_path).read_text()
        assert content == token

    def test_persister_requires_existing_file(self, tmp_path):
        """Persister raises FileNotFoundError for non-existent file."""
        token_path = str(tmp_path / "nonexistent")
        with pytest.raises(FileNotFoundError):
            new_filesystem_token_persister(token_path)

    def test_persister_requires_writable_file(self, tmp_path):
        """Persister raises ValueError for read-only file."""
        token_path = str(tmp_path / "readonly")
        Path(token_path).touch(mode=0o444)

        with pytest.raises(ValueError, match="not writable"):
            new_filesystem_token_persister(token_path)

    def test_overwrite_existing_token(self, tmp_path):
        """Persisting a new token overwrites the previous one."""
        token_path = str(tmp_path / "token")
        Path(token_path).write_text("old-token")

        persister = new_filesystem_token_persister(token_path)
        persister.persist("new-token")

        content = Path(token_path).read_text()
        assert content == "new-token"

    def test_persist_empty_token(self, tmp_path):
        """Persisting an empty token writes an empty file."""
        token_path = str(tmp_path / "token")
        Path(token_path).touch(mode=0o600)

        persister = new_filesystem_token_persister(token_path)
        persister.persist("")

        content = Path(token_path).read_text()
        assert content == ""
