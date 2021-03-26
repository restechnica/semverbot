package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type readFile func(path string) ([]byte, error)
type unmarshal func(data []byte, config *Root) (err error)

type YAMLLoader struct {
	readFile
	unmarshal
}

func NewYAMLLoader() YAMLLoader {
	return YAMLLoader{
		readFile: ioutil.ReadFile,
		unmarshal: func(data []byte, config *Root) (err error) {
			return yaml.Unmarshal(data, config)
		},
	}
}

// Load loads a yaml file into a config root.
// It returns the newly loaded config root.
func (loader YAMLLoader) Load(path string) (config Root, err error) {
	return loader.Overload(path, config)
}

// Overload loads a yaml file into an existing config root, overwrites existing data.
// It returns the newly loaded config root.
func (loader YAMLLoader) Overload(path string, config Root) (Root, error) {
	var data []byte
	var err error

	if data, err = loader.readFile(path); err == nil {
		err = loader.unmarshal(data, &config)
	}

	return config, err
}
