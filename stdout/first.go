package stdout

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/utils"
	"strings"
)

type FirstLine struct {
	cols []string
}

func (line *FirstLine) Println(lens []int) {
	for i, v := range line.cols {
		fmt.Print(Format(color.LightBlue(strings.ToUpper(v)), lens[i]) + " ")
	}
	fmt.Println()
}

func CreateFirstLine(line string) Line {
	return &FirstLine{
		cols: utils.Split(line),
	}
}
