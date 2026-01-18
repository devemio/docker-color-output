package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
)

const (
	DockerImagesRepository = "REPOSITORY"
	DockerImagesTag        = "TAG"
	DockerImagesDigest     = "DIGEST"
	DockerImagesImageID    = "IMAGE ID"
	DockerImagesCreated    = "CREATED"
	DockerImagesCreatedAt  = "CREATED AT"
	DockerImagesSize       = "SIZE"
)

type DockerImages struct {
	rules rules.CommandConfig
}

func (c *DockerImages) Columns() []string {
	return []string{
		DockerImagesImageID,    //
		DockerImagesRepository, //
		DockerImagesTag,        //
		DockerImagesDigest,     // opt
		DockerImagesCreated,    //
		DockerImagesCreatedAt,  // opt
		DockerImagesSize,       //
	}
}

func (c *DockerImages) Format(row layout.Row, col layout.Column) string {
	x := string(row[col])

	if rule, ok := c.rules.Columns[string(col)]; ok {
		return rule.Pipeline.Apply(x, rules.Context{Row: row, Column: col})
	}

	return x
}
