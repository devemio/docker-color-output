//go:build !test

package cli

import (
	"fmt"
	"os"

	"docker-color-output/internal/cmd"
	"docker-color-output/pkg/color"
)

func Exit(err error) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	fmt.Println("💥 Docker color output " + color.Green(Ver))
	fmt.Println("⚡️ Usage:")
	fmt.Println("    " + color.Green(cmd.DockerComposePs.String()) + " [-a] | " + color.Brown(App))
	fmt.Println("    " + color.Green(cmd.DockerImages.String()) + " [--format] | " + color.Brown(App))
	fmt.Println("    " + color.Green(cmd.DockerPs.String()) + " [-a] [--format] | " + color.Brown(App))
	os.Exit(1)
}
