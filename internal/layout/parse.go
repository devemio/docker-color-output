package layout

import (
	"strings"

	"github.com/devemio/docker-color-output/internal/util"
)

func ParseHeader(rows []string) Header {
	parts := util.Split(rows[0])

	res := make(Header, len(parts))
	for i, part := range parts {
		res[i] = &Cell{
			Name:      Column(part),
			MaxLength: len(part),
		}
	}

	return res
}

func ParseRows(rows []string, header *Header) []Row {
	res := make([]Row, len(rows)-1)

	for i, row := range rows[1:] {
		offset := 0
		parts := util.Split(row)
		res[i] = make(Row, len(*header))
		mismatch := len(parts) < len(*header)

		for j, col := range *header {
			if mismatch && col.IsNullable() {
				offset++

				continue
			}

			x := parts[j-offset]

			length := len(x)
			if strings.Contains(x, "…") {
				length -= 2
			}

			if length > col.MaxLength {
				col.MaxLength = length
			}

			res[i][col.Name] = Value(x)
		}
	}

	return res
}
