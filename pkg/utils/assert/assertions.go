package assert

import "testing"

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Not equal:\nexpected: %s\nactual  : %s", expected, actual)
	}
}
