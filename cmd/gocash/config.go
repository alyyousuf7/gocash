package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/alyyousuf7/gocash"
	"gopkg.in/yaml.v2"
)

func loadConfiguration(configFile string) (*gocash.Configuration, error) {
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		if err := initConfiguration(configFile); err != nil {
			return nil, err
		}

		return loadConfiguration(configFile)
	}

	config := &gocash.Configuration{}
	return config, config.UnmarshalConfiguration(buf)
}

func initConfiguration(configFile string) error {
	configPath := path.Dir(configFile)
	dbFile := path.Join(configPath, "gocash.sqlite")

	config := &gocash.Configuration{
		Storage: "sqlite",
		StorageConfig: struct {
			File string `yaml:"file"`
		}{dbFile},
	}

	buf, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFile, buf, 0644)
}
