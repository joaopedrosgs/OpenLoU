package constructions

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
type construction struct {
	Name         string
	Image        string
	Bonus        []bonus
	ResourceCost [][2]uint
	Adjascent    []adjascent
	Score        []uint
	Shared       string
}

var constructions []construction

func RegisterConstructions() {
	defer println("Construções carregadas!")

	dir := "constructions/modules/"
	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print("Erro ao ler o diretorio:", dir, " - ", err)
	}
	for _, module := range modules {
		var element construction
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			println("Erro ao ler o arquivo", dir, module.Name(), "-", err)
		}

		json.Unmarshal([]byte(string(fileContent)), &element)
		constructions = append(constructions, element)
		fmt.Printf("Construção carregada: %s\n", strings.Title(element.Name))
	}

}
