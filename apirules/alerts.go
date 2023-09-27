package apirules

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/mitchellh/mapstructure"
)

/*
let body = {
  hostUUID: "hos_c8ab2fe07e7a413c",
  entityType: "device",
  type: "ping",
  status: "active",
  severity: "crucial",
};

let alert = RQL.AddAlert("hos_c8ab2fe07e7a413c", body);

RQL.Result = alert.CreatedAt;
*/

func alertBody(body any) (*model.Alert, error) {
	result := &model.Alert{}
	err := mapstructure.Decode(body, &result)
	return result, err
}

func (inst *RQL) GetAlerts() any {
	resp, err := cli.GetAlerts()
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) AddAlert(hostIDName string, body any) any {
	b, err := alertBody(body)
	if err != nil {
		return err
	}
	resp, err := cli.AddAlert(hostIDName, b)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) UpdateAlertStatus(hostIDName, uuid, status string) any {
	resp, err := cli.UpdateAlertStatus(hostIDName, uuid, status)
	if err != nil {
		return err
	}
	return resp
}
