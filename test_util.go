package magazine

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
