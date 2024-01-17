package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeListModules(hostUUID string) ([]dto.Module, error, error) {
	url := fmt.Sprintf("/api/host/ros/modules")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]dto.Module{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]dto.Module)
	return *data, nil, nil
}

func (inst *Client) EdgeUploadModule(hostUUID string, body *dto.Module) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/modules/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeMoveFromDownloadToInstallModules(hostUUID string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/modules/move-from-download-to-install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeDeleteModule(hostUUID, pluginName string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/modules/name/%s", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeDeleteDownloadModules(hostUUID string) (*dto.Message, error, error) {
	url := fmt.Sprintf("/api/host/ros/modules/download-modules")
	// we use v2 here, coz it shows requestErr when there is no plugins' directory on download path
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Delete(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return resp.Result().(*dto.Message), nil, nil
}
