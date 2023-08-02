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

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"
)

type Config struct {
	ModPath   string
	GoVersion string
}

var (
	modPath   string
	directory string
)

func main() {
	flag.StringVar(&modPath, "module", "example.com/cosi-osp", "Provide name for your new module.")
	flag.StringVar(&directory, "dir", "cosi-osp", "Location, where the module will be created.")
	flag.Parse()

	if err := realMain(modPath, directory); err != nil {
		log.Fatal(err)
	}
}

func realMain(modPath, location string) error {
	if modPath == "" || location == "" {
		return errors.New("invalid argument")
	}

	cfg := Config{
		ModPath:   modPath,
		GoVersion: "1.20",
	}

	for _, dir := range []string{
		location,
		path.Join(location, "servers/provisioner"),
		path.Join(location, "servers/identity"),
	} {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			return fmt.Errorf("unable to create '%s' directory: %w", dir, err)
		}
	}

	for _, tpl := range []struct {
		filepath string
		template string
	}{
		{
			filepath: path.Join(location, "go.mod"),
			template: goMod,
		},
		{
			filepath: path.Join(location, "main.go"),
			template: mainGo,
		},
		{
			filepath: path.Join(location, "./servers/provisioner/provisioner.go"),
			template: provisionerGo,
		},
		{
			filepath: path.Join(location, "./servers/identity/identity.go"),
			template: identityGo,
		},
	} {
		err := execTemplate(tpl.filepath, tpl.template, cfg)
		if err != nil {
			return fmt.Errorf("unable to execute teplate '%s': %w", tpl.filepath, err)
		}
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = location
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
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
