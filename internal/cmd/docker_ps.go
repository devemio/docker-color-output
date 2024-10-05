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
	x := string(rows[col])

	switch col {
	case DockerPsContainerID:
		return c.ContainerID(x)
	case DockerPsImage:
		return c.Image(x)
	case DockerPsCommand:
		return c.Command(x)
	case DockerPsCreated:
		return c.Created(x)
	case DockerPsStatus:
		return c.Status(x)
	case DockerPsPorts:
		return c.Ports(x)
	case DockerPsNames:
		return c.Names(x)
	default:
		return x
	}
}

func (*DockerPs) ContainerID(x string) string {
	return color.DarkGray(x)
}

func (*DockerPs) Image(x string) string {
	parts := strings.Split(x, ":") //nolint:ifshort
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(x)
}

func (*DockerPs) Command(x string) string {
	return color.DarkGray(x)
}

func (*DockerPs) Created(x string) string {
	if strings.Contains(x, "months") {
		return color.Brown(x)
	}

	if strings.Contains(x, "years") {
		return color.Red(x)
	}

	return color.Green(x)
}

func (*DockerPs) Status(x string) string {
	if strings.Contains(x, "Exited") {
		return color.Red(x)
	}

	return color.LightGreen(x)
}

func (*DockerPs) Ports(x string) string {
	ports := make([]string, 0)

	for _, port := range strings.Split(x, ", ") {
		parts := strings.Split(port, "->")
		if len(parts) == ValidTotalParts {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}

		ports = append(ports, port)
	}

	return strings.Join(ports, ", ")
}

func (*DockerPs) Names(x string) string {
	return color.White(x)
}
