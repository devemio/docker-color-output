package stdout

import (
	"fmt"
	"strings"

	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
)

type DockerComposePsLine struct {
	name    string
	command string
	service string
	status  string
	ports   string
}

func (line *DockerComposePsLine) Name() string {
	if strings.Contains(line.status, "exited") {
		return color.DarkGray(line.name)
	}
	return color.White(line.name)
}

func (line *DockerComposePsLine) Command() string {
	return color.DarkGray(line.command)
}

func (line *DockerComposePsLine) Service() string {
	if strings.Contains(line.status, "exited") {
		return color.DarkGray(line.service)
	}
	return color.Yellow(line.service)
}

func (line *DockerComposePsLine) Status() string {
	if strings.Contains(line.status, "exited") {
		return color.Red(line.status)
	}
	return color.LightGreen(line.status)
}

func (line *DockerComposePsLine) Ports() string {
	var ports []string
	for _, port := range strings.Split(line.ports, ", ") {
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
		Format(line.Service(), lens[2]),
		Format(line.Status(), lens[3]),
		Format(line.Ports(), lens[4]),
	)
}

func CreateDockerComposePsLine(line string) *DockerComposePsLine {
	cols := utils.Split(line)
	return &DockerComposePsLine{
		name:    cols[0],
		command: cols[1],
		service: cols[2],
		status:  cols[3],
		ports:   cols[4],
	}
}
