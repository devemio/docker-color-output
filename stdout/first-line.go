package stdout

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
)

type FirstLine struct {
	cols []string
}

func (line *FirstLine) Println(lens []int) {
	for i, v := range line.cols {
		fmt.Print(Format(color.LightBlue(v), lens[i]) + " ")
	}
	fmt.Println()
}

func CreateFirstLine(line string) *FirstLine {
	return &FirstLine{
		cols: utils.Split(line),
	}
}
