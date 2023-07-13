package pkg

import (
	"context"
	"fmt"
	"github.com/NubeIO/module-core-rql/apirules"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m *Module) Enable() error {
	log.Infof("plugin is enabling...%s", name)
	eng := rules.NewRuleEngine()
	n := "Core"
	props := make(rules.PropertiesMap)
	props[n] = eng
	client := "RQL"
	newStorage := storage.New(fmt.Sprintf("%s.db", name))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newClient := &apirules.Client{
		CTX:     ctx,
		Storage: newStorage,
		PdfApplication: &apirules.PDFApplication{
			PandocPath:     "/usr/share/pandoc",
			UserHome:       "/home/aidan",
			PandocDataDir:  "/.pandoc",
			CommandTimeout: 10 * time.Second,
		},
	}
	props[client] = newClient
	m.Rules = eng
	m.Client = newClient
	m.Props = props
	m.Storage = newStorage
	log.Infof("plugin is enabled...%s", name)
	return nil
}

func (m *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}
