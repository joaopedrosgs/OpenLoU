package gameEntities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type adjascent struct {
	Builds   []string
	Resource string
	Bonus    []uint
}
type bonus struct {
	Name  string
	Value []uint
}

type ConstructionType struct {
	Id           uint8
	Name         string
	Image        string
	Bonus        []bonus
	ResourceCost [][2]uint
	Adjascent    []adjascent
	Score        []uint
	Shared       string
}

var registeredConstructions map[uint]ConstructionType

func LoadAllConstructions() {
	defer println("Construções carregadas!")
	println("-- Carregando Construções --")
	registeredConstructions = make(map[uint]ConstructionType)
	dir := "constructions/modules/"

	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print("Erro ao ler o diretorio:", dir, " - ", err)
	}
	for _, module := range modules {
		var element ConstructionType
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			println("Erro ao ler o arquivo", dir, module.Name(), "-", err)
		}

		json.Unmarshal(fileContent, &element)

		registeredConstructions[uint(element.Id)] = element
		fmt.Printf("Construção carregada: %.2d => %s\n", element.Id, strings.Title(element.Name))
	}

}
