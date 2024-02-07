package cmd

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/pkg/color"
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

type DockerPs struct{}

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
	v := string(rows[col])

	switch col {
	case DockerPsContainerID:
		return c.ContainerID(v)
	case DockerPsImage:
		return c.Image(v)
	case DockerPsCommand:
		return c.Command(v)
	case DockerPsCreated:
		return c.Created(v)
	case DockerPsStatus:
		return c.Status(v)
	case DockerPsPorts:
		return c.Ports(v)
	case DockerPsNames:
		return c.Names(v)
	default:
		return v
	}
}

func (*DockerPs) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (*DockerPs) Image(v string) string {
	parts := strings.Split(v, ":") //nolint:ifshort
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(v)
}

func (*DockerPs) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerPs) Created(v string) string {
	if strings.Contains(v, "months") {
		return color.Brown(v)
	}

	if strings.Contains(v, "years") {
		return color.Red(v)
	}

	return color.Green(v)
}

func (*DockerPs) Status(v string) string {
	if strings.Contains(v, "Exited") {
		return color.Red(v)
	}

	return color.LightGreen(v)
}

func (*DockerPs) Ports(v string) string {
	ports := make([]string, 0)

	for _, port := range strings.Split(v, ", ") {
		parts := strings.Split(port, "->")
		if len(parts) == ValidTotalParts {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}

		ports = append(ports, port)
	}

	return strings.Join(ports, ", ")
}

func (*DockerPs) Names(v string) string {
	return color.White(v)
}
