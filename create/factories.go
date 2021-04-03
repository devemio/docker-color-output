package create

import (
	"github.com/devemio/docker-color-output/console"
	"github.com/devemio/docker-color-output/stdout"
	"github.com/devemio/docker-color-output/utils"
)

func Split(cmd string) utils.SplitContract {
	switch cmd {
	case console.DockerPs:
		return stdout.SplitDockerPsLine
	default:
		return utils.Split
	}
}

func Line(cmd string, line string) stdout.Line {
	switch cmd {
	case console.DockerImages:
		return stdout.CreateDockerImageLine(line)
	case console.DockerPs:
		return stdout.CreateDockerPsLine(line)
	case console.DockerComposePs:
		return stdout.CreateDockerComposePsLine(line)
	default:
		return nil
	}
}
