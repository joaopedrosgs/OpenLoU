package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type ConstructionsMap map[uint]ConstructionType

func (registeredConstructions ConstructionsMap) LoadAllConstructions() {

	println("Loading constructions")
	registeredConstructions = make(map[uint]ConstructionType)
	dir := "modules/constructions/"

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
		err = json.Unmarshal(fileContent, &element)
		if err == nil {
			registeredConstructions[uint(element.Id)] = element
			println("Construction:", element.Name, " Id: ", element.Id, " - Successfull")
		} else {
			println("File Name: ", module.Name(), " - Error")
		}
	}
	println("Loading constructions ended")
}
