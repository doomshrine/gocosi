// Copyright © 2023 gocosi authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package template

import (
	"embed"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/doomshrine/gocosi/cmd/bootstrap/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed test
	testData embed.FS

	testRoot = "test"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "*")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	testConfig, err := config.New(testRoot,
		config.WithTemplateRoot(testRoot),
		config.WithOutputDir(dir),
	)
	require.NoError(t, err)

	err = Write(testData, testConfig, "")
	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	t.Parallel()
}

func TestExecTemplate(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		template string
		config   *config.Config
	}{
		{
			name:     "happy path",
			template: "{{ .GoVersion }}",
			config:   &config.Config{GoVersion: "1.18"},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, tc.config)

			dir, err := os.MkdirTemp("", "*")
			require.NoError(t, err)

			defer os.RemoveAll(dir)

			out := path.Join(dir, "test.out")

			err = execTemplate(out, tc.template, tc.config)
			assert.NoError(t, err)
			assert.FileExists(t, out)
		})
	}
}

func TestRemovePathPerfix(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		s        string
		prefix   string
		expected string
	}{
		{
			s:        "test/path",
			prefix:   "test",
			expected: "path",
		},
	} {
		tc := tc

		t.Run(fmt.Sprintf("s:'%s' prefix:'%s'", tc.s, tc.prefix), func(t *testing.T) {
			t.Parallel()

			out := removePathPrefix(tc.s, tc.prefix)
			assert.Equal(t, tc.expected, out)
		})
	}
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()

	for s, expected := range map[string]string{
		"a":            "a",
		"a.tpl":        "a",
		"HIDDEN.a":     ".a",
		"HIDDEN.a.tpl": ".a",
	} {
		s, expected := s, expected

		t.Run(s, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, expected, replaceAll(s))
		})
	}
}
