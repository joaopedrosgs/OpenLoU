package configuration

import (
	"encoding/json"
	"github.com/jackc/pgx"
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
			KeySize       uint `json:"keySize"`
			KeyStringSize uint `json:"keyStringSize"`
		} `json:"security"`
		General struct {
			WorldSize       uint   `json:"worldSize"`
			OnlyCastle      string `json:"onlyCastle"`
			NoMoral         string `json:"noMoral"`
			ContinentSize   uint   `json:"continentSize"`
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

func Load() {
	arquivo, err := ioutil.ReadFile("configuration/default.json")
	if err != nil {
		context.WithField("Error", err.Error()).Info("The default configuration couldn't be loaded")
		return
	}
	err = json.Unmarshal(arquivo, &configuration)
	if err != nil {
		context.WithField("Error", err.Error()).Info("The default configuration couldn't be parsed")
		return
	}

	context.Info("Configuration loaded")

}

func GetConnConfig() pgx.ConnConfig {
	if configuration == nil {
		Load()
	}
	config := pgx.ConnConfig{
		Host:     configuration.Db.Host,
		Port:     uint16(configuration.Db.Port),
		Database: configuration.Db.Name,
		User:     configuration.Db.User,
		Password: configuration.Db.Password}
	return config
}

func GetSingleton() *Config {
	if configuration == nil {
		Load()
	}
	return configuration

}
