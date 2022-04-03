package fmt

type Formatable interface {
	Format(col string, v string) string
}
