package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal:\nexpected: %s\nactual  : %s", expected, actual)
	}
}
