package numbers

import (
	"testing"

	"docker-color-output/pkg/utils/assert"
)

func TestParseFloat(t *testing.T) {
	tests := map[string]struct {
		in   string
		want float64
	}{
		"plain":       {in: "100", want: 100},
		"float":       {in: "100.10", want: 100.10},
		"measurement": {in: "100MB", want: 100},
		"trim":        {in: " 100 ", want: 100},
		"empty":       {in: "", want: 0},
		"unparsed":    {in: "-", want: 0},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ParseFloat(tt.in))
		})
	}
}
