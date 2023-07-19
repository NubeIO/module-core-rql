package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"time"
)

func (inst *Module) SelectAllRules() ([]byte, error) {
	resp, err := inst.Storage.SelectAllRules()
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) SelectRule(uuid string) ([]byte, error) {
	resp, err := inst.Storage.SelectRule(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) AddRule(b []byte) ([]byte, error) {
	var body *storage.RQLRule
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := inst.Storage.AddRule(body)
	if err != nil {
		return nil, err
	}
	err = inst.Rules.AddRule(body, inst.Props)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) UpdateRule(uuid string, b []byte) ([]byte, error) {
	var body *storage.RQLRule
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := inst.Storage.UpdateRule(uuid, body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) DeleteRule(uuid string) ([]byte, error) {
	err := inst.Storage.DeleteRule(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Message{Message: "ok"})
}

type Message struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (inst *Module) Dry(b []byte) ([]byte, error) {
	start := time.Now()
	inst.Client.Err = ""
	inst.Client.Return = nil
	inst.Client.TimeTaken = ""

	var body *rules.Body
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	if err != nil {
		inst.Client.Err = err.Error()
		return json.Marshal(inst.Client)
	}

	name := uuid.ShortUUID("")
	schedule := "1 sec"
	script := fmt.Sprint(body.Script)

	newRule := &storage.RQLRule{
		Name:     name,
		Script:   script,
		Schedule: schedule,
	}
	err = inst.Rules.AddRule(newRule, inst.Props)
	if err != nil {
		inst.Client.Err = err.Error()
		return json.Marshal(inst.Client)
	}
	value, err := inst.Rules.ExecuteAndRemove(name, inst.Props, false)
	if err != nil {
		inst.Client.Err = err.Error()
		return json.Marshal(inst.Client)
	}
	inst.Client.Return = value
	inst.Client.TimeTaken = time.Since(start).String()
	return json.Marshal(inst.Client)
}

func (inst *Module) SelectAllVariables() ([]byte, error) {
	resp, err := inst.Storage.SelectAllVariables()
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) SelectVariable(uuid string) ([]byte, error) {
	resp, err := inst.Storage.SelectVariable(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) AddVariable(b []byte) ([]byte, error) {
	var body *storage.RQLVariables
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := inst.Storage.AddVariable(body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) UpdateVariable(b []byte, uuid string) ([]byte, error) {
	var body *storage.RQLVariables
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := inst.Storage.UpdateVariable(uuid, body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (inst *Module) DeleteVariable(uuid string) ([]byte, error) {
	err := inst.Storage.DeleteVariable(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Message{Message: "ok"})
}
