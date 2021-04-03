package utils

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"os"
)

func Exit(err error, usage func()) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	usage()
	os.Exit(1)
}
