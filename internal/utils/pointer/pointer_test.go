package pointer

import (
	"testing"

	"docker-color-output/pkg/utils/assert"
)

func TestToInt(t *testing.T) {
	v := 0

	assert.Equal(t, &v, ToInt(v))
}
