package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) EdgeListModules(hostIDName string) ([]interfaces.Module, error, error) {
	url := fmt.Sprintf("/api/eb/ros/modules")
	resp, connectionErr, requestErr := nresty.FormatRestyV2Response(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]interfaces.Module{}).
		Get(url))
	if connectionErr != nil || requestErr != nil {
		return nil, connectionErr, requestErr
	}
	data := resp.Result().(*[]interfaces.Module)
	return *data, nil, nil
}

func (inst *Client) EdgeUploadModule(hostIDName string, body *interfaces.Module) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/modules/upload")
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

func (inst *Client) EdgeMoveFromDownloadToInstallModules(hostIDName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/modules/move-from-download-to-install")
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

func (inst *Client) EdgeDeleteModule(hostIDName, pluginName string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/api/eb/ros/modules/name/%s", pluginName)
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

func (inst *Client) EdgeDeleteDownloadModules(hostIDName string) (*interfaces.Message, error, error) {
	url := fmt.Sprintf("/api/eb/ros/modules/download-modules")
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
