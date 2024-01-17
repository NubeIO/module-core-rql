package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/services/systemctl"
)

func (inst *Client) EdgeAppUpload(hostUUID string, app *dto.AppUpload) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/apps/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeAppInstall(hostUUID string, app *systemctl.ServiceFile) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/apps/install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		SetBody(app).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeAppUninstall(hostUUID string, appName string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/apps/uninstall")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetQueryParam("app_name", appName).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeListAppsStatus(hostUUID string) ([]dto.AppsStatus, error) {
	url := fmt.Sprintf("/api/host/apps/status")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]dto.AppsStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]dto.AppsStatus)
	return *data, nil
}

func (inst *Client) EdgeAppStatus(hostUUID, appName string) (*dto.AppsStatus, error, error) {
	url := fmt.Sprintf("/api/host/apps/name/%s/status", appName)
	resp, connectionError, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.AppsStatus{}).
		Get(url))
	if connectionError != nil || requestErr != nil {
		return nil, connectionError, requestErr
	}
	return resp.Result().(*dto.AppsStatus), nil, nil
}

func (inst *Client) EdgeListRestartJobs(hostUUID string) ([]dto.RestartJob, error) {
	url := fmt.Sprintf("/host/ros/api/restart-jobs")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]dto.RestartJob{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]dto.RestartJob)
	return *data, nil
}
