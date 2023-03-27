package util_test

import (
	"testing"

	"docker-color-output/internal/util"
	"docker-color-output/pkg/color"
	"docker-color-output/pkg/util/assert"
)

func TestSplit(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, util.Split(tt.in))
		})
	}
}

func TestPad(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, util.Pad(tt.value, tt.length))
		})
	}
}

func TestIntersect(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, util.Intersect(tt.needle, tt.haystack))
		})
	}
}
