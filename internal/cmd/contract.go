package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
)

const ValidTotalParts = 2

type Command interface {
	Format(rows layout.Row, col layout.Column) string
}
