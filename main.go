package main

import (
	"github.com/devemio/docker-color-output/console"
	"github.com/devemio/docker-color-output/stdin"
	"github.com/devemio/docker-color-output/stdout"
	"github.com/devemio/docker-color-output/utils"
)

func main() {
	// Parse lines
	lines, eLines := stdin.GetLines()
	if eLines != nil {
		exit(eLines)
	}

	// Parse cmd
	cmd, eCmd := console.ParseCmd(lines)
	if eCmd != nil {
		exit(eCmd)
	}

	// Get max column lengths
	lens := utils.GetMaxLens(lines)

	// Print first line
	stdout.CreateFirstLine(lines[0]).Println(lens)

	// Print lines
	for _, line := range lines[1:] {
		createLine(cmd, line).Println(lens)
	}
}

func createLine(cmd string, line string) stdout.Line {
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
