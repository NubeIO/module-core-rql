package apirules

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

func (inst *RQL) ToNetwork(body []byte) any {
	var out *model.Network
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToNetworks(body []byte) any {
	var out []model.Network
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToDevice(body []byte) any {
	var out *model.Device
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToDevices(body []byte) any {
	var out []model.Device
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToPoint(body []byte) any {
	var out *model.Point
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToPoints(body []byte) any {
	var out []model.Point
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToAlert(body []byte) any {
	var out *model.Alert
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToAlerts(body []byte) any {
	var out []model.Alert
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

/*
let apiGet = RQL.Get("rc", "groups?with_hosts=true");
let groups = RQL.ToGroups(apiGet.Body());
*/

func (inst *RQL) ToGroup(body []byte) any {
	var out *model.Group
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) ToGroups(body []byte) any {
	var out []model.Group
	err := json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return out
}
