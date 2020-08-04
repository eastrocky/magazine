package magazine

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Magazine loads provides access to configurations
type Magazine struct {
	config map[string]interface{}
}

// GetBool returns the bool value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetBool(key string) bool {
	var defaultBool bool
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case bool:
			defaultBool = v.(bool)
		}
	}
	return defaultBool
}

// GetFloat64 returns the float64 value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetFloat64(key string) float64 {
	var defaultFloat64 float64
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case float64:
			defaultFloat64 = v.(float64)
		}
	}
	return defaultFloat64
}

// GetInt returns the int value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetInt(key string) int {
	var value int
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case int:
			value = v.(int)
		}
	}
	return value
}

// GetString returns the string value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetString(key string) string {
	var value string
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case string:
			value = v.(string)
		}
	}
	return value
}

// Load returns a flattened map[string]interface{} representing contents of the file located at `filename`.
// Environment variables can be used to override key values.
func Load(filename string) (*Magazine, error) {
	config := make(map[string]interface{})
	magazine := &Magazine{
		config: config,
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return magazine, err
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		return magazine, err
	}

	config = flattenMap(config)
	if err := applyEnv(config); err != nil {
		return magazine, err
	}

	return magazine, nil
}
