package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	SpaceLength                  = 3
	NonPrintableCharactersLength = 11
)

func Split(v string) []string {
	return regexp.MustCompile(`\s{2,}`).Split(v, -1)
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

func Debug(value any) { // @fixme
	fmt.Printf("\n---\n%#v\n---\n\n", value)
	os.Exit(0)
}
