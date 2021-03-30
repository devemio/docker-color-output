package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Split(line string) []string {
	return regexp.MustCompile("\\s{2,}").Split(line, -1)
}

func GetMaxLens(lines []string) []int {
	lens := make([]int, len(Split(lines[0])))
	for _, line := range lines {
		for i, v := range Split(line) {
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
