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

	"github.com/doomshrine/gocosi/cmd/bootstrap/internal/config"
	"github.com/doomshrine/gocosi/cmd/bootstrap/internal/template"
)

var (
	modPath       string
	directory     string
	image         string
	rootlessImage string
	rootless      bool
)

func main() {
	flag.StringVar(&modPath, "module", "example.com/cosi-osp", "Override name for your new module.")
	flag.StringVar(&directory, "dir", "cosi-osp", "Location/Path, where the module will be created.")
	flag.StringVar(&image, "image", config.DefaultImage, "Override the default base Docker image.")
	flag.StringVar(&rootlessImage, "rootless-image", config.DefaultRootlessImage, "Override the default base Docker image for rootless container.")
	flag.BoolVar(&rootless, "rootless", false, "Generate the Dockerfile for rootless container.")
	flag.Parse()

	if err := realMain(modPath,
		directory,
		image,
		rootlessImage,
		rootless); err != nil {
		log.Fatal(err)
	}
}

func realMain(modPath, location, image, rootlessImage string, rootless bool) error {
	if modPath == "" || location == "" || image == "" || rootlessImage == "" {
		return errors.New("invalid argument")
	}

	if _, err := os.Stat(location); err == nil {
		return errors.New("location already exists")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("unexpected error: %w", err)
	}

	cfg, err := config.New(modPath,
		config.WithOutputDir(location),
		config.WithDockerImage(image),
		config.WithDockerRootlessImage(rootlessImage),
		config.WithRootless(rootless),
	)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	err = template.Write(templateFS, cfg, "")
	if err != nil {
		return fmt.Errorf("error writing template: %w", err)
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = location
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed running 'go mod tidy': %w", err)
	}

	return nil
}
