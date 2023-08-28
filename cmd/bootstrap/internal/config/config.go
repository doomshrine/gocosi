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
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type Config struct {
	ModPath      *url.URL
	TemplateRoot string
	OutputDir    string
	GoVersion    string
	Year         int
	Comment      string
	Rootless     bool

	COSISpecification *COSISpecification
	Docker            *Docker
}

// New returns new Config struct, given that the modPath is a valid URL.
func New(modPath string, opts ...Option) (*Config, error) {
	// This should probably be more resilient to invalid module paths.
	modURL, err := url.Parse(modPath)
	if err != nil {
		return nil, fmt.Errorf("invalid module path format: %w", err)
	}

	cfg := &Config{
		ModPath:      modURL,
		GoVersion:    trimVersion(runtime.Version()),
		Year:         time.Now().Year(),
		Comment:      "//",
		TemplateRoot: "template",
		Rootless:     false,

		Docker:            newDocker(),
		COSISpecification: newSpecification(),
	}

	var combinedErrors error

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			combinedErrors = errors.Join(combinedErrors, err)
		}
	}

	return cfg, combinedErrors
}

func (c *Config) WithComment(comment string) *Config {
	c.Comment = comment
	return c
}

// trimVersion is supposed to remove prefix and suffix from 'goX.Y.Z', so that the
// obtained value is of the 'X.Y' format.
func trimVersion(input string) string {
	// Define a regular expression to match the "goX.Y.Z" format
	re := regexp.MustCompile(`^go(\d+\.\d+\.\d+)$`)

	// Check if the input matches the expected format
	match := re.FindStringSubmatch(input)
	if len(match) != 2 {
		panic("Invalid input format")
	}

	// Extract and return the trimmed version
	versionParts := strings.Split(match[1], ".")

	return versionParts[0] + "." + versionParts[1]
}
