package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-ufw/ufw"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-os/services/system"
)

func (inst *Client) EdgeFirewallList(hostIDName string) ([]ufw.UFWStatus, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]ufw.UFWStatus{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]ufw.UFWStatus), nil
}

func (inst *Client) EdgeFirewallStatus(hostIDName string) (*ufw.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/status/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&ufw.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

func (inst *Client) EdgeFirewallEnable(hostIDName string) (*ufw.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/enable/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&ufw.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

func (inst *Client) EdgeFirewallDisable(hostIDName string) (*ufw.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/disable/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&ufw.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

func (inst *Client) EdgeFirewallPortOpen(hostIDName string, body system.UFWBody) (*ufw.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/port/open/")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&ufw.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}

func (inst *Client) EdgeFirewallPortClose(hostIDName string, body system.UFWBody) (*ufw.Message, error) {
	url := fmt.Sprintf("/proxy/ros/api/networking/firewall/port/close/")
	if body.Port == 1662 {
		return nil, errors.New("port 1662 can not be closed")
	}
	if body.Port == 22 {
		return nil, errors.New("port 22 can not be closed")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		SetResult(&ufw.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ufw.Message), nil
}
