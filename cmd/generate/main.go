// Copyright © 2023 Doomshrine and gocosi authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type Config struct {
	ModPath   string
	GoVersion string
}

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	if len(os.Args) < 2 {
		return errors.New("no name specified")
	}

	cfg := Config{
		ModPath:   os.Args[1],
		GoVersion: "1.20",
	}

	err := os.MkdirAll("./servers/provisioner", 0o755)
	if err != nil {
		return fmt.Errorf("unable to create './servers/provisioner' directory: %w", err)
	}

	err = os.MkdirAll("./servers/identity", 0o755)
	if err != nil {
		return fmt.Errorf("unable to create './servers/identity' directory: %w", err)
	}

	for _, tpl := range []struct {
		filepath string
		template string
	}{
		{
			filepath: "go.mod",
			template: goMod,
		},
		{
			filepath: "main.go",
			template: mainGo,
		},
		{
			filepath: "./servers/provisioner/provisioner.go",
			template: provisionerGo,
		},
		{
			filepath: "./servers/identity/identity.go",
			template: identityGo,
		},
	} {
		err := execTemplate(tpl.filepath, tpl.template, cfg)
		if err != nil {
			return fmt.Errorf("unable to execute teplate '%s': %w", tpl.filepath, err)
		}
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed running 'go mod tidy': %w", err)
	}

	return nil
}

func execTemplate(filepath, tpl string, cfg Config) error {
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
