package military

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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
	Cost struct {
		Food  uint `json:"food"`
		Wood  uint `json:"wood"`
		Iron  uint `json:"iron"`
		Stone uint `json:"stone"`
		Gold  uint `json:"gold"`
	} `json:"cost"`
	Upkeep struct {
		Food  uint `json:"food"`
		Wood  uint `json:"wood"`
		Iron  uint `json:"iron"`
		Stone uint `json:"stone"`
		Gold  uint `json:"gold"`
	} `json:"upkeep"`
}

var registeredTroops map[uint]troopType

func RegisterAll() {

	defer println("Tropas carregadas!")
	println("Carregando Tropas")
	registeredTroops = make(map[uint]troopType)
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

		registeredTroops[element.Id] = element
		fmt.Printf("Tropa carregada: %s\n", strings.Title(element.Name))
	}
}
