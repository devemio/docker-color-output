package stdout

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
	"strings"
)

type DockerComposePsLine struct {
	name    string
	command string
	state   string
	ports   string
}

func (line *DockerComposePsLine) Name() string {
	if strings.Contains(line.state, "Exit") {
		return line.name
	}
	return color.White(line.name)
}

func (line *DockerComposePsLine) Command() string {
	return color.DarkGray(line.command)
}

func (line *DockerComposePsLine) State() string {
	if strings.Contains(line.state, "Exit") {
		return color.Red(line.state)
	}
	return color.LightGreen(line.state)
}

func (line *DockerComposePsLine) Ports() string {
	var ports []string
	for _, port := range strings.Split(line.ports, ",") {
		parts := strings.Split(port, "->")
		if len(parts) == 2 {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}
		ports = append(ports, port)
	}
	return strings.Join(ports, ", ")
}

func (line *DockerComposePsLine) Println(lens []int) {
	fmt.Println(
		Format(line.Name(), lens[0]),
		Format(line.Command(), lens[1]),
		Format(line.State(), lens[2]),
		Format(line.Ports(), lens[3]),
	)
}

func CreateDockerComposePsLine(line string) Line {
	cols := utils.Split(line)
	return &DockerComposePsLine{
		name:    cols[0],
		command: cols[1],
		state:   cols[2],
		ports:   cols[3],
	}
}
