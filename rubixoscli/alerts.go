package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) AddAlert(hostIDName string, body *model.Alert) (*model.Alert, error) {
	if body == nil {
		return nil, errors.New("alerts body can not be empty")
	}
	if body.HostUUID == "" {
		body.HostUUID = hostIDName
	}
	if body.Severity == "" {
		body.Severity = "info"
	}
	if body.Status == "" {
		body.Status = "active"
	}
	path := fmt.Sprintf("%s", Paths.Alerts.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Alert{}).
		SetBody(body).
		Post(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Alert), nil
}

func (inst *Client) GetAlerts(hostIDName string) ([]model.Alert, error) {
	path := fmt.Sprintf("%s", Paths.Alerts.Path)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Alert{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	var out []model.Alert
	out = *resp.Result().(*[]model.Alert)
	return out, nil
}

func (inst *Client) GetAlertsByHost(hostIDName string) ([]model.Alert, error) {
	path := fmt.Sprintf("%s/host/%s", Paths.Alerts.Path, hostIDName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Alert{}).
		Get(path))
	if err != nil {
		return nil, err
	}
	var out []model.Alert
	out = *resp.Result().(*[]model.Alert)
	return out, nil
}

func (inst *Client) DeleteAlert(hostIDName, uuid string) (*Message, error) {
	path := fmt.Sprintf("%s/%s", Paths.Alerts.Path, uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&Message{}).
		Delete(path))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Message), nil
}
