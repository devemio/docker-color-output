package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal:\nexpected: %s\nactual  : %s", expected, actual)
	}
}
