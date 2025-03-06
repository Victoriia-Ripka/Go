package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func LoadCableData(filePath string) ([]map[string]interface{}, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("the file cable_data.json was not found")
	}

	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, errors.New("failed to parse JSON data")
	}

	return data, nil
}