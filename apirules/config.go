package apirules

import (
	"encoding/json"
)

type config struct {
	SlackToken string
	LogLevel   string
}

func (inst *RQL) getConfig() (*config, error) {
	marshal, err := json.Marshal(inst.Config)
	if err != nil {
		return nil, err
	}
	c := &config{}
	err = json.Unmarshal(marshal, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
