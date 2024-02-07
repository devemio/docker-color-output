package color

type Palette struct {
	Reset       string
	Black       string
	DarkGray    string
	Red         string
	LightRed    string
	Green       string
	LightGreen  string
	Brown       string
	Yellow      string
	Blue        string
	LightBlue   string
	Purple      string
	LightPurple string
	Cyan        string
	LightCyan   string
	LightGray   string
	White       string
}

//nolint:gochecknoglobals
var palette = Palette{
	Reset:       "\u001B[0m",
	Black:       "\u001B[0;30m",
	DarkGray:    "\u001B[1;30m",
	Red:         "\u001B[0;31m",
	LightRed:    "\u001B[1;31m",
	Green:       "\u001B[0;32m",
	LightGreen:  "\u001B[1;32m",
	Brown:       "\u001B[0;33m",
	Yellow:      "\u001B[1;33m",
	Blue:        "\u001B[0;34m",
	LightBlue:   "\u001B[1;34m",
	Purple:      "\u001B[0;35m",
	LightPurple: "\u001B[1;35m",
	Cyan:        "\u001B[0;36m",
	LightCyan:   "\u001B[1;36m",
	LightGray:   "\u001B[0;37m",
	White:       "\u001B[1;37m",
}

func SetPalette(p Palette) {
	palette = p
}

func Black(value string) string {
	return palette.Black + value + palette.Reset
}

func DarkGray(value string) string {
	return palette.DarkGray + value + palette.Reset
}

func Red(value string) string {
	return palette.Red + value + palette.Reset
}

func LightRed(value string) string {
	return palette.LightRed + value + palette.Reset
}

func Green(value string) string {
	return palette.Green + value + palette.Reset
}

func LightGreen(value string) string {
	return palette.LightGreen + value + palette.Reset
}

func Brown(value string) string {
	return palette.Brown + value + palette.Reset
}

func Yellow(value string) string {
	return palette.Yellow + value + palette.Reset
}

func Blue(value string) string {
	return palette.Blue + value + palette.Reset
}

func LightBlue(value string) string {
	return palette.LightBlue + value + palette.Reset
}

func Purple(value string) string {
	return palette.Purple + value + palette.Reset
}

func LightPurple(value string) string {
	return palette.LightPurple + value + palette.Reset
}

func Cyan(value string) string {
	return palette.Cyan + value + palette.Reset
}

func LightCyan(value string) string {
	return palette.LightCyan + value + palette.Reset
}

func LightGray(value string) string {
	return palette.LightGray + value + palette.Reset
}

func White(value string) string {
	return palette.White + value + palette.Reset
}
