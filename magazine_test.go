package magazine

import (
	"testing"
)

func TestLoad(t *testing.T) {
	var (
		magazine, _     = Load("config.yml")
		expectedBool    = bool(true)
		expectedFloat64 = float64(1.0)
		expectedInt     = int(1)
		expectedString  = string("string")
	)
	assertEqual(t, expectedBool, magazine.GetBool("types.bool"))
	assertEqual(t, expectedFloat64, magazine.GetFloat64("types.float64"))
	assertEqual(t, expectedInt, magazine.GetInt("types.int"))
	assertEqual(t, expectedString, magazine.GetString("types.string"))
}
