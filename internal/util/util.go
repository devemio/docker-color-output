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

func Split(in string) []string {
	return regexp.MustCompile(`\s{2,}`).Split(in, -1)
}

func Pad(value string, length int) string {
	length += NonPrintableCharactersLength * strings.Count(value, "\033[0m")

	return fmt.Sprintf("%-"+strconv.Itoa(length+SpaceLength)+"s", value)
}

func Intersect(needle, haystack []string) bool {
	m := make(map[string]struct{}, len(haystack))
	for _, v := range haystack {
		m[v] = struct{}{}
	}

	for _, v := range needle {
		if _, found := m[v]; !found {
			return false
		}
	}

	return true
}
