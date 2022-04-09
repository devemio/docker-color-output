package fmt

type Formatable interface {
	Format(vals map[string]string, col string) string
}
