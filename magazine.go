package magazine

import "io/ioutil"

// Load returns a flattened map[string]interface{} representing contents of the file located at `filename`.
// Environment variables can be used to override key values.
func Load(filename string) (map[string]interface{}, error) {
	config := make(map[string]interface{})
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	mapYaml(content, config)
	flattenMap(config)
	applyEnv(config)
	return config, nil
}
