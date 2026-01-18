package app

import (
	"github.com/devemio/docker-color-output/internal/stderr"
	"github.com/devemio/docker-color-output/pkg/color"
)

const indent = "    "

func Usage(err error) {
	if err != nil {
		stderr.Println(color.LightRed("💡 Error: " + err.Error()))
	}

	stderr.Println("💥 Version: " + color.Green(Ver))

	stderr.Println("👌 Usage:")
	stderr.Println(indent + color.Green("docker compose ps") + " [-a] | " + color.Brown(Name))
	stderr.Println(indent + color.Green("docker images") + " [--format] 2>/dev/null | " + color.Brown(Name))
	stderr.Println(indent + color.Green("docker ps") + " [-a] [--format] | " + color.Brown(Name))
	stderr.Println(indent + color.Green("docker stats") + " [--no-stream] | " + color.Brown(Name))

	stderr.Println("🚀 Flags:")
	stderr.Println(indent + color.Green("-c") + " " + color.Brown("string") + " Path to configuration file")
	stderr.Println(indent + color.Green("-s") + " " + color.Brown("bool") + " Silent mode (suppress errors)")
}
