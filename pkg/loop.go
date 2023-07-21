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
		allRules, err := inst.Storage.SelectAllEnabledRules()
		if err != nil {
			//return
		}
		if firstLoop {
			inst.addAll(allRules) // add all existing rules from DB
		}

		for _, rule := range allRules {
			canRun, err := inst.Rules.CanExecute(rule.Name)
			if err != nil {
				//fmt.Println(err)
			}
			if canRun != nil && rule.Enable {
				if canRun.CanRun {
					result, _ := inst.Rules.ExecuteWithScript(rule.Name, inst.Props, rule.Script, rule.Schedule)
					if result.String() != "undefined" {
						inst.Storage.UpdateResult(rule.UUID, result)
					}

				}
			}
		}
		firstLoop = false
		time.Sleep(1 * time.Second)
	}

}
