package lib

import (
	"encoding/json"
	"os"
	"io"
)

type ConfigReader struct {
	ConfigPath string
}

func NewConfigReader(configPath string) *ConfigReader {
	return &ConfigReader{ConfigPath: configPath}
} 

func (c *ConfigReader) ReadConfig() ([]string, error) {
	file, err := os.Open(c.ConfigPath)
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