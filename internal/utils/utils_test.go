package utils

import (
	"testing"

	"docker-color-output/pkg/color"
	"docker-color-output/pkg/utils/assert"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		in   string
		want []string
	}{
		"success": {
			in:   "CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES",
			want: []string{"CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS", "PORTS", "NAMES"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Split(tt.in))
		})
	}
}

func TestPad(t *testing.T) {
	tests := map[string]struct {
		value  string
		length int
		want   string
	}{
		"colored": {
			value:  color.Red("IMAGE"),
			length: 10,
			want:   "\u001B[0;31mIMAGE\u001B[0m        ",
		},
		"plain": {
			value:  "CONTAINER ID",
			length: 15,
			want:   "CONTAINER ID      ",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Pad(tt.value, tt.length))
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := map[string]struct {
		needle   []string
		haystack []string
		want     bool
	}{
		"empty":            {[]string{}, []string{}, true},
		"contains":         {[]string{"1", "10"}, []string{"100", "1", "10"}, true},
		"does_not_contain": {[]string{"1", "10"}, []string{"0", "1", "5"}, false},
		"empty_value":      {[]string{"", "1"}, []string{"1", "100", ""}, true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Intersect(tt.needle, tt.haystack))
		})
	}
}
