package entities

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

var RegisteredConstructions map[int]ConstructionType
var RegisteredTroops map[int]TroopType
var context = log.WithField("Entity", "OpenLoU")

func RegisterAllConstructions() {
	dir := "modules/constructions/"
	context.WithField("From", dir).Info("Loading constructions")
	RegisteredConstructions = make(map[int]ConstructionType)

	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		context.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read constructions directory")
	}
	for _, module := range modules {
		var element ConstructionType
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			context.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read construction module file")
			continue
		}
		err = json.Unmarshal(fileContent, &element)
		if err == nil {
			RegisteredConstructions[int(element.ID)] = element
			context.WithFields(log.Fields{"Module": element.Name, "Status": "Successful"}).Info("Construction type Loaded")
		} else {
			context.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Construction type could not be loaded")

		}
	}
	context.Info("All constructions loaded!")

}

func RegisterAllTroops() {

	println("Loading troops")
	RegisteredTroops = make(map[int]TroopType)
	dir := "modules/military/"

	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		context.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read troops directory")
	}
	for _, module := range modules {
		var element TroopType
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			context.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read troop module file")
			continue
		}

		err = json.Unmarshal(fileContent, &element)
		if err == nil {
			RegisteredTroops[element.Id] = element
			context.WithFields(log.Fields{"Module": element.Name, "Status": "Successful"}).Info("Troop type Loaded")
		} else {
			context.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Troop type could not be loaded")
		}
	}
	context.Info("All troops loaded!")
}
