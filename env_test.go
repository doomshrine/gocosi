package gocosi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name          string
		envKey        string
		defaultValue  string
		expectedValue string
		options       func(key, value string) error
	}{
		{
			name:          "set variable",
			envKey:        "TEST_SET_VAR",
			defaultValue:  "default",
			expectedValue: "set",
			options:       os.Setenv,
		},
		{
			name:          "unset variable",
			envKey:        "TEST_UNSET_VAR",
			defaultValue:  "default",
			expectedValue: "default",
			options:       nil,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.options != nil {
				err := tc.options(tc.envKey, tc.expectedValue)
				require.NoError(t, err)
			}

			actual := getEnv(tc.envKey, tc.defaultValue)
			assert.Equal(t, tc.expectedValue, actual)
		})
	}
}
