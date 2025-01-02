package services_utils

import (
	"encoding/json"
)

func JsonUnmarshalInterface[T any](data_map map[string]interface{}, target *T) error {
	data_json, err := json.Marshal(data_map)
	if err != nil {
		return err
	}

	return json.Unmarshal(data_json, target)
}
