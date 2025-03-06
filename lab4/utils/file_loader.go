package utils

import (
	"encoding/json"
	"errors"
	"os"
)

// LoadCableData loads the cable_data.json file and returns parsed data
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