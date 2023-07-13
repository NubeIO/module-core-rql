package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-os/interfaces"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/NubeIO/rubix-ui/backend/constants"
)

func (inst *Client) GetEdgeDeviceInfoPublic(hostIDName string) (*rubixregistry.DeviceInfo, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/device/public")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&rubixregistry.DeviceInfo{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	deviceInfo := resp.Result().(*rubixregistry.DeviceInfo)
	if deviceInfo.DeviceType == "" {
		deviceInfo.DeviceType = constants.RubixCompute.String()
	}
	return deviceInfo, nil
}

func (inst *Client) GetEdgeDeviceInfo(hostIDName string) (*rubixregistry.DeviceInfo, error) {
	url := fmt.Sprintf("/proxy/ros/api/system/device")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&rubixregistry.DeviceInfo{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	deviceInfo := resp.Result().(*rubixregistry.DeviceInfo)
	if deviceInfo.DeviceType == "" {
		deviceInfo.DeviceType = constants.RubixCompute.String()
	}
	return deviceInfo, nil
}

func (inst *Client) UpdateEdgeDeviceInfo(hostIDName string, body *rubixregistry.DeviceInfo) (*rubixregistry.DeviceInfo,
	error) {
	url := fmt.Sprintf("/proxy/ros/api/system/device")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&rubixregistry.DeviceInfo{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	deviceInfo := resp.Result().(*rubixregistry.DeviceInfo)
	if deviceInfo.DeviceType == "" {
		deviceInfo.DeviceType = constants.RubixCompute.String()
	}
	return deviceInfo, nil
}

func (inst *Client) EdgeSystemCtlAction(hostIDName, serviceName string, action interfaces.Action) (*interfaces.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/systemctl/%s?unit=%s", action, serviceName)
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

func (inst *Client) EdgeSystemCtlState(hostIDName, serviceName string) (*systemctl.SystemState, error) {
	url := fmt.Sprintf("/proxy/ros/api/systemctl/state?unit=%s", serviceName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&systemctl.SystemState{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*systemctl.SystemState), nil
}
