package data

import (
	"archeboot/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

const DEFAULT_CONFIG_FILE_NAME = "archeboot.json"

type Config struct {
	ConfigPath      string
	RepoPath        string
	DefaultFilePath string
	Info            *ConfigInfoJSON
}

type ConfigInfoJSON struct {
	GitCommand          string `json:"git_command"`
	DefaultBootFileName string `json:"default_boot_file_name"`
}

func InitDefaultFile(conf *Config) (*ConfigInfoJSON, error) {
	confInfo := &ConfigInfoJSON{
		GitCommand:          "git",
		DefaultBootFileName: "arche.yaml",
	}

	exists, err := utils.PathExists(conf.DefaultFilePath)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}

	var bs []byte
	var err2 error
	var err3 error

	if exists {
		bs, err2 = ioutil.ReadFile(conf.DefaultFilePath)
		if err2 != nil {
			utils.LogError(err2)
			return nil, err2
		}
		newInfo := &ConfigInfoJSON{}
		json.Unmarshal(bs, newInfo)
		if newInfo.GitCommand != "" {
			confInfo.GitCommand = newInfo.GitCommand
		}
		if newInfo.DefaultBootFileName != "" {
			confInfo.DefaultBootFileName = newInfo.DefaultBootFileName
		}
	} else {
		bs, err3 = json.MarshalIndent(confInfo, "", "    ")
		if err2 != nil {
			utils.LogError(err3)
			return nil, err3
		}
		ioutil.WriteFile(conf.DefaultFilePath, bs, os.ModeAppend)
	}
	fmt.Println("Config Info:", *confInfo)
	return confInfo, nil
}

func NewConfig() *Config {
	conf := &Config{}
	u, _ := user.Current()
	homeDir := u.HomeDir
	homeDir = strings.Replace(homeDir, "\\", "/", -1)
	conf.ConfigPath = homeDir + "/.archeboot"
	conf.RepoPath = conf.ConfigPath + "/repo"
	conf.DefaultFilePath = conf.ConfigPath + "/" + DEFAULT_CONFIG_FILE_NAME
	utils.Mkdir(conf.ConfigPath)
	utils.Mkdir(conf.RepoPath)
	conf.Info, _ = InitDefaultFile(conf)
	return conf
}
