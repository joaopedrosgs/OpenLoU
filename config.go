package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	Connection string `json:"connection"`
	DbName     string `json:"dbname"`
	Parameters struct {
		Speed struct {
			Resource     string `json:"resource"`
			Military     string `json:"military"`
			Construction string `json:"construction"`
			CaveSpawn    string `json:"caveSpawn"`
		} `json:"speed"`
		General struct {
			WorldSize       uint   `json:"worldSize"`
			OnlyCastle      string `json:"onlyCastle"`
			NoMoral         string `json:"noMoral"`
			ContinentSize   string `json:"continentSize"`
			NightProtection struct {
				Activate   string `json:"activate"`
				Start      string `json:"start"`
				End        string `json:"end"`
				Percentage string `json:"percentage"`
			} `json:"nightProtection"`
			Limits struct {
				Alliance      string `json:"alliance"`
				Cities        string `json:"cities"`
				Constructions string `json:"constructions"`
			} `json:"limits"`
			Starter struct {
				Resources []int `json:"resources"`
			} `json:"starter"`
		} `json:"general"`
	} `json:"parameters"`
}

func (instance *Config) Load(fileName string) {
	log.Info("Loading configuration")
	arquivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Warn("Configuration file couldn't be loaded, using default settings")
		arquivo, err = ioutil.ReadFile("default.json")
	}
	err = json.Unmarshal(arquivo, &instance)
	if err != nil {
		log.Fatal("The configuration file couldn't be loaded")
	} else {
		log.Info("Configuration loaded")
	}

}
