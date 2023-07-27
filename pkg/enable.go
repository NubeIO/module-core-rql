package pkg

import (
	"github.com/NubeIO/module-core-rql/apirules"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	log "github.com/sirupsen/logrus"
)

func (inst *Module) Enable() error {
	log.Infof("plugin is enabling...%s", name)
	eng := rules.NewRuleEngine()
	n := "Core"
	props := make(rules.PropertiesMap)
	props[n] = eng
	client := "RQL"
	newStorage, err := storage.New("data")
	if err != nil {
		log.Errorf("%s: error in making DB: %s", name, err)
		inst.ErrorOnDB = true
	}
	newClient := &apirules.RQL{
		Storage: newStorage,
	}
	props[client] = newClient
	inst.Rules = eng
	inst.Client = newClient
	inst.Props = props
	inst.Storage = newStorage
	log.Infof("plugin is enabled...%s", name)
	go inst.Loop()
	return nil
}

func (inst *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}
