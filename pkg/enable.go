package pkg

import (
	"github.com/NubeIO/module-core-rql/apirules"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	log "github.com/sirupsen/logrus"
)

func (m *Module) Enable() error {
	log.Infof("plugin is enabling...%s", name)
	eng := rules.NewRuleEngine()
	n := "Core"
	props := make(rules.PropertiesMap)
	props[n] = eng
	client := "RQL"
	newStorage, err := storage.New(m.moduleDirectory)
	if err != nil {
		log.Errorf("%s: error in making DB: %s", name, err)
		m.ErrorOnDB = true
	}
	newClient := &apirules.RQL{
		Storage: newStorage,
		Config:  m.GetConfig(),
		ROS:     m.grpcMarshaller,
	}

	props[client] = newClient
	m.Rules = eng
	m.Client = newClient
	m.Props = props
	m.Storage = newStorage
	log.Infof("plugin is enabled...%s", name)
	go m.Loop()
	m.pluginIsEnabled = true
	return nil
}

func (m *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}
