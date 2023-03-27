package cmd

import (
	"docker-color-output/internal/layout"
)

const ValidTotalParts = 2

type Command interface {
	Format(layout.Row, layout.Column) string
}
