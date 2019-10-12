package util

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// ConvertInterfaceToPostgresJSONB converts a field in an jsonlike interface form to postgres jsonb type
func ConvertInterfaceToPostgresJSONB(modelMap interface{}, key string) postgres.Jsonb {
	switch v := modelMap.(type) {
	default:
		fmt.Printf("unexpected type %T", v)
	case map[string]interface{}:
		modelMap = modelMap.(map[string]interface{})[key]
	case []interface{}:
		modelMap = modelMap.(map[string]interface{})[key]
	}

	if modelMap == nil {
		modelMap = json.RawMessage("{}")
	}

	jsonEnc, err := json.Marshal(modelMap)
	if err != nil {
		panic("Could not marshal json")
	}

	return postgres.Jsonb{json.RawMessage(jsonEnc)}
}
