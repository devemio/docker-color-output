package stdout

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
	"strings"
)

type DockerPsLine struct {
	containerId string
	image       string
	command     string
	created     string
	status      string
	ports       string
	names       string
}

func (line DockerPsLine) ContainerId() string {
	return color.DarkGray(line.containerId)
}

func (line DockerPsLine) Image() string {
	parts := strings.Split(line.image, ":")
	if len(parts) == 2 {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}
	return color.Yellow(line.image)
}

func (line DockerPsLine) Command() string {
	return color.DarkGray(line.command)
}

func (line DockerPsLine) Created() string {
	if strings.Contains(line.created, "second") {
		return color.Green(line.created)
	}
	if strings.Contains(line.created, "minute") {
		return color.Green(line.created)
	}
	if strings.Contains(line.created, "hour") {
		return color.Green(line.created)
	}
	if strings.Contains(line.created, "days") {
		return color.Brown(line.created)
	}
	if strings.Contains(line.created, "months") {
		return color.Red(line.created)
	}
	return line.created
}

func (line DockerPsLine) Status() string {
	if strings.Contains(line.status, "Exited") {
		return color.Red(line.status)
	}
	return color.LightGreen(line.status)
}

func (line DockerPsLine) Ports() string {
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

func (line DockerPsLine) Names() string {
	return color.White(line.names)
}

func (line DockerPsLine) Println(lens []int) {
	fmt.Println(
		Format(line.ContainerId(), lens[0]),
		Format(line.Image(), lens[1]),
		Format(line.Command(), lens[2]),
		Format(line.Created(), lens[3]),
		Format(line.Status(), lens[4]),
		Format(line.Ports(), lens[5]),
		Format(line.Names(), lens[6]),
	)
}

func CreateDockerPsLine(line string) *DockerPsLine {
	cols := utils.Split(line)
	return &DockerPsLine{
		containerId: cols[0],
		image:       cols[1],
		command:     cols[2],
		created:     cols[3],
		status:      cols[4],
		ports:       cols[5],
		names:       cols[6],
	}
}
