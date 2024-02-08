//go:build !test

package app

import (
	"github.com/devemio/docker-color-output/internal/stdout"
	"github.com/devemio/docker-color-output/pkg/color"
)

const indent = "    "

func Usage(err error) {
	stdout.Println(color.LightRed("💡 Error: " + err.Error()))
	stdout.Println("💥 Version: " + color.Green(Ver))
	stdout.Println("👌 Usage:")
	stdout.Println(indent + color.Green("docker compose ps") + " [-a] | " + color.Brown(Name))
	stdout.Println(indent + color.Green("docker images") + " [--format] | " + color.Brown(Name))
	stdout.Println(indent + color.Green("docker ps") + " [-a] [--format] | " + color.Brown(Name))
	stdout.Println(indent + color.Green("docker stats") + " [--no-stream] | " + color.Brown(Name))
}
