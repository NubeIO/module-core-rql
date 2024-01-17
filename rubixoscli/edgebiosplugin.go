package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeListPlugins(hostUUID string) ([]dto.Plugin, error, error) {
	url := fmt.Sprintf("/api/host/ros/plugins")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]dto.Plugin{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]dto.Plugin)
	return *data, nil, nil
}

func (inst *Client) EdgeUploadPlugin(hostUUID string, body *dto.Plugin) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/plugins/upload")
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

func (inst *Client) EdgeMoveFromDownloadToInstallPlugins(hostUUID string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/plugins/move-from-download-to-install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeDeletePlugin(hostUUID, pluginName, arch string) (*dto.Message, error) {
	url := fmt.Sprintf("/api/host/ros/plugins/name/%s?arch=%s", pluginName, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&dto.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

func (inst *Client) EdgeDeleteDownloadPlugins(hostUUID string) (*dto.Message, error, error) {
	url := fmt.Sprintf("/api/host/ros/plugins/download-plugins")
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
