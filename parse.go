package magazine

import (
	"os"
	"strconv"
	"strings"
)

func flatten(nested map[string]interface{}) (flattened map[string]interface{}) {
	m := make(map[string]interface{})
	for k, v := range nested {
		assign(m, k, v)
	}
	return m
}

func assign(m map[string]interface{}, parentKey string, parentValue interface{}) {
	switch parentValue.(type) {
	case map[string]interface{}:
		for childKey, childValue := range parentValue.(map[string]interface{}) {
			assign(m, combineKeys(parentKey, childKey), childValue)
		}
	default:
		m[parentKey] = parentValue
	}
}

func combineKeys(parentKey string, childKey string) string {
	return parentKey + "." + childKey
}

func applyEnv(m map[string]interface{}) error {
	for k, v := range m {
		envKev := strings.ReplaceAll(strings.ToUpper(k), ".", "_")
		envValue, set := os.LookupEnv(envKev)
		if set {
			switch v.(type) {
			case bool:
				if err := mapBool(m, k, envValue); err != nil {
					return err
				}
			case float64:
				if err := mapFloat64(m, k, envValue); err != nil {
					return err
				}
			case int:
				if err := mapInt64(m, k, envValue); err != nil {
					return err
				}
			case string:
				m[k] = envValue
			}
		}
	}
	return nil
}

func mapBool(m map[string]interface{}, k string, v string) error {
	converted, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	m[k] = converted
	return nil
}

func mapFloat64(m map[string]interface{}, k string, v string) error {
	converted, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	m[k] = converted
	return nil
}

func mapInt64(m map[string]interface{}, k string, v string) error {
	converted, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}
	m[k] = converted
	return nil
}
