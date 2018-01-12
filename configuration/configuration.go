package configuration

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

var context = log.WithFields(log.Fields{"Entity": "Configuration"})

type Config struct {
	Db struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
		SSL      string `json:"SSL"`
	} `json:"db"`
	Parameters struct {
		Speed struct {
			Resource     string `json:"resource"`
			Military     string `json:"military"`
			Construction string `json:"construction"`
			CaveSpawn    string `json:"caveSpawn"`
		} `json:"speed"`
		Security struct {
			KeySize int `json:"keySize"`
		} `json:"security"`
		General struct {
			WorldSize       int    `json:"worldSize"`
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

var configuration *Config

func (instance *Config) Load() {
	arquivo, err := ioutil.ReadFile("default.json")
	err = json.Unmarshal(arquivo, &instance)
	if err != nil {
		context.Info("The default configuration couldn't be loaded")
	} else {
		context.Info("Configuration loaded")
	}

}

func GetConnectionString() string {
	GetInstance()
	connectionString := "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString, configuration.Db.User, configuration.Db.Password, configuration.Db.Host, configuration.Db.Port, configuration.Db.Name, configuration.Db.SSL)
	return connectionString
}

func GetInstance() *Config {
	if configuration == nil {
		configuration = &Config{}
		configuration.Load()
	}
	return configuration
}
