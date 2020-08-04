package magazine

import (
	"testing"
)

func TestLoad(t *testing.T) {
	var (
		magazine, _     = Load("config.yml")
		expectedBool    = true
		expectedFloat64 = 1.0
		expectedInt     = 1
		expectedString  = "string"
	)
	assertEqual(t, expectedBool, magazine.GetBool("bool"))
	assertEqual(t, expectedFloat64, magazine.GetFloat64("float64"))
	assertEqual(t, expectedInt, magazine.GetInt("int"))
	assertEqual(t, expectedString, magazine.GetString("string"))
}
