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
	"html/template"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/doomshrine/gocosi/cmd/bootstrap/internal/config"
)

func Write(fs embed.FS, config *config.Config, root string) error {
	if root == "" {
		root = config.TemplateRoot
	}

	children, err := fs.ReadDir(root)
	if err != nil {
		return err
	}

	for _, child := range children {
		if child.IsDir() {
			if err := Write(fs, config, path.Join(root, child.Name())); err != nil {
				return err
			}

			continue
		}

		if err := writeFile(fs, config, path.Join(root, child.Name())); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(fs embed.FS, config *config.Config, filepath string) error {
	dir, filename := path.Split(filepath)

	b, err := fs.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("unable to read '%s' file: %w", dir, err)
	}

	out := path.Join(config.OutputDir, removePathPrefix(dir, config.TemplateRoot))

	if err := os.MkdirAll(out, 0o755); err != nil {
		return fmt.Errorf("unable to create '%s' directory: %w", dir, err)
	}

	filepath = path.Join(out, replaceAll(filename))

	return execTemplate(filepath, string(b), config)
}

func execTemplate(filepath, tpl string, cfg *config.Config) error {
	t, err := template.New(filepath).Parse(tpl)
	if err != nil {
		return fmt.Errorf("unable to build '%s' template: %w", filepath, err)
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("unable to create '%s': %w", filepath, err)
	}

	err = t.Execute(f, cfg)
	if err != nil {
		return fmt.Errorf("unable to write '%s' template: %w", filepath, err)
	}

	return nil
}

// removePathPrefix removes the specified prefix and any leading forward slashes from the given string.
// If the 'prefix' is not found in the input string 's', the function returns the original 's' unchanged.
// The function is case-sensitive, requiring an exact match of the 'prefix' characters in the input string.
//
// Example:
//
//	inputPath := "images/pic123.jpg"
//	prefixToRemove := "images"
//	result := removePathPrefix(inputPath, prefixToRemove)
//	// result will be "pic123.jpg"
func removePathPrefix(s, prefix string) string {
	return strings.TrimLeft(strings.TrimPrefix(s, prefix), "/")
}

var replacementRules = map[*regexp.Regexp]string{
	regexp.MustCompile(`^HIDDEN\.`): ".",
	regexp.MustCompile(`\.tpl$`):    "",
}

func replaceAll(s string) string {
	for k, v := range replacementRules {
		s = k.ReplaceAllString(s, v)
	}

	return s
}
