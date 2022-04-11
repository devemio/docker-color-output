package fmt

import (
	"strings"

	"docker-color-output/pkg/color"
)

type DockerComposePsV1LineFmt struct{}

func NewDockerComposePsV1LineFmt() *DockerComposePsV1LineFmt {
	return &DockerComposePsV1LineFmt{}
}

func (*DockerComposePsV1LineFmt) Name(v, state string) string {
	if strings.Contains(state, "Exit") {
		return color.DarkGray(v)
	}
	return color.White(v)
}

func (*DockerComposePsV1LineFmt) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerComposePsV1LineFmt) State(v string) string {
	if strings.Contains(v, "Exit") {
		return color.Red(v)
	}
	return color.LightGreen(v)
}

func (*DockerComposePsV1LineFmt) Ports(v string) string {
	var ports []string
	for _, port := range strings.Split(v, ",") {
		parts := strings.Split(port, "->")
		if len(parts) == 2 {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}
		ports = append(ports, port)
	}
	return strings.Join(ports, ", ")
}

func (f *DockerComposePsV1LineFmt) Format(vals map[string]string, col string) string {
	v := vals[col]
	switch col {
	case "NAME":
		return f.Name(v, vals["STATE"])
	case "COMMAND":
		return f.Command(v)
	case "STATE":
		return f.State(v)
	case "PORTS":
		return f.Ports(v)
	default:
		return v
	}
}
