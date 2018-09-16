package modules

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var RegisteredModules map[string]map[int]Module
var logger = log.WithField("Entity", "Modules")

func RegisterAllModules() {
	RegisterAllConstructions()
	RegisterAllTroops()
}
func RegisterAllConstructions() {
	dir := "modules/constructions/"
	logger.WithField("From", dir).Info("Loading constructions")
	if RegisteredModules == nil {
		RegisteredModules = make(map[string]map[int]Module)

	}
	if RegisteredModules["construction"] == nil {
		RegisteredModules["construction"] = make(map[int]Module)
	}
	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read modules directory: " + dir)
		return

	}
	for _, module := range modules {
		element := &ConstructionType{}
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			logger.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read construction module file")
			continue
		}
		err = json.Unmarshal(fileContent, &element)
		if err != nil {
			logger.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Construction type could not be loaded")
			continue
		}
		RegisteredModules["construction"][element.ID] = element
		logger.WithFields(log.Fields{"Module": int(element.ID), "Name": element.Name, "Status": "Successful"}).Info("Construction type Loaded")
	}
	logger.Info("All constructions loaded!")

}

func RegisterAllTroops() {
	dir := "modules/troops/"
	logger.WithField("From", dir).Info("Loading troops")
	if RegisteredModules == nil {
		RegisteredModules = make(map[string]map[int]Module)

	}
	if RegisteredModules["troop"] == nil {
		RegisteredModules["troop"] = make(map[int]Module)
	}
	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read modules directory: " + dir)
		return

	}
	for _, module := range modules {
		element := &TroopType{}
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			logger.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read troop module file")
			continue
		}
		err = json.Unmarshal(fileContent, &element)
		if err == nil {
			RegisteredModules["troop"][element.ID] = element
			logger.WithFields(log.Fields{"Module": int(element.ID), "Name": element.Name, "Status": "Successful"}).Info("Construction type Loaded")
		} else {
			logger.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Construction type could not be loaded")

		}
	}
	logger.Info("All troops loaded!")
}
