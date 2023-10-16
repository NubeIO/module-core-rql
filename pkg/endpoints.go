package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/dop251/goja"
)

type Message struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (inst *Module) check() error {
	if !inst.pluginIsEnabled {
		return errors.New("please enable module")
	}
	if inst.Storage == nil {
		return errors.New("failed to init module storage")
	}
	return nil
}

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

// RunExistingRuleOnce will be used when wanting to run a pre created rule
// get existing rule script
// pass in some parameters
// run the rule and return the result
/* rql code
let x = Input.Body.a;
let y = RQL.RandInt(10, 20);
let calc = x + y;
let out = {
  inputValue: x,
  randomNumber: y,
  calc: calc,
};

RQL.Result = out;
*/
/* body
{
    "body":{
        "a":100
    }
}
*/
func (inst *Module) RunExistingRuleOnce(nameUUID string) ([]byte, error) {
	existingRule, err := inst.Storage.SelectRule(nameUUID)
	if err != nil {
		inst.Client.Err = err.Error()
		return json.Marshal(inst.Client)
	}
	if existingRule == nil {
		inst.Client.Err = "failed to get existing rule to run"
		return json.Marshal(inst.Client)
	}

	return inst.executeRuleOnce(existingRule.Script)
}

func (inst *Module) DryRunFromEditor(b []byte) ([]byte, error) {
	var body *rules.Body
	err := json.Unmarshal(b, &body)
	if err != nil {
		inst.Client.Err = err.Error()
		return json.Marshal(inst.Client)
	}
	script := fmt.Sprint(body.Script)
	return inst.executeRuleOnce(script)
}

func (inst *Module) executeRuleOnce(script string) ([]byte, error) {
	start := time.Now()
	inst.Client.Err = ""
	inst.Client.Return = nil
	inst.Client.TimeTaken = ""

	newRule := &storage.RQLRule{
		Name:     uuid.ShortUUID(""),
		Script:   script,
		Schedule: "1 sec",
	}

	err := inst.Rules.AddRule(newRule, inst.Props)
	if err != nil {
		return nil, err
	}
	value, err := inst.Rules.ExecuteAndRemove(newRule.Name, inst.Props, false)
	if err != nil {
		inst.Client.Err = err.Error()
		inst.Client.TimeTaken = time.Since(start).String()
		return json.Marshal(inst.Client)
	}

	inst.Client.Return = returnType(value)
	inst.Client.TimeTaken = time.Since(start).String()
	return json.Marshal(inst.Client)
}

func returnType(value goja.Value) any {
	if value.String() == "undefined" {
		return value.String()
	}
	t := value.ExportType().String()
	if t == "string" {
		return value.String()
	}
	if t == "*errors.errorString" {
		return value.String()
	}
	if t == "string" {
		return value.String()
	}
	if t == "int64" {
		return value.ToInteger()
	}
	if t == "float64" {
		return value.ToFloat()
	}
	return value
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
