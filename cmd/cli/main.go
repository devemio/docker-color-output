package main

import (
	"docker-color-output/internal/cli"
	"docker-color-output/internal/cmd"
	"docker-color-output/internal/lines"
	"docker-color-output/internal/lines/fmt"
	"docker-color-output/internal/stdin"
	"docker-color-output/internal/stdout"
)

func main() {
	in, err := stdin.Get()
	if err != nil {
		cli.Exit(err)
	}

	c, err := cmd.ParseCmd(in)
	if err != nil {
		cli.Exit(err)
	}

	res := make([]*lines.Line, len(in))

	cols, values := cmd.ParseFirstLine(in[0])
	res[0] = lines.NewLine(values, cols, fmt.NewFirstLineFmt())

	for i, values := range cmd.ParseLines(in[1:], cols) {
		res[i+1] = lines.NewLine(values, cols, c.GetFmt())
	}

	for _, v := range res {
		stdout.Println(v)
	}
}
