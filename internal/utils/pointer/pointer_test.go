package pointer

import (
	"testing"

	"docker-color-output/pkg/utils/assert"
)

func TestToInt(t *testing.T) {
	v := 0

	tests := map[string]struct {
		in   int
		want *int
	}{
		"success": {v, &v},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToInt(tt.in))
		})
	}
}
