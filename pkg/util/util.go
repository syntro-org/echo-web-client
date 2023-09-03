package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

const DEBUG_MODE = "debug"
const RELEASE_MODE = "release"

type AppParams struct {
	BaseAPIUrl    string `yaml:"base_api_url"`
	CookieSecrets string `yaml:"cookie_secrets"`
	AppMode       string `yaml:"app_mode"`
	Port          string `yaml:"port"`
}

func ReadParams(path string, a interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), a)
	if err != nil {
		return err
	}

	return nil
}

func MakeServerAddr(mode, port string) string {
	if mode == RELEASE_MODE {
		return ":" + port
	}

	return "localhost:" + port
}
