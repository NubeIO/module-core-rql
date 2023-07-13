package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/services/systemctl"
)

func (inst *Client) EdgeAppUpload(hostIDName string, app *interfaces.AppUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/edge/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeAppInstall(hostIDName string, app *systemctl.ServiceFile) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/edge/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeAppUninstall(hostIDName string, appName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/edge/apps/uninstall")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetQueryParam("app_name", appName).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeListAppsStatus(hostIDName string) ([]interfaces.AppsStatus, error) {
	url := fmt.Sprintf("/api/edge/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]interfaces.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]interfaces.AppsStatus)
	return *data, nil
}

func (inst *Client) EdgeAppStatus(hostIDName, appName string) (*interfaces.AppsStatus, error, error) {
	url := fmt.Sprintf("/api/edge/apps/status/%s", appName)
	resp, connectionError, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.AppsStatus{}).
		Get(url))
	if connectionError != nil || requestErr != nil {
		return nil, connectionError, requestErr
	}
	return resp.Result().(*interfaces.AppsStatus), nil, nil
}
