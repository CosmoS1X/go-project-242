package main

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	appname := "hexlet-path-size"

	t.Run("prints size and path for existing file", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		path := filepath.Join("..", "..", "testdata", "test.txt")
		args := []string{appname, path}

		code := run(args, &stdout, &stderr)

		require.Equal(t, 0, code, "expected exit code 0")
		require.Equal(t, "6B\t"+path+"\n", stdout.String(), "unexpected stdout")
		require.Empty(t, stderr.String(), "expected empty stderr")
	})

	t.Run("prints size and path for existing directory", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		path := filepath.Join("..", "..", "testdata")
		flags := "-raH"
		args := []string{appname, flags, path}

		code := run(args, &stdout, &stderr)

		require.Equal(t, 0, code, "expected exit code 0")
		require.Equal(t, "293.9KB\t"+path+"\n", stdout.String(), "unexpected stdout")
		require.Empty(t, stderr.String(), "expected empty stderr")
	})

	t.Run("returns error for nonexistent path", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		args := []string{appname, "unknown"}

		code := run(args, &stdout, &stderr)

		require.Equal(t, 1, code, "expected exit code 1")
		require.Empty(t, stdout.String(), "expected empty stdout")
		require.Contains(t, stderr.String(), "Error:", "expected error message in stderr")
	})

	t.Run("returns error if too many arguments are provided", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		args := []string{appname, "path1", "path2"}

		code := run(args, &stdout, &stderr)

		require.Equal(t, 1, code, "expected exit code 1")
		require.Empty(t, stdout.String(), "expected empty stdout")
		require.Contains(t, stderr.String(), "Error:", "expected error message in stderr")
	})

	t.Run("returns error if no path argument is provided", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		args := []string{appname}

		code := run(args, &stdout, &stderr)

		require.Equal(t, 1, code, "expected exit code 1")
		require.Empty(t, stdout.String(), "expected empty stdout")
		require.Contains(t, stderr.String(), "Error:", "expected error message in stderr")
	})
}
