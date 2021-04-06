package main

import (
	"flag"
	"io/ioutil"
)

const defaultConfigPath = "./configs"

const defaultConfigName = "/default_config.yaml"

var configName string

func init() {
	flag.StringVar(&configName, "config", defaultConfigName + defaultConfigName, "config name")
}

func main() {
	flag.Parse()
}

func getConfig(path string) (*configs.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
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