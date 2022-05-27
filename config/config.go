package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	DbSettings    DbSettings    `json:"DbSettings"`
}

type DbSettings struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}

func ReadSettingsFromFile(settingFilePath string) (config *Config) {
	var Config Config
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		panic("No such file named " + settingFilePath)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Panic(err)
	}
	config = &Config
	return config
}
