package cli

import (
	"docker-color-output/internal/cmd"
	"docker-color-output/internal/lines"
	"docker-color-output/internal/lines/fmt"
)

func Execute(in []string) ([]*lines.Line, error) {
	c, err := cmd.ParseCmd(&in)
	if err != nil {
		return nil, err
	}

	res := make([]*lines.Line, len(in))

	cols, vals := cmd.ParseFirstLine(in[0])

	res[0] = lines.NewLine(vals, cols, fmt.NewFirstLineFmt())

	pl, err := cmd.ParseLines(in[1:], cols)
	if err != nil {
		return nil, err
	}

	for i, vals := range pl {
		res[i+1] = lines.NewLine(vals, cols, c.GetFmt())
	}

	return res, nil
}
