package magazine

import (
	"os"
	"testing"
)

func TestFlattenMap(t *testing.T) {
	var (
		nestedMap = map[string]interface{}{
			"title":  "Good Reads",
			"volume": 2,
			"price":  4.99,
			"author": map[string]interface{}{
				"name":    "John Doe",
				"address": "123 Fake Street",
			},
		}
		expectedMap = map[string]interface{}{
			"title":          "Good Reads",
			"volume":         2,
			"price":          4.99,
			"author.name":    "John Doe",
			"author.address": "123 Fake Street",
		}
	)

	actualMap := flatten(nestedMap)

	assertEqual(t, expectedMap, actualMap)
}

func TestNestMap(t *testing.T) {
	var (
		flattenedMap = map[string]interface{}{
			"title":          "Good Reads",
			"volume":         2,
			"price":          4.99,
			"author.name":    "John Doe",
			"author.address": "123 Fake Street",
		}
		expectedMap = map[string]interface{}{
			"title":  "Good Reads",
			"volume": 2,
			"price":  4.99,
			"author": map[string]interface{}{
				"name":    "John Doe",
				"address": "123 Fake Street",
			},
		}
	)

	actualMap := expand(flattenedMap)

	assertEqual(t, expectedMap, actualMap)
}

func TestApplyEnv(t *testing.T) {
	var (
		actualMap = map[string]interface{}{
			"title":          "",
			"volume":         0,
			"price":          0.00,
			"author.name":    "",
			"author.address": "",
			"rare":           false,
		}
		expectedMap = map[string]interface{}{
			"title":          "Better Reads",
			"volume":         int64(3),
			"price":          float64(5.99),
			"author.name":    "Jane Doe",
			"author.address": "123 Faker Street",
			"rare":           true,
		}
	)
	os.Setenv("TITLE", "Better Reads")
	os.Setenv("VOLUME", "3")
	os.Setenv("PRICE", "5.99")
	os.Setenv("AUTHOR_NAME", "Jane Doe")
	os.Setenv("AUTHOR_ADDRESS", "123 Faker Street")
	os.Setenv("RARE", "true")

	applyEnv(actualMap)

	assertEqual(t, expectedMap, actualMap)
}
