package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeListPlugins(hostIDName string) ([]interfaces.Plugin, error, error) {
	url := fmt.Sprintf("/api/eb/ros/plugins")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]interfaces.Plugin{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]interfaces.Plugin)
	return *data, nil, nil
}

func (inst *Client) EdgeUploadPlugin(hostIDName string, body *interfaces.Plugin) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/plugins/upload")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeMoveFromDownloadToInstallPlugins(hostIDName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/plugins/move-from-download-to-install")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeDeletePlugin(hostIDName, pluginName, arch string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/plugins/name/%s?arch=%s", pluginName, arch)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}

func (inst *Client) EdgeDeleteDownloadPlugins(hostIDName string) (*interfaces.Message, error, error) {
	url := fmt.Sprintf("/api/eb/ros/plugins/download-plugins")
	// we use v2 here, coz it shows requestErr when there is no plugins' directory on download path
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Delete(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	return resp.Result().(*interfaces.Message), nil, nil
}
