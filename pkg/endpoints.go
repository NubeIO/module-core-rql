package pkg

import (
	"encoding/json"
	"github.com/NubeIO/module-core-rql/storage"
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
