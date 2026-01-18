package layout

//nolint:gochecknoglobals
var nullableCols = map[Column]struct{}{
	"PORTS":    {},
	"MOUNTS":   {},
	"NETWORKS": {},
	"EXTRA":    {},
}

type Cell struct {
	Name      Column
	MaxLength int
}

func (c *Cell) IsNullable() bool {
	_, ok := nullableCols[c.Name]

	return ok
}

type Header []*Cell

func (h Header) NullableCols() byte {
	var res byte

	for _, col := range h {
		if col.IsNullable() {
			res++
		}
	}

	return res
}
