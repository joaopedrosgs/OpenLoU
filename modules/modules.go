package modules

import (
	"encoding/json"
	"io/ioutil"

	"github.com/joaopedrosgs/OpenLoU/models"
	log "github.com/sirupsen/logrus"
)

var RegisteredConstructions map[int]models.ConstructionType
var RegisteredTroops map[int]models.TroopType
var logger = log.WithField("Entity", "Modules")

func RegisterAllConstructions() {
	dir := "modules/constructions/"
	logger.WithField("From", dir).Info("Loading constructions")
	RegisteredConstructions = make(map[int]models.ConstructionType)

	if modules, err := ioutil.ReadDir(dir); err != nil {
		logger.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read constructions directory")

	} else {
		for _, module := range modules {
			var element models.ConstructionType
			fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
			if err != nil {
				logger.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read construction module file")
				continue
			}
			err = json.Unmarshal(fileContent, &element)
			if err == nil {
				RegisteredConstructions[int(element.ID)] = element
				logger.WithFields(log.Fields{"Module": int(element.ID), "Name": element.Name, "Status": "Successful"}).Info("Construction type Loaded")
			} else {
				logger.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Construction type could not be loaded")

			}
		}
		logger.Info("All constructions loaded!")
	}

}

func RegisterAllTroops() {
	RegisteredTroops = make(map[int]models.TroopType)
	dir := "modules/military/"
	logger.WithField("From", dir).Info("Loading troops")

	modules, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.WithFields(log.Fields{"Error": err.Error()}).Error("Failed to read troops directory")
	}
	for _, module := range modules {
		var element models.TroopType
		fileContent, err := ioutil.ReadFile(string(dir + module.Name()))
		if err != nil {
			logger.WithFields(log.Fields{"Error": err.Error(), "Module": module.Name()}).Error("Failed to read troop module file")
			continue
		}

		err = json.Unmarshal(fileContent, &element)
		if err == nil {
			RegisteredTroops[element.ID] = element
			logger.WithFields(log.Fields{"Module": element.Name, "Status": "Successful"}).Info("Troop type Loaded")
		} else {
			logger.WithFields(log.Fields{"Module": element.Name, "Error": err.Error()}).Info("Troop type could not be loaded")
		}
	}
	logger.Info("All troops loaded!")
}
