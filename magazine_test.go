package magazine

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	var (
		magazine, _     = Load("testdata/config.yml")
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

func TestLoadWithInterface(t *testing.T) {
	type Config struct {
		Types struct {
			Bool    bool
			Float64 float64
			Int     int
			String  string
		}
	}
	var (
		configInterface = &Config{}
		expectedBool    = bool(true)
		expectedFloat64 = float64(1.0)
		expectedInt     = int(1)
		expectedString  = string("string")
	)

	Load("testdata/config.yml", &configInterface)

	assertEqual(t, expectedBool, configInterface.Types.Bool)
	assertEqual(t, expectedFloat64, configInterface.Types.Float64)
	assertEqual(t, expectedInt, configInterface.Types.Int)
	assertEqual(t, expectedString, configInterface.Types.String)
}

func TestLoadWithInterfaceAndEnvironment(t *testing.T) {
	type Config struct {
		Types struct {
			Bool    bool
			Float64 float64
			Int     int
			String  string
		}
	}
	var (
		configInterface = &Config{}
		expectedBool    = bool(true)
		expectedFloat64 = float64(2.0)
		expectedInt     = int(2)
		expectedString  = string("strings")
	)

	os.Setenv("TYPES_BOOL", "true")
	os.Setenv("TYPES_FLOAT64", "2.0")
	os.Setenv("TYPES_INT", "2")
	os.Setenv("TYPES_STRING", "strings")

	Load("testdata/config.yml", &configInterface)

	assertEqual(t, expectedBool, configInterface.Types.Bool)
	assertEqual(t, expectedFloat64, configInterface.Types.Float64)
	assertEqual(t, expectedInt, configInterface.Types.Int)
	assertEqual(t, expectedString, configInterface.Types.String)
}
