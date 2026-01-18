package cmd

import (
	"errors"
	"sort"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
	"github.com/devemio/docker-color-output/internal/util"
)

var ErrInvalidFirstLine = errors.New("invalid first line")

func Parse(header layout.Header, cfg rules.Config) (Command, error) { //nolint:ireturn
	columns := make([]string, len(header))
	for i, col := range header {
		columns[i] = string(col.Name)
	}

	candidates := make([]string, 0, len(cfg.Commands))
	for name := range cfg.Commands {
		candidates = append(candidates, name)
	}
	sort.Strings(candidates)

	for _, name := range candidates {
		command := cfg.Commands[name]
		if util.Intersect(columns, command.Columns) {
			return &DynamicCommand{rules: command}, nil
		}
	}

	return nil, ErrInvalidFirstLine
}
