package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/services/system"
)

func (inst *Client) EdgeGetNetworks(hostUUID string) ([]networking.NetworkInterfaces, error) {
	url := fmt.Sprintf("/host/ros/api/networking")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]networking.NetworkInterfaces{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]networking.NetworkInterfaces)
	return *data, nil
}

// EdgeDHCPPortExists check if the interface is a static or fixed ip, will return true if port is a set to dhcp
func (inst *Client) EdgeDHCPPortExists(hostUUID string, body *system.NetworkingBody) (*system.DHCPPortExists, error) {
	url := fmt.Sprintf("/host/ros/api/networking/interfaces/exists")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.DHCPPortExists{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.DHCPPortExists), nil
}

func (inst *Client) EdgeDHCPSetAsAuto(hostUUID string, body *system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/networking/interfaces/auto")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&system.Message{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeDHCPSetStaticIP(hostUUID string, body *dhcpd.SetStaticIP) (string, error) {
	url := fmt.Sprintf("/host/ros/api/networking/interfaces/static")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(url))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (inst *Client) EdgeRestartNetworking(hostUUID string) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/networking/networks/restart")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeInterfaceUpDown(hostUUID string, port system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/networking/networks/interfaces/reset")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(port).
		SetResult(system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeInterfaceUp(hostUUID string, port system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/networking/networks/interfaces/up")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(port).
		SetResult(system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}

func (inst *Client) EdgeInterfaceDown(hostUUID string, port system.NetworkingBody) (*system.Message, error) {
	url := fmt.Sprintf("/host/ros/api/networking/networks/interfaces/down")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(port).
		SetResult(system.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*system.Message), nil
}
