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
	v := string(row[col])

	switch col {
	case DockerComposePsName:
		return c.Name(v, row)
	case DockerComposePsImage:
		return c.Image(v)
	case DockerComposePsCommand:
		return c.Command(v)
	case DockerComposePsService:
		return c.Service(v, row)
	case DockerComposePsCreated:
		return c.Created(v)
	case DockerComposePsStatus:
		return c.Status(v)
	case DockerComposePsPorts:
		return c.Ports(v)
	default:
		return v
	}
}

func (*DockerComposePs) Name(v string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(v)
	}

	return color.White(v)
}

func (*DockerComposePs) Image(v string) string {
	parts := strings.Split(v, ":") //nolint:ifshort
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(v)
}

func (*DockerComposePs) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerComposePs) Service(v string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(v)
	}

	return color.Yellow(v)
}

func (*DockerComposePs) Created(v string) string {
	if strings.Contains(v, "months") {
		return color.Brown(v)
	}

	if strings.Contains(v, "years") {
		return color.Red(v)
	}

	return color.Green(v)
}

func (*DockerComposePs) Status(v string) string {
	if strings.Contains(v, "exited") {
		return color.Red(v)
	}

	return color.LightGreen(v)
}

func (*DockerComposePs) Ports(v string) string {
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
