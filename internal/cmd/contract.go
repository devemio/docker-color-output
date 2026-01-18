package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
)

type Command interface {
	Format(rows layout.Row, col layout.Column) string
}
