package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type troopType struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CanAttack bool   `json:"canAttack"`
	Image     string `json:"image"`
	Attack    uint   `json:"attack"`
	Defense   uint   `json:"defense"`
	Loot      uint   `json:"loot"`
	Speed     uint   `json:"speed"`
	Requires  []struct {
		Name  string `json:"name"`
		Id    string `json:"id"`
		Value string `json:"value"`
	} `json:"requires"`
	Cost   [5]uint
	Upkeep [5]uint
}

var RegisteredTroops map[uint]troopType

func RegisterAll() {

	defer println("Tropas carregadas!")
	println("Carregando Tropas")
	RegisteredTroops = make(map[uint]troopType)
	dir := "military/modules/"

	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print("Erro ao ler o diretorio:", dir, " - ", err)
	}
	for _, module := range modules {
		var element troopType
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			println("Erro ao ler o arquivo", dir, module.Name(), "-", err)
		}

		json.Unmarshal(fileContent, &element)

		RegisteredTroops[element.Id] = element
		fmt.Printf("Tropa carregada: %s\n", strings.Title(element.Name))
	}
}

type MilitaryAction struct {
	Id     uint
	Type   uint8
	Troops []struct {
		ID    uint8
		Quant uint
	}
	Depart   time.Time
	Duration time.Duration
}
