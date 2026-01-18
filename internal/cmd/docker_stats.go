package cmd

import (
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/rules"
)

const (
	DockerStatsContainerID = "CONTAINER ID"
	DockerStatsName        = "NAME"
	DockerStatsCPUPercent  = "CPU %"
	DockerStatsMemUsage    = "MEM USAGE / LIMIT"
	DockerStatsMemPercent  = "MEM %"
	DockerStatsNetIO       = "NET I/O"
	DockerStatsBlockIO     = "BLOCK I/O"
	DockerStatsPIDs        = "PIDS"
)

type DockerStats struct {
	rules rules.CommandConfig
}

func (c *DockerStats) Columns() []string {
	return []string{
		DockerStatsContainerID, //
		DockerStatsName,        //
		DockerStatsCPUPercent,  //
		DockerStatsMemUsage,    //
		DockerStatsMemPercent,  //
		DockerStatsNetIO,       //
		DockerStatsBlockIO,     //
		DockerStatsPIDs,        //
	}
}

func (c *DockerStats) Format(rows layout.Row, col layout.Column) string {
	x := string(rows[col])

	if rule, ok := c.rules.Columns[string(col)]; ok {
		return rule.Pipeline.Apply(x, rules.Context{Row: rows, Column: col})
	}

	return x
}
