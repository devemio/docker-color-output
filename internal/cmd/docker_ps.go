package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
)

const (
	DockerPsContainerID = "CONTAINER ID"
	DockerPsImage       = "IMAGE"
	DockerPsCommand     = "COMMAND"
	DockerPsCreatedAt   = "CREATED AT"
	DockerPsCreated     = "CREATED"
	DockerPsPorts       = "PORTS"
	DockerPsState       = "STATE"
	DockerPsStatus      = "STATUS"
	DockerPsSize        = "SIZE"
	DockerPsNames       = "NAMES"
	DockerPsLabels      = "LABELS"
	DockerPsMounts      = "MOUNTS"
	DockerPsNetworks    = "NETWORKS"
)

type DockerPs struct {
	rules rules.CommandConfig
}

func (c *DockerPs) Columns() []string {
	return []string{
		DockerPsContainerID, //
		DockerPsImage,       //
		DockerPsCommand,     //
		DockerPsCreatedAt,   // opt
		DockerPsCreated,     //
		DockerPsPorts,       // nullable
		DockerPsState,       // opt
		DockerPsStatus,      //
		DockerPsSize,        // opt
		DockerPsNames,       //
		DockerPsLabels,      // opt
		DockerPsMounts,      // opt | nullable
		DockerPsNetworks,    // opt | nullable
	}
}

func (c *DockerPs) Format(rows layout.Row, col layout.Column) string {
	x := string(rows[col])

	if rule, ok := c.rules.Columns[string(col)]; ok {
		return rule.Pipeline.Apply(x, rules.Context{Row: rows, Column: col})
	}

	return x
}
