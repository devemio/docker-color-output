package main

import (
	"fmt"
	"github.com/devemio/docker-color-output/color"
	"github.com/devemio/docker-color-output/console"
	"os"
)

func exit(err error) {
	fmt.Println(color.LightRed("💡 Error: " + err.Error()))
	console.Usage()
	os.Exit(1)
}
