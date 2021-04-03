package stdout

import (
	"fmt"
	"strconv"
	"strings"
)

func Format(value string, length int) string {
	length += NonPrintableCharactersLength * strings.Count(value, "\033[0m")
	return fmt.Sprintf("%-"+strconv.Itoa(length+SpaceLength)+"s", value)
}
