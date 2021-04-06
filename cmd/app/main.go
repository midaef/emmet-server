package main

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"

	"github.com/midaef/emmet-server/configs"
	"github.com/midaef/emmet-server/internal/app"
)

const defaultConfigPath = "./configs/"

const defaultConfigName = "default_config.yaml"

var configName string

var configPath string

func init() {
	flag.StringVar(&configName, "config-name", defaultConfigName, "config name")
	flag.StringVar(&configPath, "config-path", defaultConfigPath, "config path")
}

func main() {
	flag.Parse()

	config, err := getConfig(configName)
	if err != nil {
		log.Printf("package main: config error \n%v", err)
	}

	app.Run(config)
}

func getConfig(name string) (*configs.Config, error) {
	configPath = configPath + name

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config *configs.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}