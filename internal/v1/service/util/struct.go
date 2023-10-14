package util

import (
	"encoding/json"
	"fmt"
)

func StructToStringMap(s interface{}) (map[string]string, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var intermediateMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &intermediateMap)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for key, value := range intermediateMap {
		result[key] = fmt.Sprintf("%v", value)
	}

	return result, nil
}
