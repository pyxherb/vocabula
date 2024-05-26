package common

import (
	"encoding/json"
	"errors"
	"os"
)

type DatabaseConfig struct {
	Uri          string `json:"uri"`
	DatabaseName string `json:"databaseName"`
}

type ServerConfig struct {
	Database    DatabaseConfig `json:"database"`
	ServicePort uint16         `json:"servicePort"`
}

var GlobalServerConfig ServerConfig

func LoadServerConfig() error {
	content, err := os.ReadFile("config.json")

	if err != nil {
		return errors.New("error reading server configuration")
	}

	err = json.Unmarshal(content, &GlobalServerConfig)

	if err != nil {
		return errors.New("error deserializing server configuration")
	}

	return nil
}
