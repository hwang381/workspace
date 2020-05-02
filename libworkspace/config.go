package libworkspace

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/atrox/homedir"
)

func ReadConfig() (*Config, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	configPath := path.Join(homeDir, ".workspace", "config.json")
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
