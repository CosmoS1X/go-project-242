package main

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	appname := "pathsize"
	path := filepath.Join("..", "..", "testdata", "test.txt")
	directoryPath := filepath.Join("..", "..", "testdata")
	errMsg := "Error:"

	cases := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "prints size and path for existing file",
			args:       []string{appname, path},
			wantStdout: "6B\t" + path + "\n",
		},
		{
			name:       "prints size and path for existing directory",
			args:       []string{appname, "-raH", directoryPath},
			wantStdout: "293.9KB\t" + directoryPath + "\n",
		},
		{
			name:       "returns error for nonexistent path",
			args:       []string{appname, "unknown"},
			wantCode:   1,
			wantStderr: errMsg,
		},
		{
			name:       "returns error if too many arguments are provided",
			args:       []string{appname, "path1", "path2"},
			wantCode:   1,
			wantStderr: errMsg,
		},
		{
			name:       "returns error if no path argument is provided",
			args:       []string{appname},
			wantCode:   1,
			wantStderr: errMsg,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			code := run(c.args, &stdout, &stderr)

			require.Equal(t, c.wantCode, code, "expected exit code")
			if c.wantStdout == "" {
				require.Empty(t, stdout.String(), "expected empty stdout")
				require.Contains(t, stderr.String(), c.wantStderr, "expected error message in stderr")
			} else {
				require.Equal(t, c.wantStdout, stdout.String(), "expected stdout")
				require.Empty(t, stderr.String(), "expected empty stderr")
			}
		})
	}
}
