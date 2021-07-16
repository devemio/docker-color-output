package stdout

import (
	"fmt"
	"strings"

	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
)

type DockerImageLine struct {
	repository string
	tag        string
	imageId    string
	created    string
	size       string
}

func (line *DockerImageLine) Repository() string {
	if strings.Contains(line.repository, "/") {
		return color.DarkGray(line.repository)
	}
	return color.White(line.repository)
}

func (line *DockerImageLine) Tag() string {
	if line.tag == "latest" {
		return color.LightGreen(line.tag)
	}
	return line.tag
}

func (line *DockerImageLine) ImageId() string {
	return color.DarkGray(line.imageId)
}

func (line *DockerImageLine) Created() string {
	if strings.Contains(line.created, "days") {
		return color.Green(line.created)
	}
	if strings.Contains(line.created, "weeks") {
		return color.Green(line.created)
	}
	if strings.Contains(line.created, "months") {
		return color.Brown(line.created)
	}
	if strings.Contains(line.created, "years") {
		return color.Red(line.created)
	}
	return line.created
}

func (line *DockerImageLine) Size() string {
	if strings.Contains(line.size, "GB") {
		return color.Red(line.size)
	}
	if strings.Contains(line.size, "MB") && utils.ParseFloat(line.size) >= 500 {
		return color.Brown(line.size)
	}
	return line.size
}

func (line *DockerImageLine) Println(lens []int) {
	fmt.Println(
		Format(line.Repository(), lens[0]),
		Format(line.Tag(), lens[1]),
		Format(line.ImageId(), lens[2]),
		Format(line.Created(), lens[3]),
		Format(line.Size(), lens[4]),
	)
}

func CreateDockerImageLine(line string) *DockerImageLine {
	cols := utils.Split(line)
	return &DockerImageLine{
		repository: cols[0],
		tag:        cols[1],
		imageId:    cols[2],
		created:    cols[3],
		size:       cols[4],
	}
}
