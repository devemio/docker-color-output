package main

import (
	"github.com/devemio/docker-color-output/console"
	"github.com/devemio/docker-color-output/stdin"
	"github.com/devemio/docker-color-output/stdout"
	"github.com/devemio/docker-color-output/utils"
)

func main() {
	// Parse lines
	lines := stdin.GetLines()

	// Parse cmd
	cmd := console.ParseCmd(lines)

	// Get max column lengths
	lens := utils.GetMaxLens(lines)

	// Print first line
	stdout.CreateFirstLine(lines[0]).Println(lens)

	// Print lines
	for _, line := range lines[1:] {
		createLine(cmd, line).Println(lens)
	}
}

func createLine(cmd string, line string) stdout.LinePrinter {
	switch cmd {
	case console.DockerImages:
		return stdout.CreateDockerImageLine(line)
	case console.DockerPs:
		return stdout.CreateDockerPsLine(line)
	case console.DockerComposePs:
		return nil // @fixme
	default:
		return nil
	}
}
