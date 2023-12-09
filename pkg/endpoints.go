package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/dop251/goja"
	"time"
)

func (m *Module) check() error {
	if !m.pluginIsEnabled {
		return errors.New("please enable module")
	}
	if m.Storage == nil {
		return errors.New("failed to init module storage")
	}
	return nil
}

func (m *Module) SelectAllRules() ([]byte, error) {
	resp, err := m.Storage.SelectAllRules()
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) SelectRule(uuid string) ([]byte, error) {
	resp, err := m.Storage.SelectRule(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) AddRule(b []byte) ([]byte, error) {
	var body *storage.RQLRule
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := m.Storage.AddRule(body)
	if err != nil {
		return nil, err
	}
	err = m.Rules.AddRule(body, m.Props)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) UpdateRule(uuid string, b []byte) ([]byte, error) {
	var body *storage.RQLRule
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := m.Storage.UpdateRule(uuid, body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) DeleteRule(uuid string) ([]byte, error) {
	err := m.Storage.DeleteRule(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Message{Message: "ok"})
}

type RunExistingBody struct {
	Body interface{} `json:"body"`
}

// ReuseRuleRun will be used when wanting to run a pre created rule
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

- add a new script called test and paste in the example code above
- call the end point below as per the curl example

curl -i -X POST -H "Content-Type: application/json" -d '{"body":{"a":100}}' http://0.0.0.0:1660/api/modules/module-core-rql/rules/run/test
*/
func (m *Module) ReuseRuleRun(b []byte, nameUUID string) ([]byte, error) {
	start := time.Now()
	m.Client.Err = ""
	m.Client.Return = nil
	m.Client.TimeTaken = ""

	var body *RunExistingBody
	err := json.Unmarshal(b, &body)
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}

	existingRule, err := m.Storage.SelectRule(nameUUID)
	if err != nil {
		return nil, errors.New("failed to get existing rule to run")
	}
	if existingRule == nil {
		return nil, errors.New("failed to get existing rule to run")
	}

	name := uuid.ShortUUID("")
	m.Props["Input"] = body
	newRule := &storage.RQLRule{
		Name:     name,
		Script:   existingRule.Script,
		Schedule: "1 sec",
	}
	err = m.Rules.AddRule(newRule, m.Props)
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}
	value, err := m.Rules.ExecuteAndRemove(name, m.Props, false)
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}
	m.Client.Return = returnType(value)
	m.Client.TimeTaken = time.Since(start).String()
	return json.Marshal(m.Client)
}

type Message struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (m *Module) Dry(b []byte) ([]byte, error) {
	start := time.Now()
	m.Client.Err = ""
	m.Client.Return = nil
	m.Client.TimeTaken = ""

	var body *rules.Body
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}

	name := uuid.ShortUUID("")
	schedule := "1 sec"
	script := fmt.Sprint(body.Script)

	newRule := &storage.RQLRule{
		Name:     name,
		Script:   script,
		Schedule: schedule,
	}
	err = m.Rules.AddRule(newRule, m.Props)
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}
	value, err := m.Rules.ExecuteAndRemove(name, m.Props, false)
	if err != nil {
		m.Client.Err = err.Error()
		return json.Marshal(m.Client)
	}
	m.Client.Return = returnType(value)
	m.Client.TimeTaken = time.Since(start).String()
	return json.Marshal(m.Client)
}

func returnType(value goja.Value) any {
	if value == nil {
		return nil
	}
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

func (m *Module) SelectAllVariables() ([]byte, error) {
	resp, err := m.Storage.SelectAllVariables()
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) SelectVariable(uuid string) ([]byte, error) {
	resp, err := m.Storage.SelectVariable(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) AddVariable(b []byte) ([]byte, error) {
	var body *storage.RQLVariables
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := m.Storage.AddVariable(body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) UpdateVariable(b []byte, uuid string) ([]byte, error) {
	var body *storage.RQLVariables
	err := json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}
	resp, err := m.Storage.UpdateVariable(uuid, body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp)
}

func (m *Module) DeleteVariable(uuid string) ([]byte, error) {
	err := m.Storage.DeleteVariable(uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Message{Message: "ok"})
}
