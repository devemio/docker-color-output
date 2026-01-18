package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	SpaceLength                  = 3
	NonPrintableCharactersLength = 11
)

var splitRegexp = regexp.MustCompile(`\s{2,}`) //nolint:gochecknoglobals

func Split(in string) []string {
	return splitRegexp.Split(in, -1)
}

func Pad(value string, length int) string {
	length += NonPrintableCharactersLength * strings.Count(value, "\033[0m")

	return fmt.Sprintf("%-"+strconv.Itoa(length+SpaceLength)+"s", value)
}

func Intersect(needle, haystack []string) bool {
	x := make(map[string]struct{}, len(haystack))
	for _, v := range haystack {
		x[v] = struct{}{}
	}

	for _, v := range needle {
		if _, found := x[v]; !found {
			return false
		}
	}

	return true
}
