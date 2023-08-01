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