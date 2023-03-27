//go:build !test

package app

import (
	"fmt"

	"docker-color-output/pkg/color"
)

//nolint:forbidigo
func Usage(err error) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	fmt.Println("💥 Version: " + color.Green(Ver))
	fmt.Println("👌 Usage:")
	fmt.Println("    " + color.Green("docker compose ps") + " [-a] | " + color.Brown(Name))
	fmt.Println("    " + color.Green("docker images") + " [--format] | " + color.Brown(Name))
	fmt.Println("    " + color.Green("docker ps") + " [-a] [--format] | " + color.Brown(Name))
}
