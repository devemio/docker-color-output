package lines

import (
	"strings"

	"docker-color-output/internal/lines/fmt"
	"docker-color-output/internal/utils"
)

type Line struct {
	vals Values
	cols []Column
	fmt  fmt.Formatable
}

func NewLine(vals Values, cols []Column, fmt fmt.Formatable) *Line {
	return &Line{
		vals: vals,
		cols: cols,
		fmt:  fmt,
	}
}

func (l *Line) Build() string {
	var sb strings.Builder
	for _, col := range l.cols {
		sb.WriteString(utils.Pad(l.fmt.Format(l.vals, col.Name), *col.Len))
	}

	return sb.String()
}
