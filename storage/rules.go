package storage

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"time"
)

func (inst *db) AddRule(rc *RQLRule) (*RQLRule, error) {
	rc.UUID = uuid.ShortUUID("rql")
	err := inst.DB.Insert(rc)
	return rc, err
}

func (inst *db) UpdateRule(uuid string, rc *RQLRule) (*RQLRule, error) {
	if rc != nil {
		rc.UUID = uuid
	}
	err := inst.DB.Update(rc)
	return rc, err
}

func (inst *db) UpdateResult(uuid string, result interface{}) (*RQLRule, error) {
	rule, err := inst.SelectRule(uuid)
	if err != nil {
		return nil, err
	}
	r := Result{
		Result:    result,
		Timestamp: time.Now().Format(time.RFC1123),
		Time:      time.Now(),
	}
	if rule.ResultStorageSize < 10 {
		rule.ResultStorageSize = 10
	}
	if rule.ResultStorageSize > 100 {
		rule.ResultStorageSize = 100
	}
	if len(rule.Result) > rule.ResultStorageSize {
		rule.Result = rule.Result[1:]
	} else {
		rule.Result = append(rule.Result, r)
	}
	return inst.UpdateRule(uuid, rule)
}

func (inst *db) DeleteRule(uuid string) error {
	rule, err := inst.SelectRule(uuid)
	if err != nil {
		return err
	}
	return inst.DB.Delete(rule)
}

func (inst *db) SelectRule(uuidName string) (*RQLRule, error) {
	resp, err := inst.selectRule(uuidName)
	if resp == nil || err != nil {
		resp, err = inst.selectRuleByName(uuidName)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (inst *db) selectRuleByName(name string) (*RQLRule, error) {
	var data *RQLRule
	err := inst.DB.Open(RQLRule{}).Where("name", "=", name).First().AsEntity(&data)
	return data, err

}

func (inst *db) selectRule(uuid string) (*RQLRule, error) {
	var data *RQLRule
	err := inst.DB.Open(RQLRule{}).Where("uuid", "=", uuid).First().AsEntity(&data)
	return data, err

}

func (inst *db) SelectAllRules() ([]RQLRule, error) {
	var resp []RQLRule
	inst.DB.Open(RQLRule{}).AsEntity(&resp)
	return resp, nil
}

func (inst *db) SelectAllEnabledRules() ([]RQLRule, error) {
	var resp []RQLRule
	rules, err := inst.SelectAllRules()
	if err != nil {
		return nil, err
	}
	for _, rule := range rules {
		if rule.Enable {
			resp = append(resp, rule)
		}
	}
	return resp, nil
}
