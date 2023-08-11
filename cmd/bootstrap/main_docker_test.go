//go:build integration

package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/doomshrine/gocosi/cmd/bootstrap/internal/config"
	"github.com/doomshrine/testcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestRealMainWithDocker(t *testing.T) {
	ctx, cancel := testcontext.FromT(context.Background(), t)
	defer cancel()

	dir, err := os.MkdirTemp("", "*")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	ospDir := path.Join(dir, "test-osp")

	err = realMain(TestModPath, ospDir, config.DefaultImage, config.DefaultRootlessImage, false)
	require.NoError(t, err)
	require.FileExists(t, path.Join(ospDir, "go.mod"))

	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	cmd := exec.CommandContext(ctx, "docker", "build", "--tag=gocosi:test", "--no-cache", ".")
	cmd.Dir = ospDir
	cmd.Stderr = bufErr
	cmd.Stdout = bufOut

	err = cmd.Run()
	assert.NoError(t, err, "stdout: >>>%s<<<, stderr: >>>%s<<<", bufOut.String(), bufErr.String())

	bufOut.Reset()
	bufErr.Reset()

	cmd = exec.CommandContext(ctx, "docker", "image", "rm", "gocosi:test")
	cmd.Dir = ospDir
	cmd.Stderr = bufErr
	cmd.Stdout = bufOut

	err = cmd.Run()
	assert.NoError(t, err, "stdout: >>>%s<<<, stderr: >>>%s<<<", bufOut.String(), bufErr.String())
}

//nolint:paralleltest
func TestRealMainWithDockerRootless(t *testing.T) {
	ctx, cancel := testcontext.FromT(context.Background(), t)
	defer cancel()

	dir, err := os.MkdirTemp("", "*")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	ospDir := path.Join(dir, "test-osp")

	err = realMain(TestModPath, ospDir, config.DefaultImage, config.DefaultRootlessImage, false)
	require.NoError(t, err)
	require.FileExists(t, path.Join(ospDir, "go.mod"))

	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	cmd := exec.CommandContext(ctx, "docker", "build", "--tag=gocosi:test", "--no-cache", ".")
	cmd.Dir = ospDir
	cmd.Stderr = bufErr
	cmd.Stdout = bufOut

	err = cmd.Run()
	assert.NoError(t, err, "stdout: >>>%s<<<, stderr: >>>%s<<<", bufOut.String(), bufErr.String())

	bufOut.Reset()
	bufErr.Reset()

	cmd = exec.CommandContext(ctx, "docker", "image", "rm", "gocosi:test")
	cmd.Dir = ospDir
	cmd.Stderr = bufErr
	cmd.Stdout = bufOut

	err = cmd.Run()
	assert.NoError(t, err, "stdout: >>>%s<<<, stderr: >>>%s<<<", bufOut.String(), bufErr.String())
}
