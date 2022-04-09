package lines

type Buildable interface {
	Build() string
}

type Values map[string]string

type Column struct {
	Name string
	Len  *int
}

func (c *Column) IsNullable() bool {
	return c.Name == "PORTS" || c.Name == "MOUNTS" || c.Name == "NETWORKS"
}
