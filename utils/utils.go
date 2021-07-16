package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type SplitContract func(line string) []string

func Split(line string) []string {
	cols := regexp.MustCompile(`\s{2,}`).Split(line, -1)
	if cols[0] == "" {
		return cols[1:]
	}
	return cols
}

func GetMaxLens(lines []string, split SplitContract) []int {
	lens := make([]int, len(split(lines[0])))
	for _, line := range lines {
		for i, v := range split(line) {
			length := len(v)
			if strings.Contains(v, "…") {
				length -= 2
			}
			if length > lens[i] {
				lens[i] = length
			}
		}
	}
	return lens
}

func ParseFloat(value string) float64 {
	res, _ := strconv.ParseFloat(strings.TrimFunc(value, func(r rune) bool {
		return !unicode.IsNumber(r)
	}), 64)
	return res
}
