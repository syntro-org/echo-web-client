package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

const DEBUG_MODE = "debug"
const RELEASE_MODE = "release"

type AppParams struct {
	DBConnectionString string `yaml:"db_conn_string"`
	AppMode            string `yaml:"app_mode"`
	Port               string `yaml:"port"`
	SigningKey         string `yaml:"signing_key"`
}

func ReadParams(path string, a *AppParams) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), &a)
	if err != nil {
		return err
	}

	return nil
}

func MakeServerAddr(configPath string) string {
	var p AppParams
	ReadParams(configPath, &p)

	if p.AppMode == RELEASE_MODE {
		return ":" + p.Port
	}

	return "localhost:" + p.Port
}
