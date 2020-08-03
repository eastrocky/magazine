package magazine

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Load returns a flattened map[string]interface{} representing contents of the file located at `filename`.
// Environment variables can be used to override key values.
func Load(filename string) (map[string]interface{}, error) {
	config := make(map[string]interface{})
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		return config, err
	}

	config = flattenMap(config)
	if err := applyEnv(config); err != nil {
		return config, err
	}

	return config, nil
}
