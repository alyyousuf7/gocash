package gocash

import "gopkg.in/yaml.v2"

type Configuration struct {
	Storage       string      `yaml:"storage"`
	StorageConfig interface{} `yaml:"storage-config"`
}

func (c *Configuration) UnmarshalConfiguration(buf []byte) error {
	if err := yaml.Unmarshal(buf, c); err != nil {
		return err
	}

	return nil
}
