package main

import (
	"github.com/devemio/docker-color-output/console"
	"github.com/devemio/docker-color-output/create"
	"github.com/devemio/docker-color-output/stdin"
	"github.com/devemio/docker-color-output/stdout"
	"github.com/devemio/docker-color-output/utils"
)

func main() {
	// Parse lines
	lines, _lines := stdin.GetLines()
	if _lines != nil {
		utils.Exit(_lines, console.Usage)
	}

	// Parse cmd
	cmd, _cmd := console.ParseCmd(lines)
	if _cmd != nil {
		utils.Exit(_cmd, console.Usage)
	}

	// Get max column lengths
	lens := utils.GetMaxLens(lines, create.Split(cmd))

	// Print first line
	stdout.CreateFirstLine(lines[0]).Println(lens)

	// Print lines
	for _, line := range lines[1:] {
		create.Line(cmd, line).Println(lens)
	}
}
