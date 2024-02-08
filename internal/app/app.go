package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/devemio/docker-color-output/internal/cmd"
	"github.com/devemio/docker-color-output/internal/layout"
	"github.com/devemio/docker-color-output/internal/util"
	"github.com/devemio/docker-color-output/pkg/color"
)

var (
	ErrNoFirstLine     = errors.New("no first line")
	ErrNullableColumns = errors.New("nullable columns more than one")
)

func Run(in []string) ([]string, error) {
	if len(in) == 0 {
		return nil, ErrNoFirstLine
	}

	header := layout.ParseHeader(in)
	rows := layout.ParseRows(in, &header)

	if header.NullableCols() > 1 {
		return nil, ErrNullableColumns
	}

	command, err := cmd.Parse(header)
	if err != nil {
		return nil, fmt.Errorf("cmd: parse: %w", err)
	}

	res := make([]string, len(in))

	// First line
	var sb strings.Builder
	for _, col := range header {
		sb.WriteString(util.Pad(color.LightBlue(string(col.Name)), col.MaxLength))
	}

	res[0] = sb.String()

	// Rows
	for i, row := range rows {
		sb.Reset()

		for _, col := range header {
			sb.WriteString(util.Pad(command.Format(row, col.Name), col.MaxLength))
		}

		res[i+1] = sb.String()
	}

	return res, nil
}
