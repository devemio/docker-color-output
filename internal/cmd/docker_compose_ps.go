package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
)

const (
	DockerComposePsName    = "NAME"
	DockerComposePsImage   = "IMAGE"
	DockerComposePsCommand = "COMMAND"
	DockerComposePsService = "SERVICE"
	DockerComposePsCreated = "CREATED"
	DockerComposePsStatus  = "STATUS"
	DockerComposePsPorts   = "PORTS"
)

type DockerComposePs struct {
	rules rules.CommandConfig
}

func (c *DockerComposePs) Columns() []string {
	return []string{
		DockerComposePsName,    //
		DockerComposePsImage,   //
		DockerComposePsCommand, //
		DockerComposePsService, //
		DockerComposePsCreated, //
		DockerComposePsStatus,  //
		DockerComposePsPorts,   // nullable
	}
}

func (c *DockerComposePs) Format(row layout.Row, col layout.Column) string {
	x := string(row[col])

	if rule, ok := c.rules.Columns[string(col)]; ok {
		return rule.Pipeline.Apply(x, rules.Context{Row: row, Column: col})
	}

	return x
}
