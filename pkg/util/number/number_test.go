package number_test

import (
	"testing"

	"docker-color-output/pkg/util/assert"
	"docker-color-output/pkg/util/number"
)

func TestParseFloat(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, number.ParseFloat(tt.in))
		})
	}
}
