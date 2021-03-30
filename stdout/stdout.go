package stdout

import (
	"fmt"
	"strconv"
	"strings"
)

type LinePrinter interface {
	Println(lens []int)
}

func Format(value string, length int) string {
	length += NonPrintableCharactersLength * strings.Count(value, "\033[0m")
	return fmt.Sprintf("%-"+strconv.Itoa(length+SpaceLength)+"s", value)
}
