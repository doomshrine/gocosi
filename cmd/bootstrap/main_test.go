package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/doomshrine/testcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const TestModPath = "main.test/module"

func TestRealMain(t *testing.T) {
	t.Parallel()

	ctx, cancel := testcontext.FromT(context.Background(), t)
	defer cancel()

	dir, err := os.MkdirTemp("", "*")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	err = realMain(TestModPath, dir)
	require.NoError(t, err)

	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	cmd := exec.CommandContext(ctx, "go", "build")
	cmd.Dir = dir
	cmd.Stderr = bufErr
	cmd.Stdout = bufOut

	err = cmd.Run()
	assert.NoError(t, err, "stdout: >>>%s<<<, stderr: >>>%s<<<", bufOut.String(), bufErr.String())
}
