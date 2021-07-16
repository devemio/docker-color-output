package utils

import (
	"fmt"
	"os"

	"github.com/devemio/docker-color-output/color"
)

func Exit(err error, usage func()) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	usage()
	os.Exit(1)
}
