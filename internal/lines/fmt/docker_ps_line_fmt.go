package fmt

import (
	"strings"

	"docker-color-output/pkg/color"
)

type DockerPsLineFmt struct{}

func NewDockerPsLineFmt() *DockerPsLineFmt {
	return &DockerPsLineFmt{}
}

func (*DockerPsLineFmt) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (*DockerPsLineFmt) Image(v string) string {
	parts := strings.Split(v, ":")
	if len(parts) == 2 {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}
	return color.Yellow(v)
}

func (*DockerPsLineFmt) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerPsLineFmt) Created(v string) string {
	if strings.Contains(v, "months") {
		return color.Brown(v)
	}
	if strings.Contains(v, "years") {
		return color.Red(v)
	}
	return color.Green(v)
}

func (*DockerPsLineFmt) Status(v string) string {
	if strings.Contains(v, "Exited") {
		return color.Red(v)
	}
	return color.LightGreen(v)
}

func (*DockerPsLineFmt) Ports(v string) string {
	var ports []string
	for _, port := range strings.Split(v, ", ") {
		parts := strings.Split(port, "->")
		if len(parts) == 2 {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}
		ports = append(ports, port)
	}
	return strings.Join(ports, ", ")
}

func (*DockerPsLineFmt) Names(v string) string {
	return color.White(v)
}

func (f *DockerPsLineFmt) Format(col string, v string) string {
	switch col {
	case "CONTAINER ID":
		return f.ContainerID(v)
	case "IMAGE":
		return f.Image(v)
	case "COMMAND":
		return f.Command(v)
	case "CREATED":
		return f.Created(v)
	case "STATUS":
		return f.Status(v)
	case "PORTS":
		return f.Ports(v)
	case "NAMES":
		return f.Names(v)
	default:
		return v
	}
}
