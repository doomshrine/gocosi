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

package config

import (
	"bytes"
	"net/url"
	"runtime"
	"testing"
	"text/template"

	"github.com/doomshrine/must"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testModPath = "example.com/module"

// TestURLParse stays here so I remember if URL Parsing works as I think it works 😳.
func TestURLParse(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		u := must.Do(url.Parse(testModPath))
		assert.Equal(t, testModPath, u.String())

		tplURL, err := template.New("test.URL").Parse("{{ .URL }}")
		require.NoError(t, err)

		buf := new(bytes.Buffer)

		err = tplURL.Execute(buf, struct{ URL *url.URL }{URL: u})
		assert.NoError(t, err)
		assert.Equal(t, testModPath, buf.String())

		buf.Reset()

		// This will actually return "example.com/module" instead of "module".
		// This is caused by combination of the following lines:
		// - (go1.20.7) net/url:519 , extracting schema
		// - (go1.20.7) net/url:553 , setting host only if URL has schema.
		//   'viaRequest' is always false, due to the (go1.20.7) net/url:469
		// - (go1.20.7) net/url:573 , the 'rest' is set as URL.Path
		tplURLPath, err := template.New("test.URL.Path").Parse("{{ .URL.Path }}")
		require.NoError(t, err)

		err = tplURLPath.Execute(buf, struct{ URL *url.URL }{URL: u})
		assert.NoError(t, err)
		assert.Equal(t, testModPath, buf.String())
	})
}

func TestNew(t *testing.T) {
	t.Parallel()

	defaultValidator := func(t *testing.T, c *Config) {
		require.NotNil(t, c)
		assert.Equal(t, testModPath, c.ModPath.String())
		assert.Equal(t, trimVersion(runtime.Version()), c.GoVersion)
	}

	emptyValidator := func(t *testing.T, c *Config) {}

	for _, tc := range []struct {
		name             string
		modPath          string
		errorRequirement func(require.TestingT, error, ...interface{})
		dataValidator    func(*testing.T, *Config)
	}{
		{
			name:             "happy path",
			modPath:          testModPath,
			errorRequirement: require.NoError,
		},
		{
			name:             "malformed modpath",
			modPath:          ":mod/path",
			errorRequirement: require.Error,
			dataValidator:    emptyValidator,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.NotNil(t, tc.errorRequirement)
			if tc.dataValidator == nil {
				tc.dataValidator = defaultValidator
			}

			c, err := New(tc.modPath)
			tc.errorRequirement(t, err)
			tc.dataValidator(t, c)
		})
	}
}

func TestTrimVersion(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name      string
		assertion func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
		caller    func()
	}{
		{
			name:      "current version",
			assertion: assert.NotPanics,
			caller: func() {
				_ = trimVersion(runtime.Version())
			},
		},
		{
			name:      "static version",
			assertion: assert.NotPanics,
			caller: func() {
				_ = trimVersion("go1.18.0")
			},
		},
		{
			name:      "invalid version",
			assertion: assert.Panics,
			caller: func() {
				_ = trimVersion("go1.18")
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.assertion(t, tc.caller)
		})
	}
}
