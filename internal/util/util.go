package util

import (
	"fmt"
	"regexp"
	"strconv"
)

const SpaceLength = 3

var splitRegexp = regexp.MustCompile(`\s{2,}`)
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func Split(in string) []string {
	return splitRegexp.Split(in, -1)
}

func Pad(value string, length int) string {
	visible := len(ansiRegexp.ReplaceAllString(value, ""))
	extra := max(len(value)-visible, 0)

	return fmt.Sprintf("%-"+strconv.Itoa(length+SpaceLength+extra)+"s", value)
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
