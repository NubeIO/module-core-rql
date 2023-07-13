package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) AddDevice(hostIDName string, device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/proxy/ros/api/devices?with_tags=true&with_meta_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Device{}).
		SetBody(device).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) GetDevices(hostIDName string, withPoints ...bool) ([]model.Device, error) {
	url := fmt.Sprintf("/proxy/ros/api/devices?with_tags=true&with_meta_tags=true")
	if len(withPoints) > 0 {
		if withPoints[0] == true {
			url = fmt.Sprintf("/proxy/ros/api/devices?with_points=true&with_tags=true&with_meta_tags=true")
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Device{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Device
	out = *resp.Result().(*[]model.Device)
	return out, nil
}

func (inst *Client) GetDevice(hostIDName, uuid string, withPoints ...bool) (*model.Device, error) {
	url := fmt.Sprintf("/proxy/ros/api/devices/%s?with_tags=true&with_meta_tags=true", uuid)
	if len(withPoints) > 0 {
		if withPoints[0] == true {
			url = fmt.Sprintf("/proxy/ros/api/devices/%s?with_points=true&with_tags=true&with_meta_tags=true", uuid)
		}
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Device{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) EditDevice(hostIDName, uuid string, device *model.Device) (*model.Device, error) {
	url := fmt.Sprintf("/proxy/ros/api/devices/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Device{}).
		SetBody(device).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Device), nil
}

func (inst *Client) DeleteDevice(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/proxy/ros/api/devices/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) SyncDevices(hostIDName, networkUUID string) (*interfaces.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks/%s/sync/devices?with_points=true", networkUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&interfaces.Message{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*interfaces.Message), nil
}
