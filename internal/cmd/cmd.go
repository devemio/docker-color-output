package cmd

import (
	"errors"
	"log"
	"strings"

	"docker-color-output/internal/lines"
	"docker-color-output/internal/utils"
	"docker-color-output/internal/utils/pointer"
	"docker-color-output/pkg/color"
)

func ParseCmd(in []string) (Cmd, error) {
	if len(in) == 0 {
		return "", errors.New("no first line")
	}

	parts := utils.Split(in[0])
	if len(parts) == 0 {
		return "", errors.New("no first line")
	}

	if utils.Intersect(parts, DockerPs.Columns()) {
		return DockerPs, nil
	}

	if utils.Intersect(parts, DockerImages.Columns()) {
		return DockerImages, nil
	}

	if utils.Intersect(parts, DockerComposePs.Columns()) {
		return DockerComposePs, nil
	}

	return "", errors.New("invalid first line")
}

func ParseFirstLine(in string) ([]lines.Column, []string) {
	parts := utils.Split(in)

	cols := make([]lines.Column, len(parts))
	values := make([]string, len(parts))

	for i, part := range parts {
		cols[i] = lines.Column{
			Name: part,
			Len:  pointer.ToInt(len(part)),
		}

		values[i] = part
	}

	return cols, values
}

func ParseLines(ins []string, cols []lines.Column) [][]string {
	rows := make([][]string, len(ins))

	for i, in := range ins {
		rows[i] = make([]string, len(cols))

		parts := utils.Split(in)

		offset := 0
		mismatches := len(cols) - len(parts)
		if mismatches > 1 {
			log.Fatal(color.Red("nullable columns more than two")) // panic
		}

		for j, col := range cols {
			if mismatches > 0 && (col.Name == "PORTS" || col.Name == "MOUNTS") {
				offset += 1
				continue
			}

			v := parts[j-offset]

			l := len(v)
			if strings.Contains(v, "…") {
				l -= 2
			}

			if l > *col.Len {
				*col.Len = l
			}

			rows[i][j] = v
		}
	}

	return rows
}
