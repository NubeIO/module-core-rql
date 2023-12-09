package pkg

import (
	"fmt"
	"github.com/NubeIO/module-core-rql/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m *Module) addAll(allRules []storage.RQLRule) {
	for _, rule := range allRules {
		name := rule.Name
		schedule := rule.Schedule
		script := fmt.Sprint(rule.Script)

		newRule := &storage.RQLRule{
			Name:     name,
			Script:   script,
			Schedule: schedule,
		}
		err := m.Rules.AddRule(newRule, m.Props)
		if err != nil {
			log.Info(fmt.Sprintf("%s", err.Error()))
		}
	}
}

func (m *Module) Loop() {
	var firstLoop = true
	for {
		if m.ErrorOnDB {
			continue
		}
		allRules, err := m.Storage.SelectAllEnabledRules()
		if err != nil {
			continue
		}
		if firstLoop {
			m.addAll(allRules) // add all existing rules from DB
		}
		for _, rule := range allRules { // TODO add a lock, so we can add a goroutine
			if !rule.Enable {
				continue
			}
			canRun, err := m.Rules.CanExecute(rule.Name)
			if err != nil {
				log.Errorf("%s: run rules loop execute err: %s", name, err.Error())
			}
			if canRun != nil && rule.Enable {
				if canRun.CanRun {
					result, err := m.Rules.ExecuteWithScript(rule.Name, m.Props, rule.Script, rule.Schedule)
					if err != nil {
						log.Errorf("%s: run rules loop execute-with-script err: %s", name, err.Error())
						_, e := m.Storage.UpdateResult(rule.UUID, err.Error())
						if e != nil {
							log.Errorf("%s: run rules loop update-result err: %s", name, e.Error())
						}
						continue
					}
					if result == nil {
						continue
					}
					if result.String() != "undefined" {
						_, err := m.Storage.UpdateResult(rule.UUID, returnType(result))
						log.Info(result)
						if err != nil {
							log.Errorf("%s: run rules loop update-result err: %s", name, err.Error())
						}
					}
				}
			}
		}
		firstLoop = false
		time.Sleep(1 * time.Second)
	}

}
