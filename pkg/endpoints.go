package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/module-core-rql/rules"
	"github.com/NubeIO/module-core-rql/storage"
	"time"
)

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
	m.Client.Return = value
	m.Client.TimeTaken = time.Since(start).String()
	return json.Marshal(m.Client)
}
