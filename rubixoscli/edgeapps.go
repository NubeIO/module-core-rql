package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/services/systemctl"
)

func (inst *Client) EdgeAppUpload(hostUUID string, app *interfaces.AppUpload) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/host/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&interfaces.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeAppInstall(hostUUID string, app *systemctl.ServiceFile) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/host/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&interfaces.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeAppUninstall(hostUUID string, appName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/host/apps/uninstall")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetQueryParam("app_name", appName).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeListAppsStatus(hostUUID string) ([]interfaces.AppsStatus, error) {
	url := fmt.Sprintf("/api/host/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]interfaces.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]interfaces.AppsStatus)
	return *data, nil
}

func (inst *Client) EdgeAppStatus(hostUUID, appName string) (*interfaces.AppsStatus, error, error) {
	url := fmt.Sprintf("/api/host/apps/name/%s/status", appName)
	resp, connectionError, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&interfaces.AppsStatus{}).
		Get(url))
	if connectionError != nil || requestErr != nil {
		return nil, connectionError, requestErr
	}
	return resp.Result().(*interfaces.AppsStatus), nil, nil
}

func (inst *Client) EdgeListRestartJobs(hostUUID string) ([]interfaces.RestartJob, error) {
	url := fmt.Sprintf("/host/ros/api/restart-jobs")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]interfaces.RestartJob{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]interfaces.RestartJob)
	return *data, nil
}
