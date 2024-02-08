package cmd

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/pkg/color"
	"github.com/devemio/docker-color-output/pkg/util/number"
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

const (
	cpuPercentThresholdMedium = 50
	cpuPercentThresholdHigh   = 90

	memPercentThresholdMedium = 50
	memPercentThresholdHigh   = 90

	netIOThresholdHigh = 10

	blockIOThresholdHigh = 10

	pidsThresholdHigh = 100
)

type DockerStats struct{}

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
	v := string(rows[col])

	switch col {
	case DockerStatsContainerID:
		return c.ContainerID(v)
	case DockerStatsName:
		return c.Name(v)
	case DockerStatsCPUPercent:
		return c.CPUPercent(v)
	case DockerStatsMemUsage:
		return c.MemUsage(v)
	case DockerStatsMemPercent:
		return c.MemPercent(v)
	case DockerStatsNetIO:
		return c.NetIO(v)
	case DockerStatsBlockIO:
		return c.BlockIO(v)
	case DockerStatsPIDs:
		return c.PIDs(v)
	default:
		return v
	}
}

func (*DockerStats) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (c *DockerStats) Name(v string) string {
	return color.White(v)
}

func (c *DockerStats) CPUPercent(v string) string {
	percent := number.ParseFloat(v)

	switch {
	case percent >= cpuPercentThresholdHigh:
		return color.Red(v)
	case percent >= cpuPercentThresholdMedium:
		return color.Brown(v)
	default:
		return v
	}
}

func (c *DockerStats) MemUsage(v string) string {
	parts := strings.Split(v, "/")

	return parts[0] + color.DarkGray("/"+parts[1])
}

func (c *DockerStats) MemPercent(v string) string {
	percent := number.ParseFloat(v)

	switch {
	case percent >= memPercentThresholdHigh:
		return color.Red(v)
	case percent >= memPercentThresholdMedium:
		return color.Brown(v)
	default:
		return v
	}
}

func (*DockerStats) NetIO(v string) string {
	parts := strings.Split(v, "/")

	for i := range parts {
		if strings.Contains(parts[i], "GB") {
			if number.ParseFloat(parts[i]) >= netIOThresholdHigh {
				parts[i] = color.Red(parts[i])
			} else {
				parts[i] = color.Brown(parts[i])
			}
		}
	}

	return parts[0] + color.DarkGray("/") + parts[1]
}

func (*DockerStats) BlockIO(v string) string {
	parts := strings.Split(v, "/")

	for i := range parts {
		if strings.Contains(parts[i], "GB") {
			if number.ParseFloat(parts[i]) >= blockIOThresholdHigh {
				parts[i] = color.Red(parts[i])
			} else {
				parts[i] = color.Brown(parts[i])
			}
		}
	}

	return parts[0] + color.DarkGray("/") + parts[1]
}

func (*DockerStats) PIDs(v string) string {
	if number.ParseFloat(v) >= pidsThresholdHigh {
		return color.Red(v)
	}

	return color.LightCyan(v)
}
