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

func ReadParams(path string, a *interface{}) error {
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

func MakeServerAddr(configPath, mode, port string, p *interface{}) string {
	ReadParams(configPath, p)

	if mode == RELEASE_MODE {
		return ":" + port
	}

	return "localhost:" + port
}
