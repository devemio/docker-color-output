package cmd

import (
	"errors"
	"strings"

	"docker-color-output/internal/lines"
	"docker-color-output/internal/utils"
	"docker-color-output/internal/utils/pointer"
)

func ParseCmd(in []string) (Cmd, error) {
	if len(in) == 0 || len(in[0]) == 0 {
		return "", errors.New("no first line")
	}

	parts := utils.Split(in[0])

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

func ParseFirstLine(in string) ([]lines.Column, lines.Values) {
	parts := utils.Split(in)

	cols := make([]lines.Column, len(parts))
	vals := make(lines.Values, len(parts))

	for i, part := range parts {
		cols[i] = lines.Column{
			Name: part,
			Len:  pointer.ToInt(len(part)),
		}

		vals[part] = part
	}

	return cols, vals
}

func ParseLines(ins []string, cols []lines.Column) ([]lines.Values, error) {
	if calculateNullableCols(cols) > 1 {
		return nil, errors.New("nullable columns more than one")
	}

	rows := make([]lines.Values, len(ins))

	for i, in := range ins {
		rows[i] = make(lines.Values, len(cols))

		parts := utils.Split(in)

		offset := 0
		mismatch := len(parts) < len(cols)

		for j, col := range cols {
			if mismatch && col.IsNullable() {
				offset++
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

			rows[i][col.Name] = v
		}
	}

	return rows, nil
}

func calculateNullableCols(cols []lines.Column) (total byte) {
	for _, col := range cols {
		if col.IsNullable() {
			total++
		}
	}

	return
}
