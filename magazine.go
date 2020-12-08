package magazine

import (
	"io/ioutil"

	"github.com/eastrocky/magazine/bellows"
	"gopkg.in/yaml.v3"
)

// Load binds previously ejected magazines back into structures.
// Environment variables can be used to override loaded key values.
func Load(filepath string, config interface{}) error {
	magazine := make(map[string]interface{})
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(fileBytes, &magazine); err != nil {
		return err
	}

	magazine = flatten(magazine)
	if err := applyEnv(magazine); err != nil {
		return err
	}

	magBytes, err := yaml.Marshal(bellows.Expand(magazine))
	if err != nil {
		return err
	}

	return yaml.Unmarshal(magBytes, config)
}

// Eject writes configurations loaded by magazine to file.
func Eject(filepath string, config interface{}) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, bytes, 0644)
}
