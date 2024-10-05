package cmd

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/pkg/color"
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

type DockerComposePs struct{}

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

	switch col {
	case DockerComposePsName:
		return c.Name(x, row)
	case DockerComposePsImage:
		return c.Image(x)
	case DockerComposePsCommand:
		return c.Command(x)
	case DockerComposePsService:
		return c.Service(x, row)
	case DockerComposePsCreated:
		return c.Created(x)
	case DockerComposePsStatus:
		return c.Status(x)
	case DockerComposePsPorts:
		return c.Ports(x)
	default:
		return x
	}
}

func (*DockerComposePs) Name(x string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(x)
	}

	return color.White(x)
}

func (*DockerComposePs) Image(x string) string {
	parts := strings.Split(x, ":") //nolint:ifshort
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(x)
}

func (*DockerComposePs) Command(x string) string {
	return color.DarkGray(x)
}

func (*DockerComposePs) Service(x string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(x)
	}

	return color.Yellow(x)
}

func (*DockerComposePs) Created(x string) string {
	if strings.Contains(x, "months") {
		return color.Brown(x)
	}

	if strings.Contains(x, "years") {
		return color.Red(x)
	}

	return color.Green(x)
}

func (*DockerComposePs) Status(x string) string {
	if strings.Contains(x, "exited") {
		return color.Red(x)
	}

	return color.LightGreen(x)
}

func (*DockerComposePs) Ports(x string) string {
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
