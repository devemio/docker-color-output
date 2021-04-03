package main

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/console"
	"github.com/devemio/docker-color-output/stdin"
	"github.com/devemio/docker-color-output/stdout"
	"github.com/devemio/docker-color-output/utils"
	"os"
)

func main() {
	// Parse lines
	lines, err := stdin.GetLines()
	if err != nil {
		exit(err)
	}

	// Parse cmd
	cmd, err := console.ParseCmd(lines)
	if err != nil {
		exit(err)
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

func exit(err error) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	console.Usage()
	os.Exit(1)
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
