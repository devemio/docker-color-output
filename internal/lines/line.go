package lines

import (
	"strings"

	"docker-color-output/internal/lines/fmt"
	"docker-color-output/internal/utils"
)

type Line struct {
	values []string
	cols   []Column
	fmt    fmt.Formatable
}

func NewLine(values []string, cols []Column, fmt fmt.Formatable) *Line {
	return &Line{
		values: values,
		cols:   cols,
		fmt:    fmt,
	}
}

func (l *Line) Build() string {
	var sb strings.Builder

	for i, col := range l.cols {
		v := l.values[i]
		v = l.fmt.Format(col.Name, v)
		v = utils.Pad(v, *col.Len)
		sb.WriteString(v)
	}

	return sb.String()
}
