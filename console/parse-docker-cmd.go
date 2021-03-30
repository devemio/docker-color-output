package console

import (
	"github.com/devemio/docker-color-output/utils"
)

func ParseCmd(lines []string) string {
	cols := utils.Split(lines[0])[:2]

	if equals(cols, []string{"REPOSITORY", "TAG"}) {
		return DockerImages
	}

	if equals(cols, []string{"CONTAINER ID", "IMAGE"}) {
		return DockerPs
	}

	if equals(cols, []string{"Name", "Command"}) {
		return DockerComposePs
	}

	panic("Failed to parse cmd")
}

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
