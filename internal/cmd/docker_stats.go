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
	x := string(rows[col])

	switch col {
	case DockerStatsContainerID:
		return c.ContainerID(x)
	case DockerStatsName:
		return c.Name(x)
	case DockerStatsCPUPercent:
		return c.CPUPercent(x)
	case DockerStatsMemUsage:
		return c.MemUsage(x)
	case DockerStatsMemPercent:
		return c.MemPercent(x)
	case DockerStatsNetIO:
		return c.NetIO(x)
	case DockerStatsBlockIO:
		return c.BlockIO(x)
	case DockerStatsPIDs:
		return c.PIDs(x)
	default:
		return x
	}
}

func (*DockerStats) ContainerID(x string) string {
	return color.DarkGray(x)
}

func (c *DockerStats) Name(x string) string {
	return color.White(x)
}

func (c *DockerStats) CPUPercent(x string) string {
	percent := number.ParseFloat(x)

	switch {
	case percent >= cpuPercentThresholdHigh:
		return color.Red(x)
	case percent >= cpuPercentThresholdMedium:
		return color.Brown(x)
	default:
		return x
	}
}

func (c *DockerStats) MemUsage(x string) string {
	parts := strings.Split(x, "/")

	return parts[0] + color.DarkGray("/"+parts[1])
}

func (c *DockerStats) MemPercent(x string) string {
	percent := number.ParseFloat(x)

	switch {
	case percent >= memPercentThresholdHigh:
		return color.Red(x)
	case percent >= memPercentThresholdMedium:
		return color.Brown(x)
	default:
		return x
	}
}

func (*DockerStats) NetIO(x string) string {
	parts := strings.Split(x, "/")

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

func (*DockerStats) BlockIO(x string) string {
	parts := strings.Split(x, "/")

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

func (*DockerStats) PIDs(x string) string {
	if number.ParseFloat(x) >= pidsThresholdHigh {
		return color.Red(x)
	}

	return color.LightCyan(x)
}
