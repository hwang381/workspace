package main

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/atrox/homedir"
)

func readConfig() (*Config, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	configPath := path.Join(homeDir, ".workspace", "config.json")
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := Config{
		GitConfig: GitConfig{
			PrivateKeyPath: path.Join(homeDir, ".ssh", "id_rsa"),
		},
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
