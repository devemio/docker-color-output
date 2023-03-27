package cmd

import (
	"errors"

	"docker-color-output/internal/layout"
	"docker-color-output/internal/util"
)

var ErrInvalidFirstLine = errors.New("invalid first line")

func Parse(header layout.Header) (Command, error) { //nolint:ireturn
	columns := make([]string, len(header))
	for i, col := range header {
		columns[i] = string(col.Name)
	}

	ps := &DockerPs{}
	if util.Intersect(columns, ps.Columns()) {
		return ps, nil
	}

	images := &DockerImages{}
	if util.Intersect(columns, images.Columns()) {
		return images, nil
	}

	composePs := &DockerComposePs{}
	if util.Intersect(columns, composePs.Columns()) {
		return composePs, nil
	}

	return nil, ErrInvalidFirstLine
}
