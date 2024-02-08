//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/devemio/docker-color-output/internal/app"
	"github.com/devemio/docker-color-output/internal/config"
	"github.com/devemio/docker-color-output/internal/stdin"
	"github.com/devemio/docker-color-output/internal/stdout"
	"github.com/devemio/docker-color-output/pkg/color"
)

func main() {
	if err := run(); err != nil {
		app.Usage(err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Get()
	if err != nil {
		return fmt.Errorf("cfg: %w", err)
	}

	color.SetPalette(color.Palette(cfg.Colors))

	return stdin.Get(func(rows []string) error { //nolint:wrapcheck
		rows, err = app.Run(rows)
		if err != nil {
			return fmt.Errorf("app: %w", err)
		}

		for _, row := range rows {
			stdout.Println(row)
		}

		return nil
	})
}
