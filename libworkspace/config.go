package libworkspace

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/atrox/homedir"
)

func fileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

func migrateSingleConfig() error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return err
	}

	singleConfigPath := path.Join(homeDir, ".workspace", "config.json")
	exists, err := fileExists(singleConfigPath)
	if err != nil {
		return err
	}
	if exists {
		glog.Infoln("Migrating single config at %v for multiple workspaces", singleConfigPath)
		newConfigPath := path.Join(homeDir, ".workspace", "default.workspace.json")
		err := os.Rename(singleConfigPath, newConfigPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadConfigs() (map[string]*Config, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	err = migrateSingleConfig()
	if err != nil {
		return nil, err
	}

	configs := make(map[string]*Config)
	err = filepath.Walk(path.Join(homeDir, ".workspace"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".workspace.json") {
			configName := strings.TrimSuffix(info.Name(), ".workspace.json")
			configBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			config := Config{}
			err = json.Unmarshal(configBytes, &config)
			if err != nil {
				return err
			}
			configs[configName] = &config
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func ReadConfig(workspaceName string) (*Config, error) {
	configs, err := ReadConfigs()
	if err != nil {
		return nil, err
	}

	config, ok := configs[workspaceName]
	if !ok {
		return nil, fmt.Errorf("Workspace %v cannot be found\n", workspaceName)
	}
	return config, nil
}
