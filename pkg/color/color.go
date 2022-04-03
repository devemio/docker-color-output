package color

const (
	reset       = "\033[0m"
	black       = "\033[0;30m"
	darkGray    = "\033[1;30m"
	red         = "\033[0;31m"
	lightRed    = "\033[1;31m"
	green       = "\033[0;32m"
	lightGreen  = "\033[1;32m"
	brown       = "\033[0;33m"
	yellow      = "\033[1;33m"
	blue        = "\033[0;34m"
	lightBlue   = "\033[1;34m"
	purple      = "\033[0;35m"
	lightPurple = "\033[1;35m"
	cyan        = "\033[0;36m"
	lightCyan   = "\033[1;36m"
	lightGray   = "\033[0;37m"
	white       = "\033[1;37m"
)

func Black(value string) string {
	return black + value + reset
}

func DarkGray(value string) string {
	return darkGray + value + reset
}

func Red(value string) string {
	return red + value + reset
}

func LightRed(value string) string {
	return lightRed + value + reset
}

func Green(value string) string {
	return green + value + reset
}

func LightGreen(value string) string {
	return lightGreen + value + reset
}

func Brown(value string) string {
	return brown + value + reset
}

func Yellow(value string) string {
	return yellow + value + reset
}

func Blue(value string) string {
	return blue + value + reset
}

func LightBlue(value string) string {
	return lightBlue + value + reset
}

func Purple(value string) string {
	return purple + value + reset
}

func LightPurple(value string) string {
	return lightPurple + value + reset
}

func Cyan(value string) string {
	return cyan + value + reset
}

func LightCyan(value string) string {
	return lightCyan + value + reset
}

func LightGray(value string) string {
	return lightGray + value + reset
}

func White(value string) string {
	return white + value + reset
}
