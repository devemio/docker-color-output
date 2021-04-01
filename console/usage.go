package console

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  " + color.Green(DockerImages) + " | " + color.Brown(App))
	fmt.Println("  " + color.Green(DockerPs) + " [-a] | " + color.Brown(App))
	fmt.Println("  " + color.Green(DockerComposePs) + " [-a] | " + color.Brown(App))
}
