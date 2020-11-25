package magazine

import (
	"io/ioutil"

	"github.com/eastrocky/magazine/bellows"
	"gopkg.in/yaml.v3"
)

// Magazine loads provides access to configurations
type Magazine struct {
	config map[string]interface{}
}

// Load returns a flattened map[string]interface{} representing contents of the file located at `filepath`.
// Environment variables can be used to override key values.
func Load(filepath string, i ...interface{}) (*Magazine, error) {
	config := make(map[string]interface{})
	magazine := &Magazine{}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return magazine, err
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		return magazine, err
	}

	config = flatten(config)
	if err := applyEnv(config); err != nil {
		return magazine, err
	}

	magazine.config = config

	if len(i) > 0 {
		bytes, err := yaml.Marshal(bellows.Expand(config))
		if err != nil {
			return magazine, err
		}

		return magazine, yaml.Unmarshal(bytes, i[0])
	}

	return magazine, nil
}

// Eject writes configurations loaded by magazine to file.
func Eject(filepath string, config interface{}) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, bytes, 0644)
}

// GetBool returns the bool value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetBool(key string) bool {
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case bool:
			return v.(bool)
		}
	}
	var defaultBool bool
	return defaultBool
}

// GetFloat64 returns the float64 value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetFloat64(key string) float64 {
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case float64:
			return v.(float64)
		}
	}
	var defaultFloat64 float64
	return defaultFloat64
}

// GetInt returns the int value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetInt(key string) int {
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case int:
			return v.(int)
		}
	}
	var defaultInt int
	return defaultInt
}

// GetString returns the string value at key.
// Returns the zero value when key is not set.
func (m *Magazine) GetString(key string) string {
	v, set := m.config[key]
	if set {
		switch v.(type) {
		case string:
			return v.(string)
		}
	}
	var defaultString string
	return defaultString
}
