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

package gocosi

import (
	"context"
	"net/url"
	"testing"

	"github.com/doomshrine/gocosi/testutils"
	"github.com/doomshrine/must"
	"github.com/doomshrine/testcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncpointListener(t *testing.T) {
	t.Parallel()

	ctx, cancel := testcontext.FromT(context.Background(), t)
	defer cancel()

	for _, tc := range []struct {
		name           string
		endpoint       *Endpoint
		errorAssertion func(assert.TestingT, error, ...interface{}) bool
		valueAssertion func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{
			name:           "valid UNIX socket",
			errorAssertion: assert.NoError,
			valueAssertion: assert.NotNil,
			endpoint: &Endpoint{
				address: testutils.MustMkUnixTemp("cosi.sock"),
			},
		},
		{
			name:           "valid TCP socket",
			errorAssertion: assert.NoError,
			valueAssertion: assert.NotNil,
			endpoint: &Endpoint{
				address: must.Do(url.Parse("tcp://:0")),
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, tc.endpoint, "endpoint is required")
			defer tc.endpoint.Close() // no matter what, I need it to be called

			lis, err := tc.endpoint.Listener(ctx)
			tc.errorAssertion(t, err)
			tc.valueAssertion(t, lis)

			lis, err = tc.endpoint.Listener(ctx)
			tc.errorAssertion(t, err)
			tc.valueAssertion(t, lis)
		})
	}
}
