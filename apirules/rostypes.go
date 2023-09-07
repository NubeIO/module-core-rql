package apirules

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/go-resty/resty/v2"
)

func (inst *RQL) ToNetwork(body *resty.Response) *model.Network {
	var out *model.Network
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToNetworks(body *resty.Response) []model.Network {
	var out []model.Network
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToDevice(body *resty.Response) *model.Device {
	var out *model.Device
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToDevices(body *resty.Response) []model.Device {
	var out []model.Device
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToPoint(body *resty.Response) *model.Point {
	var out *model.Point
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToPoints(body *resty.Response) []model.Point {
	var out []model.Point
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToAlert(body *resty.Response) *model.Alert {
	var out *model.Alert
	json.Unmarshal(body.Body(), &out)
	return out
}

func (inst *RQL) ToAlerts(body *resty.Response) []model.Alert {
	var out []model.Alert
	json.Unmarshal(body.Body(), &out)
	return out
}
