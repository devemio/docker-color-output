//go:build !test

package main

import (
	"os"

	"github.com/devemio/docker-color-output/internal/app"
	"github.com/devemio/docker-color-output/internal/stdin"
	"github.com/devemio/docker-color-output/internal/stdout"
)

func main() {
	if err := run(); err != nil {
		app.Usage(err)
		os.Exit(1)
	}
}

func run() error {
	in, err := stdin.Get()
	if err != nil {
		return err //nolint:wrapcheck
	}

	rows, err := app.Run(in)
	if err != nil {
		return err //nolint:wrapcheck
	}

	for _, row := range rows {
		stdout.Println(row)
	}

	return nil
}
