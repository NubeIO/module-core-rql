package pkg

import (
	log "github.com/sirupsen/logrus"
)

func (m *Module) Enable() error {
	log.Infof("plugin is enabled...%s", name)
	return nil
}

func (m *Module) Disable() error {
	log.Infof("plugin is disabled...%s", name)
	return nil
}
