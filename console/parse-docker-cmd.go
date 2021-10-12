package console

import (
	"errors"

	"github.com/devemio/docker-color-output/utils"
)

func ParseCmd(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", errors.New("no first line")
	}

	cols := utils.Split(lines[0])

	if len(cols) < 2 {
		return "", errors.New("invalid first line")
	}

	cols = cols[:2]

	if equals(cols, []string{"REPOSITORY", "TAG"}) {
		return DockerImages, nil
	}

	if equals(cols, []string{"CONTAINER ID", "IMAGE"}) {
		return DockerPs, nil
	}

	if equals(cols, []string{"NAME", "COMMAND"}) {
		return DockerComposePs, nil
	}

	return "", errors.New("invalid cmd")
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
