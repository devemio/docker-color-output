//go:build !test

package app

import (
	"fmt"

	"github.com/devemio/docker-color-output/pkg/color"
)

const indent = "    "

//nolint:forbidigo
func Usage(err error) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	fmt.Println("💥 Version: " + color.Green(Ver))
	fmt.Println("👌 Usage:")
	fmt.Println(indent + color.Green("docker compose ps") + " [-a] | " + color.Brown(Name))
	fmt.Println(indent + color.Green("docker images") + " [--format] | " + color.Brown(Name))
	fmt.Println(indent + color.Green("docker ps") + " [-a] [--format] | " + color.Brown(Name))
}
