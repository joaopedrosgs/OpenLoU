package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
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
			WorldSize       uint   `json:worldSize`
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

var Parameters config

func Load() {

	arquivo, err := ioutil.ReadFile("settings.json")
	if err != nil {
		println("Erro ao carregar as configurações: " + err.Error())
	}
	json.Unmarshal(arquivo, &Parameters)
	println("Configurações carregadas!")

}
