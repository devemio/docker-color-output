package fmt

import (
	"strings"

	"docker-color-output/pkg/color"
)

type DockerComposePsLineFmt struct{}

func NewDockerComposePsLineFmt() *DockerComposePsLineFmt {
	return &DockerComposePsLineFmt{}
}

func (*DockerComposePsLineFmt) Name(v string) string {
	if strings.Contains(v, "exited") { // @fixme status
		return color.DarkGray(v)
	}
	return color.White(v)
}

func (*DockerComposePsLineFmt) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerComposePsLineFmt) Service(v string) string {
	if strings.Contains(v, "exited") { // @fixme status
		return color.DarkGray(v)
	}
	return color.Yellow(v)
}

func (*DockerComposePsLineFmt) Status(v string) string {
	if strings.Contains(v, "exited") {
		return color.Red(v)
	}
	return color.LightGreen(v)
}

func (*DockerComposePsLineFmt) Ports(v string) string {
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

func (f *DockerComposePsLineFmt) Format(col string, v string) string {
	switch col {
	case "NAME":
		return f.Name(v)
	case "COMMAND":
		return f.Command(v)
	case "SERVICE":
		return f.Service(v)
	case "STATUS":
		return f.Status(v)
	case "PORTS":
		return f.Ports(v)
	default:
		return v
	}
}