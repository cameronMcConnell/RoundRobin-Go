package lib

import (
	"encoding/json"
	"os"
	"io"
)

func readConfig() ([]string, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var servers []string
	err = json.Unmarshal(bytes, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}