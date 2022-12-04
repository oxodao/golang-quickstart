package config

import (
	"io/ioutil"
	"os"

	"github.com/oxodao/go-quickstart/utils"
	"gopkg.in/yaml.v2"
)

var GET Config

type Config struct {
	Web struct {
		ListeningAddr string `yaml:"listening_addr"`
	} `yaml:"web"`

	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`
}

func Load() error {
	cfg := Config{}

	configPath := os.Getenv(utils.APP_NAME + "CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "./config.yaml"
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	GET = cfg

	return nil
}
