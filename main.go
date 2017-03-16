package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type adjascent struct {
	Builds   []string `json:"builds"`
	Resource string   `json:"resource"`
	Bonus    []uint   `json:"bonus"`
}
type bonus struct {
	Name  string `json:"name"`
	Value []uint `json:"value"`
}
type construction struct {
	Name         string
	Bonus        []bonus
	ResourceCost [][2]uint
	Adjascent    []adjascent `json:"adjascent"`
	Score        []uint
	Shared       string
}

func main() {
	jsons, err := ioutil.ReadFile("constructions.json")
	if err != nil {
		fmt.Println(err)
	}
	var constructions []construction
	errs := json.Unmarshal(jsons, &constructions)
	if errs != nil {
		fmt.Println("error:", errs)
	}
	fmt.Printf("%+v", constructions[0].Name)

}
