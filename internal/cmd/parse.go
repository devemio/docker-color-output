package cmd

import (
	"errors"

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

	ps := &DockerPs{rules: cfg.DockerPs}
	if util.Intersect(columns, ps.Columns()) {
		return ps, nil
	}

	images := &DockerImages{rules: cfg.DockerImages}
	if util.Intersect(columns, images.Columns()) {
		return images, nil
	}

	composePs := &DockerComposePs{rules: cfg.DockerComposePs}
	if util.Intersect(columns, composePs.Columns()) {
		return composePs, nil
	}

	stats := &DockerStats{rules: cfg.DockerStats}
	if util.Intersect(columns, stats.Columns()) {
		return stats, nil
	}

	return nil, ErrInvalidFirstLine
}
