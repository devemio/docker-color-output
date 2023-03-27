package number

import (
	"strconv"
	"strings"
	"unicode"
)

func ParseFloat(value string) float64 {
	res, _ := strconv.ParseFloat(strings.TrimFunc(value, func(r rune) bool {
		return !unicode.IsNumber(r)
	}), 64)

	return res
}
