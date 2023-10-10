package pkg

import (
	"fmt"
	"github.com/NubeIO/module-core-rql/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

func (inst *Module) addAll(allRules []storage.RQLRule) {
	for _, rule := range allRules {
		name := rule.Name
		schedule := rule.Schedule
		script := fmt.Sprint(rule.Script)

		newRule := &storage.RQLRule{
			Name:     name,
			Script:   script,
			Schedule: schedule,
		}
		err := inst.Rules.AddRule(newRule, inst.Props)
		if err != nil {
			log.Info(fmt.Sprintf("%s", err.Error()))
		}
	}
}

func (inst *Module) Loop() {
	var firstLoop = true
	for {
		if inst.ErrorOnDB {
			continue
		}
		allRules, err := inst.Storage.SelectAllEnabledRules()
		if err != nil {
			continue
		}
		if firstLoop {
			inst.addAll(allRules) // add all existing rules from DB
		}
		for _, rule := range allRules {
			canRun, err := inst.Rules.CanExecute(rule.Name)
			if err != nil {
				log.Errorf("%s: run rules loop execute err: %s", name, err.Error())
			}
			if canRun != nil && rule.Enable {
				if canRun.CanRun {
					result, err := inst.Rules.ExecuteWithScript(rule.Name, inst.Props, rule.Script, rule.Schedule)
					if err != nil {
						_, err := inst.Storage.UpdateResult(rule.UUID, err.Error())
						log.Errorf("%s: run rules loop update-result err: %s", name, err.Error())
						continue
					}
					if result == nil {
						continue
					}
					if result.String() != "undefined" {
						_, err := inst.Storage.UpdateResult(rule.UUID, returnType(result))
						log.Info(1111)
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
